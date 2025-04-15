package tcgcollector

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient("test-api-key")
	if client == nil {
		t.Fatal("Expected client to be created")
	}
	if client.apiKey != "test-api-key" {
		t.Errorf("Expected apiKey to be 'test-api-key', got %s", client.apiKey)
	}
}

func TestWithBaseURL(t *testing.T) {
	client := NewClient("test-api-key", WithBaseURL("http://localhost:8080"))
	if client.baseURL.String() != "http://localhost:8080" {
		t.Errorf("Expected baseURL to be 'http://localhost:8080', got %s", client.baseURL.String())
	}
}

func TestWithHTTPClient(t *testing.T) {
	customClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	client := NewClient("test-api-key", WithHTTPClient(customClient))
	if client.httpClient != customClient {
		t.Error("Expected httpClient to be set to custom client")
	}
}

func TestListAuditLogEventTypes(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/api/audit-log-event-types" {
			t.Errorf("Expected path /api/audit-log-event-types, got %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-api-key" {
			t.Errorf("Expected Authorization header to be 'Bearer test-api-key', got %s", r.Header.Get("Authorization"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"items": [
				{
					"id": 1,
					"codeName": "test",
					"name": "Test Event"
				}
			],
			"itemCount": 1,
			"totalItemCount": 1,
			"page": 1,
			"pageCount": 1
		}`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Make request
	result, err := client.ListAuditLogEventTypes(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result.Items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(result.Items))
	}

	if result.Items[0].ID != 1 {
		t.Errorf("Expected item ID to be 1, got %d", result.Items[0].ID)
	}
}

func TestGetAuditLogEventType(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/api/audit-log-event-types/1" {
			t.Errorf("Expected path /api/audit-log-event-types/1, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": 1,
			"codeName": "test",
			"name": "Test Event"
		}`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Make request
	result, err := client.GetAuditLogEventType(context.Background(), 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", result.ID)
	}
}

func TestErrorResponse(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{
			"message": "Invalid request",
			"code": "INVALID_REQUEST"
		}`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Make request
	_, err := client.GetAuditLogEventType(context.Background(), 1)
	if err == nil {
		t.Fatal("Expected an error")
	}

	expectedErr := "API error: Invalid request (code: INVALID_REQUEST)"
	if err.Error() != expectedErr {
		t.Errorf("Expected error '%s', got '%s'", expectedErr, err.Error())
	}
}

func TestDoRequestErrors(t *testing.T) {
	client := NewClient("test-api-key")

	// Test invalid request body marshaling
	t.Run("invalid request body", func(t *testing.T) {
		ctx := context.Background()
		body := make(chan int) // channels cannot be marshaled to JSON
		err := client.doRequest(ctx, http.MethodPost, "/test", body, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to marshal request body")
	})

	// Test invalid path parsing
	t.Run("invalid path", func(t *testing.T) {
		ctx := context.Background()
		err := client.doRequest(ctx, http.MethodGet, ":\\invalid", nil, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to parse path")
	})

	// Test request creation failure
	t.Run("invalid method", func(t *testing.T) {
		ctx := context.Background()
		err := client.doRequest(ctx, "\n", "/test", nil, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create request")
	})

	// Test HTTP client error
	t.Run("http client error", func(t *testing.T) {
		client := NewClient("test-api-key", WithHTTPClient(&http.Client{
			Transport: &errorRoundTripper{},
		}))
		ctx := context.Background()
		err := client.doRequest(ctx, http.MethodGet, "/test", nil, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to send request")
	})

	// Test error response decoding failure
	t.Run("error response decode failure", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid json"))
		}))
		defer ts.Close()

		client := NewClient("test-api-key", WithBaseURL(ts.URL))
		ctx := context.Background()
		err := client.doRequest(ctx, http.MethodGet, "/test", nil, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to decode error response")
	})

	// Test response decoding failure
	t.Run("response decode failure", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("invalid json"))
		}))
		defer ts.Close()

		client := NewClient("test-api-key", WithBaseURL(ts.URL))
		ctx := context.Background()
		var result struct{ Field string }
		err := client.doRequest(ctx, http.MethodGet, "/test", nil, &result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to decode response")
	})
}

func TestWithBaseURLPanic(t *testing.T) {
	assert.Panics(t, func() {
		WithBaseURL(":\\invalid")(NewClient("test-api-key"))
	})
}

func TestWithHTTPClientNil(t *testing.T) {
	client := NewClient("test-api-key", WithHTTPClient(nil))
	assert.Nil(t, client.httpClient)
}

// errorRoundTripper is a mock http.RoundTripper that always returns an error
type errorRoundTripper struct{}

func (e *errorRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("mock transport error")
}

func TestNewClientInvalidBaseURL(t *testing.T) {
	// Test that NewClient handles invalid default base URL
	// This is a bit tricky since we can't modify the defaultBaseURL constant
	// But we can test that WithBaseURL handles invalid URLs
	assert.Panics(t, func() {
		WithBaseURL(":\\invalid")(NewClient("test-api-key"))
	})
}

func TestDoRequestContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	ctx, cancel := context.WithCancel(context.Background())

	// Cancel the context immediately
	cancel()

	err := client.doRequest(ctx, http.MethodGet, "/test", nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context canceled")
}

func TestDoRequestContextTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	err := client.doRequest(ctx, http.MethodGet, "/test", nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context deadline exceeded")
}

func TestClientWithCustomTimeout(t *testing.T) {
	customTimeout := 5 * time.Second
	client := NewClient("test-api-key", WithHTTPClient(&http.Client{
		Timeout: customTimeout,
	}))

	assert.Equal(t, customTimeout, client.httpClient.Timeout)
}

func TestClientWithInvalidRequestBody(t *testing.T) {
	client := NewClient("test-api-key")
	invalidBody := make(chan int)
	err := client.doRequest(context.Background(), http.MethodPost, "/test", invalidBody, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to marshal request body")
}

func TestClientWithInvalidPath(t *testing.T) {
	client := NewClient("test-api-key")
	err := client.doRequest(context.Background(), http.MethodGet, ":\\invalid", nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse path")
}

func TestClientWithInvalidMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Invalid method", "code": "INVALID_METHOD"}`))
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	err := client.doRequest(context.Background(), "INVALID", "/test", nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "API error: Invalid method")
}

func TestClientWithEmptyAPIKey(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "API key is required", "code": "UNAUTHORIZED"}`))
	}))
	defer server.Close()

	client := NewClient("", WithBaseURL(server.URL))
	err := client.doRequest(context.Background(), http.MethodGet, "/test", nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "API error: API key is required")
}

func TestClientWithNilResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data": "test"}`))
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	err := client.doRequest(context.Background(), http.MethodGet, "/test", nil, nil)
	assert.NoError(t, err)
}

