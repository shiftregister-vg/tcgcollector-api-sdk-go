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

func TestListCardVariantReferences(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-variant-references", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := ListCardVariantReferencesResponse{
			Items: []CardVariantReference{
				{
					ID:            1,
					CardVariantID: 100,
					ReferenceID:   200,
					Type:          "related",
					CreatedAt:     time.Now(),
					UpdatedAt:     time.Now(),
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
	result, err := client.ListCardVariantReferences(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Items, 1)
	assert.Equal(t, 1, result.Items[0].ID)
	assert.Equal(t, 100, result.Items[0].CardVariantID)
	assert.Equal(t, 200, result.Items[0].ReferenceID)
	assert.Equal(t, "related", result.Items[0].Type)
}

func TestListCardVariantReferencesWithPagination(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "2", r.URL.Query().Get("page"))
		assert.Equal(t, "50", r.URL.Query().Get("pageSize"))

		response := ListCardVariantReferencesResponse{
			Items:          []CardVariantReference{},
			ItemCount:      0,
			TotalItemCount: 100,
			Page:           2,
			PageCount:      3,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	page := 2
	pageSize := 50
	result, err := client.ListCardVariantReferences(context.Background(), &ListCardVariantReferencesParams{
		Page:     &page,
		PageSize: &pageSize,
	})
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.Page)
	assert.Equal(t, 3, result.PageCount)
	assert.Equal(t, 100, result.TotalItemCount)
}

func TestListCardVariantReferencesError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "internal server error"}`))
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	result, err := client.ListCardVariantReferences(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestListCardVariantReferencesInvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{invalid json`))
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	result, err := client.ListCardVariantReferences(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetCardVariantReference(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-variant-references/1", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := CardVariantReference{
			ID:            1,
			CardVariantID: 100,
			ReferenceID:   200,
			Type:          "related",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	result, err := client.GetCardVariantReference(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, 100, result.CardVariantID)
	assert.Equal(t, 200, result.ReferenceID)
	assert.Equal(t, "related", result.Type)
}

func TestGetCardVariantReferenceNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "card variant reference not found"}`))
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	result, err := client.GetCardVariantReference(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetCardVariantReferenceError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "internal server error"}`))
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	result, err := client.GetCardVariantReference(context.Background(), 1)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetCardVariantReferenceInvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{invalid json`))
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	result, err := client.GetCardVariantReference(context.Background(), 1)
	assert.Error(t, err)
	assert.Nil(t, result)
}
