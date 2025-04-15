package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCardSupertypes(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-supertypes", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"name": "Basic",
				"description": "A basic Pokemon card that cannot evolve from another card",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			},
			{
				"id": 2,
				"name": "Stage 1",
				"description": "A Pokemon card that evolves from a Basic Pokemon",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			}
		]`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	supertypes, err := client.ListCardSupertypes(context.Background())
	assert.NoError(t, err)
	assert.Len(t, supertypes, 2)
	assert.Equal(t, 1, supertypes[0].ID)
	assert.Equal(t, "Basic", supertypes[0].Name)
	assert.Equal(t, "A basic Pokemon card that cannot evolve from another card", supertypes[0].Description)
	assert.Equal(t, 2, supertypes[1].ID)
	assert.Equal(t, "Stage 1", supertypes[1].Name)
	assert.Equal(t, "A Pokemon card that evolves from a Basic Pokemon", supertypes[1].Description)
}

func TestGetCardSupertype(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-supertypes/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "Basic",
			"description": "A basic Pokemon card that cannot evolve from another card",
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	supertype, err := client.GetCardSupertype(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, supertype.ID)
	assert.Equal(t, "Basic", supertype.Name)
	assert.Equal(t, "A basic Pokemon card that cannot evolve from another card", supertype.Description)
}
