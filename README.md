# Seolyticize for Go

Every [Seolyticize](https://seolyticize.com) API tool as a typed method, and package `site` for analytics, IndexNow and deploys. Standard library only.

```sh
go get github.com/Seolyticize/sdk-go
```

## Quick start

```go
import seolyticize "github.com/Seolyticize/sdk-go"

client := seolyticize.New("") // reads SEOLYTICIZE_API_KEY

manifest, err := client.OnboardSite(ctx, seolyticize.OnboardSiteInput{Domain: "example.com", Framework: "html"})
// …deploy the files and snippet in manifest, then:
checklist, err := client.VerifySite(ctx, seolyticize.VerifySiteInput{ClientID: "client_…"})
```

Results are the tool's JSON as `json.RawMessage`; unmarshal into the shape you need. Optional numbers and booleans are pointers: `Limit: seolyticize.Ptr(int64(50))`. `client.Call(ctx, "tool_name", input)` reaches tools newer than this module; `client.ListTools(ctx, "onboarding")` lists what your key may call.

## Errors

```go
var apiErr *seolyticize.Error
if errors.As(err, &apiErr) {
	fmt.Println(apiErr.Status, apiErr.Code, apiErr.Next, apiErr.RequestID)
}
if seolyticize.IsCode(err, "search_console_no_access") { /* … */ }
```

429s are retried after `Retry-After`; server and network errors are retried only for read-only tools.

## On your site

```go
import "github.com/Seolyticize/sdk-go/site"

tag := site.AnalyticsTag("SITE_KEY", site.TagOptions{Exclude: []string{"/account/*"}}) // template.HTML(tag)
mux.Handle("/"+key+".txt", site.IndexNowKeyHandler(key))
site.NotifyDeploy(ctx, hookURL, nil) // after each deploy
site.TrackEvent(ctx, site.Event{SiteKey: "SITE_KEY", URL: "https://example.com/pricing", Name: "Signup"})
```

Docs: [seolyticize.com/docs/go-sdk](https://seolyticize.com/docs/go-sdk). Generated from the registry that runs the API. License: MIT.
