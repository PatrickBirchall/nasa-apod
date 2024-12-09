package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
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
	os.Unsetenv("NASA_KEY")
	_, err := FetchAPOD("https://api.nasa.gov/planetary/apod")
	if err == nil || err.Error() != "NASA_KEY environment variable not set" {
		t.Errorf("expected error 'NASA_KEY environment variable not set', got %v", err)
	}

	os.Setenv("NASA_KEY", "DEMO_KEY")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"date":"2023-10-01","explanation":"Test explanation","media_type":"image","service_version":"v1","title":"Test Title","url":"https://example.com/test.jpg"}`))
	}))
	defer server.Close()

	resp, err := FetchAPOD(server.URL)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if resp.Title != "Test Title" {
		t.Errorf("expected title 'Test Title', got %v", resp.Title)
	}
}
