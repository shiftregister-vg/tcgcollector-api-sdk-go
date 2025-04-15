package tcgcollector

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestListAuditLogEntries(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/api/audit-log" {
			t.Errorf("Expected path /api/audit-log, got %s", r.URL.Path)
		}

		// Verify query parameters
		query := r.URL.Query()
		if query.Get("eventTypeId") != "1" {
			t.Errorf("Expected eventTypeId=1, got %s", query.Get("eventTypeId"))
		}
		if query.Get("userId") != "2" {
			t.Errorf("Expected userId=2, got %s", query.Get("userId"))
		}
		if query.Get("startDate") != "2024-01-01T00:00:00Z" {
			t.Errorf("Expected startDate=2024-01-01T00:00:00Z, got %s", query.Get("startDate"))
		}
		if query.Get("endDate") != "2024-01-31T23:59:59Z" {
			t.Errorf("Expected endDate=2024-01-31T23:59:59Z, got %s", query.Get("endDate"))
		}
		if query.Get("page") != "1" {
			t.Errorf("Expected page=1, got %s", query.Get("page"))
		}
		if query.Get("pageSize") != "10" {
			t.Errorf("Expected pageSize=10, got %s", query.Get("pageSize"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"items": [
				{
					"id": 1,
					"eventTypeId": 1,
					"userId": 2,
					"ipAddress": "127.0.0.1",
					"createdAt": "2024-01-15T12:00:00Z",
					"details": "Test audit log entry"
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

	// Create parameters
	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 1, 31, 23, 59, 59, 0, time.UTC)
	eventTypeID := 1
	userID := 2
	page := 1
	pageSize := 10

	params := &ListAuditLogEntriesParams{
		EventTypeID: &eventTypeID,
		UserID:      &userID,
		StartDate:   &startDate,
		EndDate:     &endDate,
		Page:        &page,
		PageSize:    &pageSize,
	}

	// Make request
	result, err := client.ListAuditLogEntries(context.Background(), params)
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

func TestGetAuditLogEntry(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/api/audit-log/1" {
			t.Errorf("Expected path /api/audit-log/1, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": 1,
			"eventTypeId": 1,
			"userId": 2,
			"ipAddress": "127.0.0.1",
			"createdAt": "2024-01-15T12:00:00Z",
			"details": "Test audit log entry"
		}`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Make request
	result, err := client.GetAuditLogEntry(context.Background(), 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", result.ID)
	}
	if result.EventTypeID != 1 {
		t.Errorf("Expected EventTypeID to be 1, got %d", result.EventTypeID)
	}
	if result.UserID != 2 {
		t.Errorf("Expected UserID to be 2, got %d", result.UserID)
	}
	if result.IPAddress != "127.0.0.1" {
		t.Errorf("Expected IPAddress to be '127.0.0.1', got %s", result.IPAddress)
	}
}
