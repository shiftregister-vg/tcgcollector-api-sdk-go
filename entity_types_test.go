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

func TestListEntityTypes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/entity-types", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := ListEntityTypesResponse{
			Items: []EntityType{
				{
					ID:          1,
					Name:        "Card",
					Description: "A trading card",
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
	types, err := client.ListEntityTypes(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, types)
	assert.Len(t, types.Items, 1)
	assert.Equal(t, 1, types.Items[0].ID)
	assert.Equal(t, "Card", types.Items[0].Name)
	assert.Equal(t, "A trading card", types.Items[0].Description)
}

func TestGetEntityType(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/entity-types/1", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := EntityType{
			ID:          1,
			Name:        "Card",
			Description: "A trading card",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	entityType, err := client.GetEntityType(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, entityType)
	assert.Equal(t, 1, entityType.ID)
	assert.Equal(t, "Card", entityType.Name)
	assert.Equal(t, "A trading card", entityType.Description)
}
