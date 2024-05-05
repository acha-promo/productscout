package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HttpClient struct holds any configuration and the http.Client instance
type HttpClient struct {
	client    *http.Client
	userAgent string
}

// Option configures the HttpClient
type Option func(*HttpClient)

// WithTimeout sets the timeout for the HttpClient
func WithTimeout(d time.Duration) Option {
	return func(hc *HttpClient) {
		hc.client.Timeout = d
	}
}

// WithUserAgent sets the userAgent for the HttpClient
func WithUserAgent(ua string) Option {
	return func(hc *HttpClient) {
		hc.userAgent = ua
	}
}

// NewHttpClient creates a new instance of HttpClient with optional configurations
func NewHttpClient(opts ...Option) *HttpClient {
	hc := &HttpClient{
		client: &http.Client{}, // Default http.Client without a timeout
	}
	// Apply all options to the HttpClient
	for _, opt := range opts {
		opt(hc)
	}
	return hc
}

// GetJSON performs a GET request and decodes the JSON response into the target interface{}
func (hc *HttpClient) GetJSON(url string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", hc.userAgent)

	resp, err := hc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code [%d] for the URL [%s]", resp.StatusCode, url)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, target)
}
