package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type MockImageDownloader struct{}

func (d MockImageDownloader) DownloadImage(url, dir, filename string) error {
	return nil
}

func TestRun(t *testing.T) {
	downloader := MockImageDownloader{}

	os.Unsetenv("NASA_KEY")
	err := run(downloader)
	if err == nil || err.Error() != "NASA_KEY environment variable not set" {
		t.Errorf("expected error 'NASA_KEY environment variable not set', got %v", err)
	}

	os.Setenv("NASA_KEY", "DEMO_KEY")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"date":"2023-10-01","explanation":"Test explanation","media_type":"image","service_version":"v1","title":"Test Title","url":"https://example.com/test.jpg"}`))
	}))
	defer server.Close()

	err = run(downloader)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
