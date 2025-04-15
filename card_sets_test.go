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

func TestListCardSets(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-sets", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := []Set{
			{
				ID:          1,
				Name:        "Test Set",
				Code:        "TST",
				ReleaseDate: "2024-01-01",
				TotalCards:  100,
				ImageURL:    "http://example.com/image.jpg",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	sets, err := client.ListCardSets(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, sets)
	assert.Len(t, sets, 1)
	assert.Equal(t, 1, sets[0].ID)
	assert.Equal(t, "Test Set", sets[0].Name)
	assert.Equal(t, "TST", sets[0].Code)
}

func TestGetCardSet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-sets/1", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := Set{
			ID:          1,
			Name:        "Test Set",
			Code:        "TST",
			ReleaseDate: "2024-01-01",
			TotalCards:  100,
			ImageURL:    "http://example.com/image.jpg",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	set, err := client.GetCardSet(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, set)
	assert.Equal(t, 1, set.ID)
	assert.Equal(t, "Test Set", set.Name)
	assert.Equal(t, "TST", set.Code)
}

func TestListCardSetsError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Internal server error",
			Code:    "INTERNAL_ERROR",
		})
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	sets, err := client.ListCardSets(context.Background())
	assert.Error(t, err)
	assert.Nil(t, sets)
	assert.Contains(t, err.Error(), "Internal server error")
}

func TestGetCardSetError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Card set not found",
			Code:    "NOT_FOUND",
		})
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	set, err := client.GetCardSet(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, set)
	assert.Contains(t, err.Error(), "Card set not found")
}
