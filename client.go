// Package seolyticize is the Go client for the Seolyticize API.
//
//	client := seolyticize.New(os.Getenv("SEOLYTICIZE_API_KEY"))
//	steps, err := client.GetNextSteps(ctx, seolyticize.GetNextStepsInput{ClientID: "client_…"})
//
// Every tool is a method with a typed input; Call reaches tools newer than
// this SDK. Results are the tool's JSON, as json.RawMessage, to decode into
// whatever shape the caller wants. The on-site helpers are in package site.
package seolyticize

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// DefaultBaseURL is the hosted Seolyticize API.
const DefaultBaseURL = "https://seolyticize.com"

// Version is this SDK's version, sent as the user agent.
const Version = "0.2.0"

// Client calls the Seolyticize API. The zero value is not usable; use New.
type Client struct {
	APIKey  string
	BaseURL string
	HTTP    *http.Client
	// MaxRetries after the first attempt. 429 is always retried (it ran
	// nothing); server and network errors only for read-only tools.
	MaxRetries int
	// MaxRetryAfter is the longest Retry-After the client will wait.
	MaxRetryAfter time.Duration
	sleep         func(time.Duration)
}

// Option configures a Client.
type Option func(*Client)

// WithBaseURL points the client at another deployment.
func WithBaseURL(base string) Option {
	return func(c *Client) { c.BaseURL = strings.TrimRight(base, "/") }
}

// WithHTTPClient replaces the HTTP client (timeouts, proxies, tests).
func WithHTTPClient(h *http.Client) Option { return func(c *Client) { c.HTTP = h } }

// WithMaxRetries sets how many times a call is retried.
func WithMaxRetries(n int) Option { return func(c *Client) { c.MaxRetries = n } }

// New makes a client. An empty apiKey falls back to SEOLYTICIZE_API_KEY.
func New(apiKey string, options ...Option) *Client {
	if apiKey == "" {
		apiKey = os.Getenv("SEOLYTICIZE_API_KEY")
	}
	base := os.Getenv("SEOLYTICIZE_BASE_URL")
	if base == "" {
		base = DefaultBaseURL
	}
	c := &Client{
		APIKey:        apiKey,
		BaseURL:       strings.TrimRight(base, "/"),
		HTTP:          &http.Client{Timeout: 60 * time.Second},
		MaxRetries:    2,
		MaxRetryAfter: 60 * time.Second,
		sleep:         time.Sleep,
	}
	for _, option := range options {
		option(c)
	}
	return c
}

// Error is what the API answers when it says no. Branch on Code, not Message.
type Error struct {
	Status    int    // HTTP status; 0 when no response arrived
	Message   string // a sentence written for the caller
	Code      string // stable reason, e.g. "search_console_no_access"
	Next      string // the call or step that fixes it, when there is one
	RequestID string // quote this when asking for help
	Tool      string
	// RetryAfter is set on a 429.
	RetryAfter time.Duration
}

func (e *Error) Error() string {
	text := e.Message
	if e.Code != "" {
		text += " (code: " + e.Code + ")"
	}
	if e.Next != "" {
		text += " Next: " + e.Next
	}
	return text
}

// IsCode reports whether err is a Seolyticize Error with that code.
func IsCode(err error, code string) bool {
	var apiErr *Error
	return errors.As(err, &apiErr) && apiErr.Code == code
}

// Call runs any tool by name. input is anything that marshals to the tool's
// JSON input (a generated *Input struct, a map, or nil).
func (c *Client) Call(ctx context.Context, tool string, input any) (json.RawMessage, error) {
	if input == nil {
		input = map[string]any{}
	}
	body, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("seolyticize: encoding %s input: %w", tool, err)
	}
	return c.do(ctx, http.MethodPost, "/api/v1/tools/"+url.PathEscape(tool), body, tool, readOnlyTools[tool])
}

