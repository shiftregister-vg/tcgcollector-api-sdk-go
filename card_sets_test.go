package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCardSets(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-sets", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"name": "Base Set",
				"code": "BS",
				"releaseDate": "1999-01-09",
				"totalCards": 102,
				"description": "The first Pokemon TCG set",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			},
			{
				"id": 2,
				"name": "Jungle",
				"code": "JU",
				"releaseDate": "1999-06-16",
				"totalCards": 64,
				"description": "The second Pokemon TCG set",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			}
		]`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	sets, err := client.ListCardSets(context.Background())
	assert.NoError(t, err)
	assert.Len(t, sets, 2)
	assert.Equal(t, 1, sets[0].ID)
	assert.Equal(t, "Base Set", sets[0].Name)
	assert.Equal(t, "BS", sets[0].Code)
	assert.Equal(t, 2, sets[1].ID)
	assert.Equal(t, "Jungle", sets[1].Name)
	assert.Equal(t, "JU", sets[1].Code)
}

func TestGetCardSet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-sets/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "Base Set",
			"code": "BS",
			"releaseDate": "1999-01-09",
			"totalCards": 102,
			"description": "The first Pokemon TCG set",
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	set, err := client.GetCardSet(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, set.ID)
	assert.Equal(t, "Base Set", set.Name)
	assert.Equal(t, "BS", set.Code)
	assert.Equal(t, 102, set.TotalCards)
}
