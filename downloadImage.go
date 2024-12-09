package main

import (
	"io"
	"net/http"
	"os"
)

type ImageDownloader interface {
	DownloadImage(url, dir, filename string) error
}

type RealImageDownloader struct{}

// DownloadImage downloads an image from a URL and saves it to a file in a subdirectory
func (d RealImageDownloader) DownloadImage(url string, subdirectory string, filename string) error {
	// Create the subdirectory if it doesn't exist
	if _, err := os.Stat(subdirectory); os.IsNotExist(err) {
		os.Mkdir(subdirectory, 0755)
	}

	// Prepend the subdirectory to the filename
	path := subdirectory + "/" + filename

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}
