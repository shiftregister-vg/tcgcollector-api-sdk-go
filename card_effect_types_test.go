package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCardEffectTypes(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-effect-types", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"name": "Ability",
				"description": "A special ability that can be used during the game",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			},
			{
				"id": 2,
				"name": "Attack",
				"description": "A move that can be used to deal damage",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			}
		]`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	effectTypes, err := client.ListCardEffectTypes(context.Background())
	assert.NoError(t, err)
	assert.Len(t, effectTypes, 2)
	assert.Equal(t, 1, effectTypes[0].ID)
	assert.Equal(t, "Ability", effectTypes[0].Name)
	assert.Equal(t, "A special ability that can be used during the game", effectTypes[0].Description)
	assert.Equal(t, 2, effectTypes[1].ID)
	assert.Equal(t, "Attack", effectTypes[1].Name)
	assert.Equal(t, "A move that can be used to deal damage", effectTypes[1].Description)
}

func TestGetCardEffectType(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-effect-types/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "Ability",
			"description": "A special ability that can be used during the game",
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	effectType, err := client.GetCardEffectType(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, effectType.ID)
	assert.Equal(t, "Ability", effectType.Name)
	assert.Equal(t, "A special ability that can be used during the game", effectType.Description)
}
