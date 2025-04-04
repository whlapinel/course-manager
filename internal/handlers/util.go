package handlers

import (
	"net/url"
)

// AddQueryParams adds arbitrary query parameters to a given URL
func AddQueryParams(baseURL string, params map[string]string) (string, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	// Get the existing query params
	query := parsedURL.Query()

	// Append new params
	for key, value := range params {
		if value != "" {
			query.Set(key, value)
		}
	}
	// Encode and set the updated query string
	parsedURL.RawQuery = query.Encode()
	return parsedURL.String(), nil
}
