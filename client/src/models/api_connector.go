package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// APIConnector :
// Handles communication with the local API
type APIConnector struct {
	baseURL    string
	httpClient *http.Client
}

// NewAPIConnector :
// Creates a new instance of APIConnector. The default timeout is 10 seconds.
func NewAPIConnector(baseURL string) *APIConnector {
	return &APIConnector{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// SendPackage :
// Sends the PackageSent structure to the specified endpoint
func (c *APIConnector) SendPackage(endpoint string, pkg PackageSent) (*http.Response, error) {
	// Marshal the package to JSON
	jsonData, err := json.Marshal(pkg)
	if err != nil {
		return nil, fmt.Errorf("error marshaling package: %v", err)
	}

	// Create the full URL
	fullURL := c.baseURL + endpoint

	// Create a new request
	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	return resp, nil
}
