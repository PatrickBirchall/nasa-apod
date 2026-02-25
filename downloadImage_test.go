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
		if _, err := w.Write([]byte("test image content")); err != nil {
			t.Fatalf("failed to write test response: %v", err)
		}
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
	defer func() {
		if cerr := file.Close(); cerr != nil {
			t.Fatalf("expected no error closing file, got %v", cerr)
		}
	}()

	content, err := io.ReadAll(file)
	if err != nil {
		t.Errorf("expected no error reading file, got %v", err)
	}

	expectedContent := "test image content"
	if string(content) != expectedContent {
		t.Errorf("expected file content %s, got %s", expectedContent, string(content))
	}
}
