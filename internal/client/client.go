package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	Host       string
	APIKey     string
	HTTPClient *http.Client
}

type APIError struct {
	StatusCode int
	Message    string
	Code       string
}

func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("API Error (%d, %s): %s", e.StatusCode, e.Code, e.Message)
	}
	return fmt.Sprintf("API Error (%d): %s", e.StatusCode, e.Message)
}

func IsNotFound(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == http.StatusNotFound
	}
	return false
}

func NewClient(host, apiKey string, insecure bool) (*Client, error) {
	if host == "" {
		return nil, fmt.Errorf("host is required")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("api_key is required")
	}

	// Normalize host URL
	if !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
		host = "https://" + host
	}
	host = strings.TrimSuffix(host, "/")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: insecure,
		},
	}

	httpClient := &http.Client{
		Transport: tr,
		Timeout:   30 * time.Second,
	}

	return &Client{
		Host:       host,
		APIKey:     apiKey,
		HTTPClient: httpClient,
	}, nil
}

func (c *Client) DoRequest(ctx context.Context, method, path string, body interface{}, out interface{}) error {
	fullURL := c.Host + path

	var reqBody io.Reader
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(buf)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, reqBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-API-Key", c.APIKey)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	respBuf, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errResp struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}
		_ = json.Unmarshal(respBuf, &errResp)
		msg := errResp.Message
		if msg == "" {
			msg = string(respBuf)
		}
		return &APIError{
			StatusCode: resp.StatusCode,
			Code:       errResp.Code,
			Message:    msg,
		}
	}

	if out != nil && len(respBuf) > 0 {
		if err := json.Unmarshal(respBuf, out); err != nil {
			return fmt.Errorf("failed to decode response JSON: %w", err)
		}
	}

	return nil
}

// Helper to format query parameters
func FormatQueryParams(offset, limit int, filter string) string {
	v := url.Values{}
	if offset > 0 {
		v.Set("offset", fmt.Sprintf("%d", offset))
	}
	if limit > 0 {
		v.Set("limit", fmt.Sprintf("%d", limit))
	}
	if filter != "" {
		v.Set("filter", filter)
	}
	if len(v) == 0 {
		return ""
	}
	return "?" + v.Encode()
}
