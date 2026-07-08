// Package docuseal provides a client for the DocuSeal API.
//
// API documentation: https://www.docuseal.com/docs/api
package docuseal

//go:generate ./scripts/generate-types.sh

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	GlobalURL = "https://api.docuseal.com"
	EuURL     = "https://api.docuseal.eu"

	version = "0.1.0"
)

type Client struct {
	key        string
	baseURL    string
	userAgent  string
	httpClient *http.Client
}

type Option func(*Client)

// WithBaseURL sets a custom API base URL. Use EuURL for the EU cloud or
// your own URL for on-premises installations, e.g. "https://yourdocuseal.com/api".
func WithBaseURL(baseURL string) Option {
	return func(c *Client) { c.baseURL = baseURL }
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) { c.httpClient = httpClient }
}

func WithUserAgent(userAgent string) Option {
	return func(c *Client) { c.userAgent = userAgent + " docuseal-go/" + version }
}

// NewClient returns a Client authenticated with the given API key.
// Get your API key at https://console.docuseal.com/api.
func NewClient(key string, opts ...Option) *Client {
	c := &Client{
		key:        key,
		baseURL:    GlobalURL,
		userAgent:  "docuseal-go/" + version,
		httpClient: &http.Client{Timeout: 60 * time.Second},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

type APIError struct {
	StatusCode int
	Message    string
	Body       []byte
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("docuseal: %d %s", e.StatusCode, e.Message)
	}

	return fmt.Sprintf("docuseal: unexpected status %d", e.StatusCode)
}

func (c *Client) do(ctx context.Context, method, path string, query url.Values, body, out any) error {
	endpoint := c.baseURL + path
	if len(query) > 0 {
		endpoint += "?" + query.Encode()
	}

	var payload io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}
		payload = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint, payload)
	if err != nil {
		return err
	}

	req.Header.Set("X-Auth-Token", c.key)
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		var apiErr struct {
			Error string `json:"error"`
		}
		_ = json.Unmarshal(data, &apiErr)

		return &APIError{StatusCode: resp.StatusCode, Message: apiErr.Error, Body: data}
	}

	if out == nil {
		return nil
	}

	return json.Unmarshal(data, out)
}

func queryValues(params any) url.Values {
	values := url.Values{}
	if params == nil {
		return values
	}

	data, err := json.Marshal(params)
	if err != nil {
		return values
	}

	var fields map[string]any
	if err := json.Unmarshal(data, &fields); err != nil {
		return values
	}

	for key, value := range fields {
		values.Set(key, fmt.Sprintf("%v", value))
	}

	return values
}
