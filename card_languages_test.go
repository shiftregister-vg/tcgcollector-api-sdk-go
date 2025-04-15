package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCardLanguages(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-languages", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"name": "English",
				"code": "en",
				"description": "English language",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			},
			{
				"id": 2,
				"name": "Japanese",
				"code": "ja",
				"description": "Japanese language",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			}
		]`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	languages, err := client.ListCardLanguages(context.Background())
	assert.NoError(t, err)
	assert.Len(t, languages, 2)
	assert.Equal(t, 1, languages[0].ID)
	assert.Equal(t, "English", languages[0].Name)
	assert.Equal(t, "en", languages[0].Code)
	assert.Equal(t, 2, languages[1].ID)
	assert.Equal(t, "Japanese", languages[1].Name)
	assert.Equal(t, "ja", languages[1].Code)
}

func TestGetCardLanguage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-languages/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "English",
			"code": "en",
			"description": "English language",
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	language, err := client.GetCardLanguage(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, language.ID)
	assert.Equal(t, "English", language.Name)
	assert.Equal(t, "en", language.Code)
	assert.Equal(t, "English language", language.Description)
}
