package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCardConditions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-conditions", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"name": "Near Mint",
				"description": "Card is in near mint condition",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			},
			{
				"id": 2,
				"name": "Lightly Played",
				"description": "Card shows minimal wear",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			}
		]`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	conditions, err := client.ListCardConditions(context.Background())
	assert.NoError(t, err)
	assert.Len(t, conditions, 2)
	assert.Equal(t, 1, conditions[0].ID)
	assert.Equal(t, "Near Mint", conditions[0].Name)
	assert.Equal(t, "Card is in near mint condition", conditions[0].Description)
	assert.Equal(t, 2, conditions[1].ID)
	assert.Equal(t, "Lightly Played", conditions[1].Name)
	assert.Equal(t, "Card shows minimal wear", conditions[1].Description)
}

func TestGetCardCondition(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-conditions/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "Near Mint",
			"description": "Card is in near mint condition",
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	condition, err := client.GetCardCondition(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, condition.ID)
	assert.Equal(t, "Near Mint", condition.Name)
	assert.Equal(t, "Card is in near mint condition", condition.Description)
}