// ListTools returns the tools this key may call, with their schemas. group
// narrows it to one job: onboarding, weekly, technical, research, reporting.
func (c *Client) ListTools(ctx context.Context, group string) (json.RawMessage, error) {
	path := "/api/v1/tools"
	if group != "" {
		path += "?group=" + url.QueryEscape(group)
	}
	return c.do(ctx, http.MethodGet, path, nil, "", true)
}

func (c *Client) do(ctx context.Context, method, path string, body []byte, tool string, retryTransient bool) (json.RawMessage, error) {
	if c.APIKey == "" {
		return nil, &Error{Message: "An API key is required: pass one to New or set SEOLYTICIZE_API_KEY.", Code: "no_api_key"}
	}
	for attempt := 0; ; attempt++ {
		var reader io.Reader
		if body != nil {
			reader = bytes.NewReader(body)
		}
		request, err := http.NewRequestWithContext(ctx, method, c.BaseURL+path, reader)
		if err != nil {
			return nil, err
		}
		request.Header.Set("Authorization", "Bearer "+c.APIKey)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Accept", "application/json")
		request.Header.Set("User-Agent", "seolyticize-go/"+Version)

		response, err := c.HTTP.Do(request)
		if err != nil {
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}
			if retryTransient && attempt < c.MaxRetries {
				c.wait(ctx, backoff(attempt))
				continue
			}
			return nil, &Error{Message: "No response from " + c.BaseURL + ": " + err.Error(), Code: "network_error", Tool: tool}
		}
		raw, readErr := io.ReadAll(io.LimitReader(response.Body, 64<<20))
		response.Body.Close()
		if readErr != nil {
			return nil, readErr
		}
		if response.StatusCode < 300 {
			var envelope struct {
				Data json.RawMessage `json:"data"`
			}
			if json.Unmarshal(raw, &envelope) == nil && envelope.Data != nil {
				return envelope.Data, nil
			}
			return raw, nil
		}

		apiErr := errorFrom(response, raw, tool)
		canRetry := attempt < c.MaxRetries
		if response.StatusCode == http.StatusTooManyRequests && canRetry && apiErr.RetryAfter <= c.MaxRetryAfter {
			wait := apiErr.RetryAfter
			if wait == 0 {
				wait = backoff(attempt)
			}
			c.wait(ctx, wait)
			continue
		}
		if response.StatusCode >= 500 && retryTransient && canRetry {
			c.wait(ctx, backoff(attempt))
			continue
		}
		return nil, apiErr
	}
}

func errorFrom(response *http.Response, raw []byte, tool string) *Error {
	var body struct {
		Error     string `json:"error"`
		Code      string `json:"code"`
		Next      string `json:"next"`
		RequestID string `json:"requestId"`
	}
	_ = json.Unmarshal(raw, &body)
	apiErr := &Error{
		Status:    response.StatusCode,
		Message:   body.Error,
		Code:      body.Code,
		Next:      body.Next,
		RequestID: body.RequestID,
		Tool:      tool,
	}
	if apiErr.Message == "" {
		apiErr.Message = "HTTP " + strconv.Itoa(response.StatusCode)
	}
	if apiErr.RequestID == "" {
		apiErr.RequestID = response.Header.Get("X-Request-Id")
	}
	if seconds, err := strconv.ParseFloat(response.Header.Get("Retry-After"), 64); err == nil {
		apiErr.RetryAfter = time.Duration(seconds * float64(time.Second))
	}
	return apiErr
}

func (c *Client) wait(ctx context.Context, d time.Duration) {
	if c.sleep != nil && ctx.Err() == nil {
		c.sleep(d)
	}
}

func backoff(attempt int) time.Duration {
	base := 500 * time.Millisecond << attempt
	if base > 8*time.Second {
		base = 8 * time.Second
	}
	return base/2 + time.Duration(rand.Int63n(int64(base/2)+1))
}

// Ptr returns a pointer to v, for the optional number and boolean fields of
// the generated inputs: MaxCostMicros: seolyticize.Ptr(int64(500000)).
func Ptr[T any](v T) *T { return &v }
