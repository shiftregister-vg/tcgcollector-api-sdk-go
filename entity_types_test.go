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

func TestListEntityTypes(t *testing.T) {
	// Test successful request with no parameters
	t.Run("success without params", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/entity-types", r.URL.Path)
			assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
			assert.Equal(t, "application/json", r.Header.Get("Accept"))

			response := ListEntityTypesResponse{
				Items: []EntityType{
					{
						ID:          1,
						Name:        "Test Entity Type",
						Description: "Test Description",
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
		defer server.Close()

		client := NewClient("test-api-key", WithBaseURL(server.URL))
		result, err := client.ListEntityTypes(context.Background(), nil)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Items, 1)
		assert.Equal(t, 1, result.Items[0].ID)
		assert.Equal(t, "Test Entity Type", result.Items[0].Name)
		assert.Equal(t, "Test Description", result.Items[0].Description)
	})

	// Test successful request with pagination parameters
	t.Run("success with pagination", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "2", r.URL.Query().Get("page"))
			assert.Equal(t, "50", r.URL.Query().Get("pageSize"))

			response := ListEntityTypesResponse{
				Items:          []EntityType{},
				ItemCount:      0,
				TotalItemCount: 100,
				Page:           2,
				PageCount:      3,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		}))
		defer server.Close()

		client := NewClient("test-api-key", WithBaseURL(server.URL))
		page := 2
		pageSize := 50
		result, err := client.ListEntityTypes(context.Background(), &ListEntityTypesParams{
			Page:     &page,
			PageSize: &pageSize,
		})
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 2, result.Page)
		assert.Equal(t, 3, result.PageCount)
		assert.Equal(t, 100, result.TotalItemCount)
	})

	// Test error response
	t.Run("error response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{
				Message: "Internal server error",
				Code:    "INTERNAL_ERROR",
			})
		}))
		defer server.Close()

		client := NewClient("test-api-key", WithBaseURL(server.URL))
		result, err := client.ListEntityTypes(context.Background(), nil)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "Internal server error")
	})

	// Test invalid response
	t.Run("invalid response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{invalid json`))
		}))
		defer server.Close()

		client := NewClient("test-api-key", WithBaseURL(server.URL))
		result, err := client.ListEntityTypes(context.Background(), nil)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to decode response")
	})
}

func TestGetEntityType(t *testing.T) {
	// Test successful request
	t.Run("success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/entity-types/1", r.URL.Path)
			assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
			assert.Equal(t, "application/json", r.Header.Get("Accept"))

			response := EntityType{
				ID:          1,
				Name:        "Test Entity Type",
				Description: "Test Description",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		}))
		defer server.Close()

		client := NewClient("test-api-key", WithBaseURL(server.URL))
		result, err := client.GetEntityType(context.Background(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 1, result.ID)
		assert.Equal(t, "Test Entity Type", result.Name)
		assert.Equal(t, "Test Description", result.Description)
	})

	// Test not found error
	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorResponse{
				Message: "Entity type not found",
				Code:    "NOT_FOUND",
			})
		}))
		defer server.Close()

		client := NewClient("test-api-key", WithBaseURL(server.URL))
		result, err := client.GetEntityType(context.Background(), 999)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "Entity type not found")
	})

	// Test invalid response
	t.Run("invalid response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{invalid json`))
		}))
		defer server.Close()

		client := NewClient("test-api-key", WithBaseURL(server.URL))
		result, err := client.GetEntityType(context.Background(), 1)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to decode response")
	})
}
