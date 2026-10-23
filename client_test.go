package seolyticize

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func fake(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	c := New("sk_seo_test", WithBaseURL(server.URL))
	c.sleep = func(time.Duration) {}
	return c
}

func TestTypedCallSendsCamelCaseAndUnwrapsData(t *testing.T) {
	c := fake(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/tools/get_audit_summary" || r.Header.Get("Authorization") != "Bearer sk_seo_test" {
			t.Errorf("unexpected request %s %s", r.URL.Path, r.Header.Get("Authorization"))
		}
		body, _ := io.ReadAll(r.Body)
		if string(body) != `{"clientId":"client_1"}` {
			t.Errorf("body = %s; optional fields should be omitted", body)
		}
		_, _ = w.Write([]byte(`{"data":{"pagesCrawled":10}}`))
	})
	out, err := c.GetAuditSummary(context.Background(), GetAuditSummaryInput{ClientID: "client_1"})
	if err != nil {
		t.Fatal(err)
	}
	var summary struct{ PagesCrawled int }
	if err := json.Unmarshal(out, &summary); err != nil || summary.PagesCrawled != 10 {
		t.Fatalf("got %s, %v", out, err)
	}
}

func TestRefusalCarriesCodeNextAndRequestID(t *testing.T) {
	c := fake(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(422)
		_, _ = w.Write([]byte(`{"error":"Cannot read that property yet.","code":"search_console_no_access","next":"verify_site","requestId":"req_1"}`))
	})
	_, err := c.Call(context.Background(), "connect_search_console_service_account", map[string]any{"clientId": "c"})
	apiErr, ok := err.(*Error)
	if !ok || apiErr.Status != 422 || apiErr.Code != "search_console_no_access" || apiErr.Next != "verify_site" || apiErr.RequestID != "req_1" {
		t.Fatalf("got %#v", err)
	}
	if !IsCode(err, "search_console_no_access") {
		t.Fatal("IsCode should match")
	}
}

func TestRateLimitIsRetriedButAWriteIsNotRetriedOnServerError(t *testing.T) {
	var calls atomic.Int32
	c := fake(t, func(w http.ResponseWriter, r *http.Request) {
		n := calls.Add(1)
		if r.URL.Path == "/api/v1/tools/get_next_steps" && n == 1 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(429)
			return
		}
		if r.URL.Path == "/api/v1/tools/create_client" {
			w.WriteHeader(500)
			return
		}
		_, _ = w.Write([]byte(`{"data":{}}`))
	})
	if _, err := c.GetNextSteps(context.Background(), GetNextStepsInput{ClientID: "c"}); err != nil {
		t.Fatalf("429 should be retried: %v", err)
	}
	calls.Store(0)
	_, err := c.CreateClient(context.Background(), CreateClientInput{Name: "A", Domain: "a.com"})
	if err == nil || calls.Load() != 1 {
		t.Fatalf("a write must not be repeated after a 500: calls=%d err=%v", calls.Load(), err)
	}
}

func TestOptionalNumbersArePointers(t *testing.T) {
	c := fake(t, func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		if string(body) != `{"clientId":"c","ruleId":"missing-title","limit":5}` {
			t.Errorf("body = %s", body)
		}
		_, _ = w.Write([]byte(`{"data":{}}`))
	})
	_, err := c.ListIssuePages(context.Background(), ListIssuePagesInput{ClientID: "c", RuleID: "missing-title", Limit: Ptr(int64(5))})
	if err != nil {
		t.Fatal(err)
	}
}
