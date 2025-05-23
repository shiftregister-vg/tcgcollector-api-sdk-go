package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCardFormats(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-formats", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"name": "Standard",
				"description": "The standard format for competitive play",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			},
			{
				"id": 2,
				"name": "Expanded",
				"description": "The expanded format allowing older cards",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			}
		]`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	formats, err := client.ListCardFormats(context.Background())
	assert.NoError(t, err)
	assert.Len(t, formats, 2)
	assert.Equal(t, 1, formats[0].ID)
	assert.Equal(t, "Standard", formats[0].Name)
	assert.Equal(t, "The standard format for competitive play", formats[0].Description)
	assert.Equal(t, 2, formats[1].ID)
	assert.Equal(t, "Expanded", formats[1].Name)
	assert.Equal(t, "The expanded format allowing older cards", formats[1].Description)
}

func TestGetCardFormat(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-formats/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "Standard",
			"description": "The standard format for competitive play",
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	format, err := client.GetCardFormat(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, format.ID)
	assert.Equal(t, "Standard", format.Name)
	assert.Equal(t, "The standard format for competitive play", format.Description)
}

func TestListCardFormatsError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{
			"message": "Internal server error",
			"code": "INTERNAL_ERROR"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	formats, err := client.ListCardFormats(context.Background())
	assert.Error(t, err)
	assert.Nil(t, formats)
	assert.Contains(t, err.Error(), "Internal server error")
}

func TestListCardFormatsInvalidResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	formats, err := client.ListCardFormats(context.Background())
	assert.Error(t, err)
	assert.Nil(t, formats)
	assert.Contains(t, err.Error(), "invalid character")
}

func TestGetCardFormatError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{
			"message": "Card format not found",
			"code": "NOT_FOUND"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	format, err := client.GetCardFormat(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, format)
	assert.Contains(t, err.Error(), "Card format not found")
}

func TestGetCardFormatInvalidResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	format, err := client.GetCardFormat(context.Background(), 1)
	assert.Error(t, err)
	assert.Nil(t, format)
	assert.Contains(t, err.Error(), "invalid character")
}
