package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var ErrMissingNASAKey = errors.New("NASA_KEY environment variable not set")

// FetchAPOD fetches the Astronomy Picture of the Day (APOD) from the NASA API
func FetchAPOD(cfg Config) (Response, error) {
	if cfg.APIKey == "" {
		return Response{}, ErrMissingNASAKey
	}

	url := cfg.BaseURL + "?api_key=" + cfg.APIKey

	resp, err := http.Get(url)
	if err != nil {
		return Response{}, fmt.Errorf("fetching APOD: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		limitedBody, readErr := io.ReadAll(io.LimitReader(resp.Body, 512))
		if readErr != nil {
			return Response{}, fmt.Errorf("fetching APOD: unexpected status %d and failed to read response body: %w", resp.StatusCode, readErr)
		}
		return Response{}, fmt.Errorf("fetching APOD: unexpected status %d: %s", resp.StatusCode, string(limitedBody))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, fmt.Errorf("reading APOD response body: %w", err)
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return Response{}, fmt.Errorf("unmarshalling APOD response: %w", err)
	}

	return result, nil
}
