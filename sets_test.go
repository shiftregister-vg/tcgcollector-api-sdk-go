package tcgcollector

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListSets(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/sets", r.URL.Path)

		// Check query parameters
		query := r.URL.Query()
		assert.Equal(t, "1", query.Get("page"))
		assert.Equal(t, "10", query.Get("pageSize"))

		// Create mock response
		response := ListResponse[Set]{
			Items: []Set{
				{
					ID:          1,
					Name:        "Test Set",
					Code:        "TST",
					ReleaseDate: "2024-01-01",
					TotalCards:  100,
					ImageURL:    "http://example.com/image.jpg",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			},
			ItemCount:      1,
			TotalItemCount: 1,
			Page:           1,
			PageCount:      1,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Create parameters
	page := 1
	pageSize := 10

	params := &ListSetsParams{
		Page:     &page,
		PageSize: &pageSize,
	}

	result, err := client.ListSets(context.Background(), params)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result.Items))
	assert.Equal(t, 1, result.Items[0].ID)
	assert.Equal(t, "Test Set", result.Items[0].Name)
	assert.Equal(t, "TST", result.Items[0].Code)
}

func TestGetSet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/sets/1", r.URL.Path)

		// Create mock response
		response := Set{
			ID:          1,
			Name:        "Test Set",
			Code:        "TST",
			ReleaseDate: "2024-01-01",
			TotalCards:  100,
			ImageURL:    "http://example.com/image.jpg",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	result, err := client.GetSet(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test Set", result.Name)
	assert.Equal(t, "TST", result.Code)
}

func TestGetSetCards(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/sets/1/cards", r.URL.Path)

		// Create mock response
		response := ListResponse[Card]{
			Items: []Card{
				{
					ID:        1,
					Name:      "Test Card",
					Number:    "001",
					ImageURL:  "http://example.com/image.jpg",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			ItemCount:      1,
			TotalItemCount: 1,
			Page:           1,
			PageCount:      1,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	result, err := client.GetSetCards(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result.Items))
	assert.Equal(t, 1, result.Items[0].ID)
	assert.Equal(t, "Test Card", result.Items[0].Name)
	assert.Equal(t, "001", result.Items[0].Number)
}
