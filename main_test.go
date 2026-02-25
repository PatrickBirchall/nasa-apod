package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockImageDownloader struct{}

func (d MockImageDownloader) DownloadImage(url, dir, filename string) error {
	return nil
}

func TestRun(t *testing.T) {
	downloader := MockImageDownloader{}

	cfg := Config{
		APIKey:    "",
		BaseURL:   defaultBaseURL,
		OutputDir: "images",
	}

	err := run(cfg, downloader)
	if !errors.Is(err, ErrMissingNASAKey) {
		t.Errorf("expected ErrMissingNASAKey, got %v", err)
	}

	cfg = Config{
		APIKey:    "DEMO_KEY",
		BaseURL:   defaultBaseURL,
		OutputDir: "images",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"date":"2023-10-01","explanation":"Test explanation","media_type":"image","service_version":"v1","title":"Test Title","url":"https://example.com/test.jpg"}`))
	}))
	defer server.Close()

	cfg.BaseURL = server.URL

	err = run(cfg, downloader)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
