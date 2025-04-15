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

func TestListPokemonStages(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/pokemon-stages", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := ListPokemonStagesResponse{
			Items: []PokemonStage{
				{
					ID:          1,
					Name:        "Basic",
					Description: "The initial stage of a Pokémon",
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
	stages, err := client.ListPokemonStages(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, stages)
	assert.Len(t, stages.Items, 1)
	assert.Equal(t, 1, stages.Items[0].ID)
	assert.Equal(t, "Basic", stages.Items[0].Name)
	assert.Equal(t, "The initial stage of a Pokémon", stages.Items[0].Description)
}

func TestGetPokemonStage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/pokemon-stages/1", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := PokemonStage{
			ID:          1,
			Name:        "Basic",
			Description: "The initial stage of a Pokémon",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	stage, err := client.GetPokemonStage(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, stage)
	assert.Equal(t, 1, stage.ID)
	assert.Equal(t, "Basic", stage.Name)
	assert.Equal(t, "The initial stage of a Pokémon", stage.Description)
}
