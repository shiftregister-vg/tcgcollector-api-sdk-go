package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListExpansions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/expansions", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"name": "Base Set",
				"description": "Original Pokemon TCG set",
				"slug": "base-set",
				"cardCount": 102,
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			},
			{
				"id": 2,
				"name": "Jungle",
				"description": "Second Pokemon TCG set",
				"slug": "jungle",
				"cardCount": 64,
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			}
		]`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	expansions, err := client.ListExpansions(context.Background())
	assert.NoError(t, err)
	assert.Len(t, expansions, 2)
	assert.Equal(t, 1, expansions[0].ID)
	assert.Equal(t, "Base Set", expansions[0].Name)
	assert.Equal(t, 2, expansions[1].ID)
	assert.Equal(t, "Jungle", expansions[1].Name)
}

func TestGetExpansion(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/expansions/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "Base Set",
			"description": "Original Pokemon TCG set",
			"slug": "base-set",
			"cardCount": 102,
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	expansion, err := client.GetExpansion(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, expansion.ID)
	assert.Equal(t, "Base Set", expansion.Name)
	assert.Equal(t, "Original Pokemon TCG set", expansion.Description)
}

func TestRecalculateExpansionCardCounts(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/expansions/recalculate-card-counts", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	err := client.RecalculateExpansionCardCounts(context.Background())
	assert.NoError(t, err)
}

func TestRegenerateExpansionSlugs(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/expansions/regenerate-slugs", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	err := client.RegenerateExpansionSlugs(context.Background())
	assert.NoError(t, err)
}
