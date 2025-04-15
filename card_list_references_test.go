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

func TestListCardListReferences(t *testing.T) {
	// Test successful request with no parameters
	t.Run("success without params", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/card-list-references", r.URL.Path)
			assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
			assert.Equal(t, "application/json", r.Header.Get("Accept"))

			response := ListCardListReferencesResponse{
				Items: []CardListReference{
					{
						ID:          1,
						CardListID:  100,
						ReferenceID: 200,
						Type:        "related",
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
		result, err := client.ListCardListReferences(context.Background(), nil)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Items, 1)
		assert.Equal(t, 1, result.Items[0].ID)
		assert.Equal(t, 100, result.Items[0].CardListID)
		assert.Equal(t, 200, result.Items[0].ReferenceID)
		assert.Equal(t, "related", result.Items[0].Type)
	})

	// Test successful request with pagination parameters
	t.Run("success with pagination", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "2", r.URL.Query().Get("page"))
			assert.Equal(t, "50", r.URL.Query().Get("pageSize"))

			response := ListCardListReferencesResponse{
				Items:          []CardListReference{},
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
		result, err := client.ListCardListReferences(context.Background(), &ListCardListReferencesParams{
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
		result, err := client.ListCardListReferences(context.Background(), nil)
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
		result, err := client.ListCardListReferences(context.Background(), nil)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to decode response")
	})
}

func TestGetCardListReference(t *testing.T) {
	// Test successful request
	t.Run("success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/card-list-references/1", r.URL.Path)
			assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
			assert.Equal(t, "application/json", r.Header.Get("Accept"))

			response := CardListReference{
				ID:          1,
				CardListID:  100,
				ReferenceID: 200,
				Type:        "related",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		}))
		defer server.Close()

		client := NewClient("test-api-key", WithBaseURL(server.URL))
		result, err := client.GetCardListReference(context.Background(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 1, result.ID)
		assert.Equal(t, 100, result.CardListID)
		assert.Equal(t, 200, result.ReferenceID)
		assert.Equal(t, "related", result.Type)
	})

	// Test not found error
	t.Run("not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorResponse{
				Message: "Card list reference not found",
				Code:    "NOT_FOUND",
			})
		}))
		defer server.Close()

		client := NewClient("test-api-key", WithBaseURL(server.URL))
		result, err := client.GetCardListReference(context.Background(), 999)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "Card list reference not found")
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
		result, err := client.GetCardListReference(context.Background(), 1)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to decode response")
	})
}
