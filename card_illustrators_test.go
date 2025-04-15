package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCardIllustrators(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-illustrators", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"name": "Ken Sugimori",
				"description": "Original Pokemon illustrator",
				"imageUrl": "https://example.com/sugimori.jpg",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			},
			{
				"id": 2,
				"name": "Mitsuhiro Arita",
				"description": "Famous Pokemon card illustrator",
				"imageUrl": "https://example.com/arita.jpg",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			}
		]`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	illustrators, err := client.ListCardIllustrators(context.Background())
	assert.NoError(t, err)
	assert.Len(t, illustrators, 2)
	assert.Equal(t, 1, illustrators[0].ID)
	assert.Equal(t, "Ken Sugimori", illustrators[0].Name)
	assert.Equal(t, "Original Pokemon illustrator", illustrators[0].Description)
	assert.Equal(t, 2, illustrators[1].ID)
	assert.Equal(t, "Mitsuhiro Arita", illustrators[1].Name)
	assert.Equal(t, "Famous Pokemon card illustrator", illustrators[1].Description)
}

func TestGetCardIllustrator(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-illustrators/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "Ken Sugimori",
			"description": "Original Pokemon illustrator",
			"imageUrl": "https://example.com/sugimori.jpg",
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	illustrator, err := client.GetCardIllustrator(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, illustrator.ID)
	assert.Equal(t, "Ken Sugimori", illustrator.Name)
	assert.Equal(t, "Original Pokemon illustrator", illustrator.Description)
	assert.Equal(t, "https://example.com/sugimori.jpg", illustrator.ImageURL)
}

func TestListCardIllustratorsError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-illustrators", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{
			"message": "Internal server error",
			"code": "INTERNAL_ERROR"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	illustrators, err := client.ListCardIllustrators(context.Background())
	assert.Error(t, err)
	assert.Nil(t, illustrators)
	assert.Contains(t, err.Error(), "Internal server error")
}

func TestListCardIllustratorsInvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-illustrators", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	illustrators, err := client.ListCardIllustrators(context.Background())
	assert.Error(t, err)
	assert.Nil(t, illustrators)
	assert.Contains(t, err.Error(), "invalid character")
}

func TestGetCardIllustratorNotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-illustrators/999", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{
			"message": "Card illustrator not found",
			"code": "NOT_FOUND"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	illustrator, err := client.GetCardIllustrator(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, illustrator)
	assert.Contains(t, err.Error(), "Card illustrator not found")
}

func TestGetCardIllustratorInvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-illustrators/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	illustrator, err := client.GetCardIllustrator(context.Background(), 1)
	assert.Error(t, err)
	assert.Nil(t, illustrator)
	assert.Contains(t, err.Error(), "invalid character")
}