func TestClientWithCustomBaseURL(t *testing.T) {
	customURL := "http://custom.example.com"
	client := NewClient("test-api-key", WithBaseURL(customURL))
	assert.Equal(t, customURL, client.baseURL.String())
}

func TestClientWithInvalidBaseURLPanic(t *testing.T) {
	assert.Panics(t, func() {
		NewClient("test-api-key", WithBaseURL(":\\invalid"))
	})
}

func TestClientWithMultipleOptions(t *testing.T) {
	customURL := "http://custom.example.com"
	customTimeout := 10 * time.Second
	client := NewClient("test-api-key",
		WithBaseURL(customURL),
		WithHTTPClient(&http.Client{Timeout: customTimeout}),
	)
	assert.Equal(t, customURL, client.baseURL.String())
	assert.Equal(t, customTimeout, client.httpClient.Timeout)
}

func TestClientWithNilHTTPClient(t *testing.T) {
	client := NewClient("test-api-key", WithHTTPClient(nil))
	assert.NotNil(t, client)
	assert.Nil(t, client.httpClient)
}

func TestClientWithInvalidResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"invalid json`))
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	var result struct{ Field string }
	err := client.doRequest(context.Background(), http.MethodGet, "/test", nil, &result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to decode response")
}

func TestClientWithInvalidErrorResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"invalid json`))
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	err := client.doRequest(context.Background(), http.MethodGet, "/test", nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to decode error response")
}

func TestClientWithRequestCreationError(t *testing.T) {
	client := NewClient("test-api-key")
	err := client.doRequest(context.Background(), string([]byte{0x7f}), "/test", nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create request")
}

func TestClientWithHTTPError(t *testing.T) {
	client := NewClient("test-api-key", WithHTTPClient(&http.Client{
		Transport: &errorRoundTripper{},
	}))
	err := client.doRequest(context.Background(), http.MethodGet, "/test", nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to send request")
}

func TestClientWithDefaultTimeout(t *testing.T) {
	client := NewClient("test-api-key")
	assert.Equal(t, defaultTimeout, client.httpClient.Timeout)
}

func TestClientWithRelativeURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/test", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	err := client.doRequest(context.Background(), http.MethodGet, "/api/test", nil, nil)
	assert.NoError(t, err)
}

func TestClientWithQueryParameters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "value", r.URL.Query().Get("param"))
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	err := client.doRequest(context.Background(), http.MethodGet, "/test?param=value", nil, nil)
	assert.NoError(t, err)
}

func TestClientWithHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	err := client.doRequest(context.Background(), http.MethodGet, "/test", nil, nil)
	assert.NoError(t, err)
}

func TestClientWithRequestBody(t *testing.T) {
	expectedBody := map[string]string{"key": "value"}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]string
		err := json.NewDecoder(r.Body).Decode(&body)
		assert.NoError(t, err)
		assert.Equal(t, expectedBody, body)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	err := client.doRequest(context.Background(), http.MethodPost, "/test", expectedBody, nil)
	assert.NoError(t, err)
}

func TestClientWithResponseBody(t *testing.T) {
	expectedResponse := map[string]string{"key": "value"}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	var response map[string]string
	err := client.doRequest(context.Background(), http.MethodGet, "/test", nil, &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}

func TestClientWithDefaultBaseURL(t *testing.T) {
	client := NewClient("test-api-key")
	assert.Equal(t, defaultBaseURL, client.baseURL.String())
}

func TestClientWithEmptyBaseURL(t *testing.T) {
	assert.Panics(t, func() {
		NewClient("test-api-key", WithBaseURL(""))
	})
}

func TestClientWithInvalidBaseURLScheme(t *testing.T) {
	assert.Panics(t, func() {
		NewClient("test-api-key", WithBaseURL("invalid://example.com"))
	})
}
