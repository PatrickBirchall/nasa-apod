package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestDownloadImage(t *testing.T) {
	downloader := RealImageDownloader{}

	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Test case: Successful download
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test image content"))
	}))
	defer server.Close()

	filename := "test_image.jpg"
	err := downloader.DownloadImage(server.URL, tempDir, filename)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Check if the file was created
	filePath := filepath.Join(tempDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("expected file %s to be created, but it does not exist", filePath)
	}

	// Check the file content
	file, err := os.Open(filePath)
	if err != nil {
		t.Errorf("expected no error opening file, got %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		t.Errorf("expected no error reading file, got %v", err)
	}

	expectedContent := "test image content"
	if string(content) != expectedContent {
		t.Errorf("expected file content %s, got %s", expectedContent, string(content))
	}
}
