package tcgcollector

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCards(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/api/cards" {
			t.Errorf("Expected path /api/cards, got %s", r.URL.Path)
		}

		// Verify query parameters
		query := r.URL.Query()
		if query.Get("setId") != "1" {
			t.Errorf("Expected setId=1, got %s", query.Get("setId"))
		}
		if query.Get("name") != "test" {
			t.Errorf("Expected name=test, got %s", query.Get("name"))
		}
		if query.Get("number") != "001" {
			t.Errorf("Expected number=001, got %s", query.Get("number"))
		}
		if query.Get("rarity") != "rare" {
			t.Errorf("Expected rarity=rare, got %s", query.Get("rarity"))
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
					"setId": 1,
					"name": "Test Card",
					"number": "001",
					"rarity": "rare",
					"imageUrl": "http://example.com/image.jpg",
					"createdAt": "2024-01-01T00:00:00Z",
					"updatedAt": "2024-01-01T00:00:00Z",
					"description": "Test card description"
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
	setID := 1
	name := "test"
	number := "001"
	rarity := "rare"
	page := 1
	pageSize := 10

	params := &ListCardsParams{
		SetID:    &setID,
		Name:     &name,
		Number:   &number,
		Rarity:   &rarity,
		Page:     &page,
		PageSize: &pageSize,
	}

	// Make request
	result, err := client.ListCards(context.Background(), params)
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

func TestGetCard(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/api/cards/1" {
			t.Errorf("Expected path /api/cards/1, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": 1,
			"setId": 1,
			"name": "Test Card",
			"number": "001",
			"rarity": "rare",
			"imageUrl": "http://example.com/image.jpg",
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z",
			"description": "Test card description"
		}`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Make request
	result, err := client.GetCard(context.Background(), 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", result.ID)
	}
	if result.SetID != 1 {
		t.Errorf("Expected SetID to be 1, got %d", result.SetID)
	}
	if result.Name != "Test Card" {
		t.Errorf("Expected Name to be 'Test Card', got %s", result.Name)
	}
}

func TestGetCardPrices(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/api/cards/1/prices" {
			t.Errorf("Expected path /api/cards/1/prices, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"items": [
				{
					"id": 1,
					"cardId": 1,
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
	result, err := client.GetCardPrices(context.Background(), 1)
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

func TestRecalculateCachedValues(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/cards/recalculate-cached-values", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	err := client.RecalculateCachedValues(context.Background())
	assert.NoError(t, err)
}

func TestRegenerateSlugs(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/cards/regenerate-slugs", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	err := client.RegenerateSlugs(context.Background())
	assert.NoError(t, err)
}

func TestRegenerateSurrogateNumbersAndFullNames(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/cards/regenerate-surrogate-numbers-and-full-names", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	err := client.RegenerateSurrogateNumbersAndFullNames(context.Background())
	assert.NoError(t, err)
}
