// Package site is what a website itself needs from Seolyticize: the
// analytics tag, the IndexNow key file, deploy notification and server-side
// events. Standard library only.
package site

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"strings"
	"time"
)

// DefaultBaseURL is the hosted Seolyticize.
const DefaultBaseURL = "https://seolyticize.com"

// TagOptions configures AnalyticsTag.
type TagOptions struct {
	Exclude []string // paths never reported: "/app/*", "/verify"
	Include []string // when set, only these paths are reported
	Track   []string // "outbound", "downloads"
	BaseURL string
}

// AnalyticsTag is the cookieless analytics <script> for every page's <head>,
// escaped for HTML. In html/template, wrap it in template.HTML.
func AnalyticsTag(siteKey string, options TagOptions) string {
	base := options.BaseURL
	if base == "" {
		base = DefaultBaseURL
	}
	attrs := []string{`defer`, fmt.Sprintf(`data-site="%s"`, html.EscapeString(siteKey))}
	if len(options.Exclude) > 0 {
		attrs = append(attrs, fmt.Sprintf(`data-exclude="%s"`, html.EscapeString(strings.Join(options.Exclude, ","))))
	}
	if len(options.Include) > 0 {
		attrs = append(attrs, fmt.Sprintf(`data-include="%s"`, html.EscapeString(strings.Join(options.Include, ","))))
	}
	if len(options.Track) > 0 {
		attrs = append(attrs, fmt.Sprintf(`data-track="%s"`, html.EscapeString(strings.Join(options.Track, " "))))
	}
	attrs = append(attrs, fmt.Sprintf(`src="%s/a.js"`, html.EscapeString(strings.TrimRight(base, "/"))))
	return "<script " + strings.Join(attrs, " ") + "></script>"
}

// IndexNowKeyHandler serves the IndexNow key file. Mount it at "/"+key+".txt":
//
//	mux.Handle("/"+key+".txt", site.IndexNowKeyHandler(key))
func IndexNowKeyHandler(key string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Cache-Control", "public, max-age=3600")
		_, _ = w.Write([]byte(key))
	})
}

// DeployResult is what a deploy hook reports.
type DeployResult struct {
	Sitemaps     []string `json:"sitemaps"`
	SitemapFound bool     `json:"sitemapFound"`
	FirstRun     bool     `json:"firstRun"`
	PagesKnown   int      `json:"pagesKnown"`
	Changed      int      `json:"changed"`
	Submitted    int      `json:"submitted"`
	Result       string   `json:"result"`
	Error        string   `json:"error"`
	Code         string   `json:"code"`
}

// NotifyDeploy tells Seolyticize a deploy finished (hookURL from
// create_deploy_hook). With no urls, the sitemap is compared with the last
// deploy's; with them, exactly those go to IndexNow.
func NotifyDeploy(ctx context.Context, hookURL string, urls []string) (*DeployResult, error) {
	body := map[string]any{}
	if urls != nil {
		body["urls"] = urls
	}
	var result DeployResult
	status, err := post(ctx, hookURL, body, &result, 2*time.Minute)
	if err != nil {
		return nil, err
	}
	if status >= 300 {
		return &result, fmt.Errorf("seolyticize: deploy hook answered %d: %s", status, result.Error)
	}
	return &result, nil
}

// Event is one analytics event sent from a server.
type Event struct {
	SiteKey  string
	URL      string
	Name     string // "pageview" when empty
	Referrer string
	Props    map[string]any
	BaseURL  string
}

// TrackEvent sends an event to /api/collect. Every event from one server
// address counts as one visitor, so send from the browser where you can.
func TrackEvent(ctx context.Context, event Event) error {
	base := event.BaseURL
	if base == "" {
		base = DefaultBaseURL
	}
	name := event.Name
	if name == "" {
		name = "pageview"
	}
	body := map[string]any{"k": event.SiteKey, "n": name, "u": event.URL}
	if event.Referrer != "" {
		body["r"] = event.Referrer
	}
	if len(event.Props) > 0 {
		body["p"] = event.Props
	}
	status, err := post(ctx, strings.TrimRight(base, "/")+"/api/collect", body, nil, 10*time.Second)
	if err != nil {
		return err
	}
	if status >= 300 {
		return fmt.Errorf("seolyticize: collect answered %d", status)
	}
	return nil
}

func post(ctx context.Context, url string, body any, into any, timeout time.Duration) (int, error) {
	encoded, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(encoded))
	if err != nil {
		return 0, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "seolyticize-go-site/0.1.0")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	if into != nil {
		_ = json.NewDecoder(response.Body).Decode(into)
	}
	return response.StatusCode, nil
}
