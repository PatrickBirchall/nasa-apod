package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

// FetchAPOD fetches the Astronomy Picture of the Day (APOD) from the NASA API
func FetchAPOD() (Response, error) {
	apiKey := os.Getenv("NASA_KEY")
	url := "https://api.nasa.gov/planetary/apod?api_key=" + apiKey

	resp, err := http.Get(url)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return Response{}, err
	}

	return result, nil
}
