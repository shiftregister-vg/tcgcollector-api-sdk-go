package tcgcollector

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListImages(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/api/images" {
			t.Errorf("Expected path /api/images, got %s", r.URL.Path)
		}

		// Verify query parameters
		query := r.URL.Query()
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
					"url": "http://example.com/image.jpg",
					"contentType": "image/jpeg",
					"size": 1024,
					"width": 800,
					"height": 600,
					"createdAt": "2024-01-01T00:00:00Z",
					"updatedAt": "2024-01-01T00:00:00Z"
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
	page := 1
	pageSize := 10

	params := &ListImagesParams{
		Page:     &page,
		PageSize: &pageSize,
	}

	// Make request
	result, err := client.ListImages(context.Background(), params)
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

func TestGetImage(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/api/images/1" {
			t.Errorf("Expected path /api/images/1, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": 1,
			"url": "http://example.com/image.jpg",
			"contentType": "image/jpeg",
			"size": 1024,
			"width": 800,
			"height": 600,
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Make request
	result, err := client.GetImage(context.Background(), 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", result.ID)
	}
	if result.ContentType != "image/jpeg" {
		t.Errorf("Expected ContentType to be 'image/jpeg', got %s", result.ContentType)
	}
}

func TestCreateImage(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/api/images" {
			t.Errorf("Expected path /api/images, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": 1,
			"url": "http://example.com/image.jpg",
			"contentType": "image/jpeg",
			"size": 1024,
			"width": 800,
			"height": 600,
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Create image
	params := &CreateImageParams{
		File: []byte("test image data"),
	}

	// Make request
	result, err := client.CreateImage(context.Background(), params)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", result.ID)
	}
}

func TestDeleteImage(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}
		if r.URL.Path != "/api/images/1" {
			t.Errorf("Expected path /api/images/1, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Make request
	err := client.DeleteImage(context.Background(), 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
