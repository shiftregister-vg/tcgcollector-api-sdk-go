package tcgcollector

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListCardListReferences(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-list-references", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := ListCardListReferencesResponse{
			Items: []CardListReference{
				{
					ID:          1,
					CardListID:  100,
					ReferenceID: 200,
					Type:        "related",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			},
			ItemCount:      1,
			TotalItemCount: 1,
			Page:           1,
			PageCount:      1,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	result, err := client.ListCardListReferences(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Items, 1)
	assert.Equal(t, 1, result.Items[0].ID)
	assert.Equal(t, 100, result.Items[0].CardListID)
	assert.Equal(t, 200, result.Items[0].ReferenceID)
	assert.Equal(t, "related", result.Items[0].Type)
}

func TestGetCardListReference(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-list-references/1", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := CardListReference{
			ID:          1,
			CardListID:  100,
			ReferenceID: 200,
			Type:        "related",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	result, err := client.GetCardListReference(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, 100, result.CardListID)
	assert.Equal(t, 200, result.ReferenceID)
	assert.Equal(t, "related", result.Type)
}
