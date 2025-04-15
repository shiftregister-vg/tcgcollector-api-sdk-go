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
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
					Source:     "TCGPlayer",
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
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	result, err := client.ListCardListPrices(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Items, 1)
	assert.Equal(t, 1, result.Items[0].ID)
	assert.Equal(t, 100, result.Items[0].CardListID)
	assert.Equal(t, 10.99, result.Items[0].Price)
	assert.Equal(t, "USD", result.Items[0].Currency)
	assert.Equal(t, "TCGPlayer", result.Items[0].Source)
}

func TestGetCardListPrice(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-list-prices/1", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := CardListPrice{
			ID:         1,
			CardListID: 100,
			Price:      10.99,
			Currency:   "USD",
			Source:     "TCGPlayer",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	result, err := client.GetCardListPrice(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, 100, result.CardListID)
	assert.Equal(t, 10.99, result.Price)
	assert.Equal(t, "USD", result.Currency)
	assert.Equal(t, "TCGPlayer", result.Source)
}
