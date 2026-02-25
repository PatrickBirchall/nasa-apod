package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type ImageDownloader interface {
	DownloadImage(url, dir, filename string) error
}

type RealImageDownloader struct{}

// DownloadImage downloads an image from a URL and saves it to a file in a subdirectory
func (d RealImageDownloader) DownloadImage(url string, subdirectory string, filename string) (err error) {
	if err := os.MkdirAll(subdirectory, 0o755); err != nil {
		return fmt.Errorf("creating directory %q: %w", subdirectory, err)
	}

	path := filepath.Join(subdirectory, filename)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("downloading image from %q: %w", url, err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("closing response body: %w", cerr)
		}
	}()

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("creating file %q: %w", path, err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("closing file %q: %w", path, cerr)
		}
	}()

	if _, err = io.Copy(file, resp.Body); err != nil {
		return fmt.Errorf("writing image to %q: %w", path, err)
	}

	return nil
}
