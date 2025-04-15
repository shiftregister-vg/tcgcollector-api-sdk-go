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

func TestListExpansionReferences(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/expansion-references", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := ListExpansionReferencesResponse{
			Items: []ExpansionReference{
				{
					ID:          1,
					ExpansionID: 100,
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
	result, err := client.ListExpansionReferences(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Items, 1)
	assert.Equal(t, 1, result.Items[0].ID)
	assert.Equal(t, 100, result.Items[0].ExpansionID)
	assert.Equal(t, 200, result.Items[0].ReferenceID)
	assert.Equal(t, "related", result.Items[0].Type)
}

func TestGetExpansionReference(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/expansion-references/1", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := ExpansionReference{
			ID:          1,
			ExpansionID: 100,
			ReferenceID: 200,
			Type:        "related",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	reference, err := client.GetExpansionReference(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, reference)
	assert.Equal(t, 1, reference.ID)
	assert.Equal(t, 100, reference.ExpansionID)
	assert.Equal(t, 200, reference.ReferenceID)
	assert.Equal(t, "related", reference.Type)
}

func TestGetExpansionReferenceError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Expansion reference not found",
			Code:    "NOT_FOUND",
		})
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	reference, err := client.GetExpansionReference(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, reference)
	assert.Contains(t, err.Error(), "Expansion reference not found")
}

func TestGetExpansionReferenceInvalidResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	reference, err := client.GetExpansionReference(context.Background(), 1)
	assert.Error(t, err)
	assert.Nil(t, reference)
	assert.Contains(t, err.Error(), "failed to decode response")
}

func TestListExpansionReferencesWithPagination(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/expansion-references", r.URL.Path)
		assert.Equal(t, "2", r.URL.Query().Get("page"))
		assert.Equal(t, "10", r.URL.Query().Get("pageSize"))
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := ListExpansionReferencesResponse{
			Items: []ExpansionReference{
				{
					ID:          2,
					ExpansionID: 101,
					ReferenceID: 201,
					Type:        "related",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			},
			ItemCount:      1,
			TotalItemCount: 15,
			Page:           2,
			PageCount:      2,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	page := 2
	pageSize := 10
	params := &ListExpansionReferencesParams{
		Page:     &page,
		PageSize: &pageSize,
	}
	result, err := client.ListExpansionReferences(context.Background(), params)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Items, 1)
	assert.Equal(t, 2, result.Items[0].ID)
	assert.Equal(t, 101, result.Items[0].ExpansionID)
	assert.Equal(t, 201, result.Items[0].ReferenceID)
	assert.Equal(t, 2, result.Page)
	assert.Equal(t, 2, result.PageCount)
	assert.Equal(t, 15, result.TotalItemCount)
}

func TestListExpansionReferencesError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Internal server error",
			Code:    "INTERNAL_SERVER_ERROR",
		})
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	result, err := client.ListExpansionReferences(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Internal server error")
}

func TestListExpansionReferencesInvalidResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"invalid json`))
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	result, err := client.ListExpansionReferences(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to decode response")
}
