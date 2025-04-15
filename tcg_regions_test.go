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

func TestListTCGRegions(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check request method and path
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/tcg-regions", r.URL.Path)

		// Check query parameters
		page := r.URL.Query().Get("page")
		pageSize := r.URL.Query().Get("pageSize")
		assert.Equal(t, "1", page)
		assert.Equal(t, "10", pageSize)

		// Create mock response
		response := ListTCGRegionsResponse{
			Items: []TCGRegion{
				{
					ID:          1,
					Name:        "North America",
					Description: "North American region",
					Code:        "NA",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
				{
					ID:          2,
					Name:        "Europe",
					Description: "European region",
					Code:        "EU",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			},
			ItemCount:      2,
			TotalItemCount: 2,
			Page:           1,
			PageCount:      1,
		}

		// Send response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create client with mock server URL
	client := NewClient("test-api-key", WithBaseURL(server.URL))

	// Set up test parameters
	params := &ListTCGRegionsParams{
		Page:     intPtr(1),
		PageSize: intPtr(10),
	}

	// Call the function
	response, err := client.ListTCGRegions(context.Background(), params)

	// Check response
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 2, len(response.Items))
	assert.Equal(t, "North America", response.Items[0].Name)
	assert.Equal(t, "Europe", response.Items[1].Name)
}

func TestGetTCGRegion(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check request method and path
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/tcg-regions/1", r.URL.Path)

		// Create mock response
		response := TCGRegion{
			ID:          1,
			Name:        "North America",
			Description: "North American region",
			Code:        "NA",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		// Send response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create client with mock server URL
	client := NewClient("test-api-key", WithBaseURL(server.URL))

	// Call the function
	response, err := client.GetTCGRegion(context.Background(), 1)

	// Check response
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "North America", response.Name)
	assert.Equal(t, "NA", response.Code)
}

// Helper function to create integer pointers
func intPtr(i int) *int {
	return &i
}
