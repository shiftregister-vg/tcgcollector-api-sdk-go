package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCardDatabaseLogEntries(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-database-log", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"action": "CREATE",
				"details": "Created new card",
				"createdAt": "2024-01-01T00:00:00Z"
			},
			{
				"id": 2,
				"action": "UPDATE",
				"details": "Updated card price",
				"createdAt": "2024-01-01T00:00:00Z"
			}
		]`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	entries, err := client.ListCardDatabaseLogEntries(context.Background())
	assert.NoError(t, err)
	assert.Len(t, entries, 2)
	assert.Equal(t, 1, entries[0].ID)
	assert.Equal(t, "CREATE", entries[0].Action)
	assert.Equal(t, 2, entries[1].ID)
	assert.Equal(t, "UPDATE", entries[1].Action)
}

func TestPruneCardDatabaseLog(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/card-database-log/prune", r.URL.Path)

		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	err := client.PruneCardDatabaseLog(context.Background())
	assert.NoError(t, err)
}
