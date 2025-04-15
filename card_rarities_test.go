package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCardRarities(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-rarities", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"name": "Common",
				"description": "Common rarity",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			},
			{
				"id": 2,
				"name": "Rare",
				"description": "Rare rarity",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			}
		]`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	rarities, err := client.ListCardRarities(context.Background())
	assert.NoError(t, err)
	assert.Len(t, rarities, 2)
	assert.Equal(t, 1, rarities[0].ID)
	assert.Equal(t, "Common", rarities[0].Name)
	assert.Equal(t, 2, rarities[1].ID)
	assert.Equal(t, "Rare", rarities[1].Name)
}

func TestGetCardRarity(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-rarities/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "Common",
			"description": "Common rarity",
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	rarity, err := client.GetCardRarity(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, rarity.ID)
	assert.Equal(t, "Common", rarity.Name)
	assert.Equal(t, "Common rarity", rarity.Description)
}
