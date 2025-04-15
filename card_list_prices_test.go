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

func TestListCardListPrices(t *testing.T) {
	// Test successful request with no parameters
	t.Run("success without params", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/api/card-list-prices", r.URL.Path)
			assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
			assert.Equal(t, "application/json", r.Header.Get("Accept"))

			response := ListCardListPricesResponse{
				Items: []CardListPrice{
					{
						ID:         1,
						CardListID: 100,
						Price:      10.99,
						Currency:   "USD",
						Source:     "tcgplayer",
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
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
		result, err := client.ListCardListPrices(context.Background(), nil)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Items, 1)
		assert.Equal(t, 1, result.Items[0].ID)
		assert.Equal(t, 100, result.Items[0].CardListID)
		assert.Equal(t, 10.99, result.Items[0].Price)
		assert.Equal(t, "USD", result.Items[0].Currency)
		assert.Equal(t, "tcgplayer", result.Items[0].Source)
	})

	// Test successful request with pagination parameters
	t.Run("success with pagination", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "2", r.URL.Query().Get("page"))
			assert.Equal(t, "20", r.URL.Query().Get("pageSize"))

			response := ListCardListPricesResponse{
				Items:          []CardListPrice{},
				ItemCount:      0,
				TotalItemCount: 40,
				Page:           2,
				PageCount:      2,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		}))
		defer ts.Close()

		client := NewClient("test-api-key", WithBaseURL(ts.URL))
		page := 2
		pageSize := 20
		params := &ListCardListPricesParams{
			Page:     &page,
			PageSize: &pageSize,
		}
		result, err := client.ListCardListPrices(context.Background(), params)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 2, result.Page)
		assert.Equal(t, 2, result.PageCount)
		assert.Equal(t, 40, result.TotalItemCount)
	})

	// Test error response
	t.Run("error response", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{
				Message: "Internal server error",
				Code:    "INTERNAL_ERROR",
			})
		}))
		defer ts.Close()

		client := NewClient("test-api-key", WithBaseURL(ts.URL))
		result, err := client.ListCardListPrices(context.Background(), nil)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "Internal server error")
	})

	// Test invalid response
	t.Run("invalid response", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("invalid json"))
		}))
		defer ts.Close()

		client := NewClient("test-api-key", WithBaseURL(ts.URL))
		result, err := client.ListCardListPrices(context.Background(), nil)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to decode response")
	})
}

func TestGetCardListPrice(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-list-prices/1", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := CardListPrice{
			ID:         1,
			CardListID: 100,
			Price:      10.99,
			Currency:   "USD",
			Source:     "tcgplayer",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	result, err := client.GetCardListPrice(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, 100, result.CardListID)
	assert.Equal(t, 10.99, result.Price)
	assert.Equal(t, "USD", result.Currency)
	assert.Equal(t, "tcgplayer", result.Source)
}

func TestGetCardListPriceError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Card list price not found",
			Code:    "NOT_FOUND",
		})
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	result, err := client.GetCardListPrice(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Card list price not found")
}

func TestGetCardListPriceInvalidResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	result, err := client.GetCardListPrice(context.Background(), 1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to decode response")
}
