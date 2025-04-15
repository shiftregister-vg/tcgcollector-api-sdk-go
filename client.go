package tcgcollector

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
	defaultBaseURL = "https://www.tcgcollector.com"
	defaultTimeout = 30 * time.Second
)

// Client represents a TCG Collector API client
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
	apiKey     string
}

// ClientOption is a function that configures a Client
type ClientOption func(*Client)

// NewClient creates a new TCG Collector API client
func NewClient(apiKey string, opts ...ClientOption) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	client := &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
		apiKey: apiKey,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// WithBaseURL sets the base URL for the client
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		if baseURL == "" {
			panic("base URL cannot be empty")
		}
		u, err := url.Parse(baseURL)
		if err != nil {
			panic(fmt.Sprintf("invalid base URL: %v", err))
		}
		if u.Scheme != "http" && u.Scheme != "https" {
			panic(fmt.Sprintf("invalid base URL scheme: %s (must be http or https)", u.Scheme))
		}
		c.baseURL = u
	}
}

// WithHTTPClient sets the HTTP client for the client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// doRequest performs an HTTP request and decodes the response
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewReader(jsonData)
	}

	// Parse the path to handle query parameters correctly
	u, err := url.Parse(path)
	if err != nil {
		return fmt.Errorf("failed to parse path: %w", err)
	}

	// Join with base URL
	reqURL := c.baseURL.ResolveReference(u)

	req, err := http.NewRequestWithContext(ctx, method, reqURL.String(), reqBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var errResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return fmt.Errorf("failed to decode error response: %w", err)
		}
		return fmt.Errorf("API error: %s (code: %s)", errResp.Message, errResp.Code)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// ListAuditLogEventTypes lists all audit log event types
func (c *Client) ListAuditLogEventTypes(ctx context.Context) (*ListResponse[AuditLogEventType], error) {
	var result ListResponse[AuditLogEventType]
	if err := c.doRequest(ctx, http.MethodGet, "/api/audit-log-event-types", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAuditLogEventType gets a single audit log event type by ID
func (c *Client) GetAuditLogEventType(ctx context.Context, id int) (*AuditLogEventType, error) {
	var result AuditLogEventType
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/audit-log-event-types/%d", id), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
