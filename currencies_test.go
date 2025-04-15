package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCurrencies(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/currencies", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"code": "USD",
				"name": "US Dollar",
				"symbol": "$",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			},
			{
				"id": 2,
				"code": "EUR",
				"name": "Euro",
				"symbol": "€",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			}
		]`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	currencies, err := client.ListCurrencies(context.Background())
	assert.NoError(t, err)
	assert.Len(t, currencies, 2)
	assert.Equal(t, 1, currencies[0].ID)
	assert.Equal(t, "USD", currencies[0].Code)
	assert.Equal(t, "US Dollar", currencies[0].Name)
	assert.Equal(t, "$", currencies[0].Symbol)
	assert.Equal(t, 2, currencies[1].ID)
	assert.Equal(t, "EUR", currencies[1].Code)
	assert.Equal(t, "Euro", currencies[1].Name)
	assert.Equal(t, "€", currencies[1].Symbol)
}

func TestGetCurrency(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/currencies/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"id": 1,
			"code": "USD",
			"name": "US Dollar",
			"symbol": "$",
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	currency, err := client.GetCurrency(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, currency.ID)
	assert.Equal(t, "USD", currency.Code)
	assert.Equal(t, "US Dollar", currency.Name)
	assert.Equal(t, "$", currency.Symbol)
}
