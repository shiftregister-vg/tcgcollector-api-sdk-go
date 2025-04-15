package tcgcollector

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
