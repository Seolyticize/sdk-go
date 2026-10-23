package site

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAnalyticsTag(t *testing.T) {
	tag := AnalyticsTag("abc123", TagOptions{Exclude: []string{"/app/*", "/verify"}, Track: []string{"outbound"}})
	for _, want := range []string{`data-site="abc123"`, `data-exclude="/app/*,/verify"`, `data-track="outbound"`, `src="https://seolyticize.com/a.js"`, "defer"} {
		if !strings.Contains(tag, want) {
			t.Errorf("tag %q lacks %q", tag, want)
		}
	}
	if strings.Contains(AnalyticsTag(`"><x`, TagOptions{}), `"><x`) {
		t.Error("site key must be escaped")
	}
}

func TestIndexNowKeyHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	IndexNowKeyHandler("k123").ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/k123.txt", nil))
	if recorder.Body.String() != "k123" || !strings.HasPrefix(recorder.Header().Get("Content-Type"), "text/plain") {
		t.Fatalf("got %q %q", recorder.Body.String(), recorder.Header().Get("Content-Type"))
	}
}

func TestNotifyDeploy(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string][]string
		_ = json.NewDecoder(r.Body).Decode(&body)
		_ = json.NewEncoder(w).Encode(map[string]any{"submitted": len(body["urls"]), "result": "ok"})
	}))
	defer server.Close()
	result, err := NotifyDeploy(context.Background(), server.URL, []string{"https://a.com/x"})
	if err != nil || result.Submitted != 1 {
		t.Fatalf("got %+v %v", result, err)
	}
}
