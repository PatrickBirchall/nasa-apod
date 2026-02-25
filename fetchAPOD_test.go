package main

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func MockResponse(statusCode int, body string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestFetchAPOD(t *testing.T) {
	cfg := Config{
		APIKey:    "",
		BaseURL:   defaultBaseURL,
		OutputDir: "images",
	}

	_, err := FetchAPOD(cfg)
	if !errors.Is(err, ErrMissingNASAKey) {
		t.Errorf("expected ErrMissingNASAKey, got %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"date":"2023-10-01","explanation":"Test explanation","media_type":"image","service_version":"v1","title":"Test Title","url":"https://example.com/test.jpg"}`))
	}))
	defer server.Close()

	cfg = Config{
		APIKey:    "DEMO_KEY",
		BaseURL:   server.URL,
		OutputDir: "images",
	}

	resp, err := FetchAPOD(cfg)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if resp.Title != "Test Title" {
		t.Errorf("expected title 'Test Title', got %v", resp.Title)
	}
}
