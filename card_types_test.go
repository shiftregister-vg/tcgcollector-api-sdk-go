package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCardTypes(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-types", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"name": "Pokemon",
				"description": "A Pokemon creature card",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			},
			{
				"id": 2,
				"name": "Trainer",
				"description": "A trainer card that provides special effects",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			}
		]`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	types, err := client.ListCardTypes(context.Background())
	assert.NoError(t, err)
	assert.Len(t, types, 2)
	assert.Equal(t, 1, types[0].ID)
	assert.Equal(t, "Pokemon", types[0].Name)
	assert.Equal(t, "A Pokemon creature card", types[0].Description)
	assert.Equal(t, 2, types[1].ID)
	assert.Equal(t, "Trainer", types[1].Name)
	assert.Equal(t, "A trainer card that provides special effects", types[1].Description)
}

func TestGetCardType(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-types/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "Pokemon",
			"description": "A Pokemon creature card",
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	cardType, err := client.GetCardType(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, cardType.ID)
	assert.Equal(t, "Pokemon", cardType.Name)
	assert.Equal(t, "A Pokemon creature card", cardType.Description)
}
