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

func TestListTCGPriceSources(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/tcg-price-sources", r.URL.Path)

		// Check query parameters
		query := r.URL.Query()
		assert.Equal(t, "1", query.Get("page"))
		assert.Equal(t, "10", query.Get("pageSize"))

		// Create mock response
		response := ListResponse[TCGPriceSource]{
			Items: []TCGPriceSource{
				{
					ID:          1,
					Name:        "Test Source",
					Description: "Test price source description",
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

	params := &ListTCGPriceSourcesParams{
		Page:     &page,
		PageSize: &pageSize,
	}

	result, err := client.ListTCGPriceSources(context.Background(), params)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result.Items))
	assert.Equal(t, 1, result.Items[0].ID)
	assert.Equal(t, "Test Source", result.Items[0].Name)
	assert.Equal(t, "Test price source description", result.Items[0].Description)
}

func TestGetTCGPriceSource(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/tcg-price-sources/1", r.URL.Path)

		// Create mock response
		response := TCGPriceSource{
			ID:          1,
			Name:        "Test Source",
			Description: "Test price source description",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	result, err := client.GetTCGPriceSource(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test Source", result.Name)
	assert.Equal(t, "Test price source description", result.Description)
}
