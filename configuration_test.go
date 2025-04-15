package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllowedExternalAccountHosts(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/configuration/allowed-external-account-hosts", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"hosts": [
				"example.com",
				"api.example.com"
			]
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	hosts, err := client.GetAllowedExternalAccountHosts(context.Background())
	assert.NoError(t, err)
	assert.Len(t, hosts.Hosts, 2)
	assert.Equal(t, "example.com", hosts.Hosts[0])
	assert.Equal(t, "api.example.com", hosts.Hosts[1])
}

func TestGetBaseTCGCurrency(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/configuration/base-tcg-currency", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"currency": "USD"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	currency, err := client.GetBaseTCGCurrency(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, "USD", currency.Currency)
}
