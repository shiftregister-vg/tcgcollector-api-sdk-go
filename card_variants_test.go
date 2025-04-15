package tcgcollector

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCardVariants(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/api/card-variants" {
			t.Errorf("Expected path /api/card-variants, got %s", r.URL.Path)
		}

		// Verify query parameters
		query := r.URL.Query()
		if query.Get("cardId") != "1" {
			t.Errorf("Expected cardId=1, got %s", query.Get("cardId"))
		}
		if query.Get("typeId") != "2" {
			t.Errorf("Expected typeId=2, got %s", query.Get("typeId"))
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
					"cardId": 1,
					"typeId": 2,
					"name": "Test Variant",
					"description": "Test variant description",
					"imageUrl": "http://example.com/image.jpg",
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
	cardID := 1
	typeID := 2
	page := 1
	pageSize := 10

	params := &ListCardVariantsParams{
		CardID:   &cardID,
		TypeID:   &typeID,
		Page:     &page,
		PageSize: &pageSize,
	}

	// Make request
	result, err := client.ListCardVariants(context.Background(), params)
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

func TestGetCardVariant(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/api/card-variants/1" {
			t.Errorf("Expected path /api/card-variants/1, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": 1,
			"cardId": 1,
			"typeId": 2,
			"name": "Test Variant",
			"description": "Test variant description",
			"imageUrl": "http://example.com/image.jpg",
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Make request
	result, err := client.GetCardVariant(context.Background(), 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", result.ID)
	}
	if result.CardID != 1 {
		t.Errorf("Expected CardID to be 1, got %d", result.CardID)
	}
	if result.TypeID != 2 {
		t.Errorf("Expected TypeID to be 2, got %d", result.TypeID)
	}
}

func TestCreateCardVariant(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/api/card-variants" {
			t.Errorf("Expected path /api/card-variants, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": 1,
			"cardId": 1,
			"typeId": 2,
			"name": "Test Variant",
			"description": "Test variant description",
			"imageUrl": "http://example.com/image.jpg",
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Create variant
	variant := &CardVariant{
		CardID:      1,
		TypeID:      2,
		Name:        "Test Variant",
		Description: "Test variant description",
		ImageURL:    "http://example.com/image.jpg",
	}

	// Make request
	result, err := client.CreateCardVariant(context.Background(), variant)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", result.ID)
	}
}

func TestUpdateCardVariant(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}
		if r.URL.Path != "/api/card-variants/1" {
			t.Errorf("Expected path /api/card-variants/1, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": 1,
			"cardId": 1,
			"typeId": 2,
			"name": "Updated Variant",
			"description": "Updated variant description",
			"imageUrl": "http://example.com/image.jpg",
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Create variant
	variant := &CardVariant{
		CardID:      1,
		TypeID:      2,
		Name:        "Updated Variant",
		Description: "Updated variant description",
		ImageURL:    "http://example.com/image.jpg",
	}

	// Make request
	result, err := client.UpdateCardVariant(context.Background(), 1, variant)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", result.ID)
	}
	if result.Name != "Updated Variant" {
		t.Errorf("Expected Name to be 'Updated Variant', got %s", result.Name)
	}
}

func TestDeleteCardVariant(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}
		if r.URL.Path != "/api/card-variants/1" {
			t.Errorf("Expected path /api/card-variants/1, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Make request
	err := client.DeleteCardVariant(context.Background(), 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetCardVariantPrices(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/api/card-variants/1/prices" {
			t.Errorf("Expected path /api/card-variants/1/prices, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"items": [
				{
					"id": 1,
					"variantId": 1,
					"price": 10.99,
					"currency": "USD",
					"source": "TCGPlayer",
					"createdAt": "2024-01-01T00:00:00Z"
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
	result, err := client.GetCardVariantPrices(context.Background(), 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result.Items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(result.Items))
	}

	if result.Items[0].ID != 1 {
		t.Errorf("Expected item ID to be 1, got %d", result.Items[0].ID)
	}
	if result.Items[0].Price != 10.99 {
		t.Errorf("Expected price to be 10.99, got %f", result.Items[0].Price)
	}
}

func TestRecalculateComputedAndCachedValues(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/card-variants/recalculate-computed-and-cached-values", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	err := client.RecalculateComputedAndCachedValues(context.Background())
	assert.NoError(t, err)
}
