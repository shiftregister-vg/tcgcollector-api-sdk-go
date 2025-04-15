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

func TestInvalidateCardListCache(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/card-collection/invalidate-card-list-cache", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	err := client.InvalidateCardListCache(context.Background())
	assert.NoError(t, err)
}

func TestInvalidateExpansionCache(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/card-collection/invalidate-expansion-cache", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	err := client.InvalidateExpansionCache(context.Background())
	assert.NoError(t, err)
}

func TestListCollections(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/collections", r.URL.Path)

		// Check query parameters
		query := r.URL.Query()
		assert.Equal(t, "1", query.Get("userId"))
		assert.Equal(t, "1", query.Get("page"))
		assert.Equal(t, "10", query.Get("pageSize"))

		// Create mock response
		response := ListResponse[Collection]{
			Items: []Collection{
				{
					ID:          1,
					UserID:      1,
					Name:        "Test Collection",
					Description: "Test collection description",
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
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Create parameters
	userID := 1
	page := 1
	pageSize := 10

	params := &ListCollectionsParams{
		UserID:   &userID,
		Page:     &page,
		PageSize: &pageSize,
	}

	result, err := client.ListCollections(context.Background(), params)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result.Items))
	assert.Equal(t, 1, result.Items[0].ID)
	assert.Equal(t, "Test Collection", result.Items[0].Name)
}

func TestGetCollection(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/collections/1", r.URL.Path)

		// Create mock response
		response := Collection{
			ID:          1,
			UserID:      1,
			Name:        "Test Collection",
			Description: "Test collection description",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	result, err := client.GetCollection(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test Collection", result.Name)
}

func TestCreateCollection(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/collections", r.URL.Path)

		// Create mock response
		response := Collection{
			ID:          1,
			UserID:      1,
			Name:        "Test Collection",
			Description: "Test collection description",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	collection := &Collection{
		Name:        "Test Collection",
		Description: "Test collection description",
	}

	result, err := client.CreateCollection(context.Background(), collection)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test Collection", result.Name)
}

func TestUpdateCollection(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "/api/collections/1", r.URL.Path)

		// Create mock response
		response := Collection{
			ID:          1,
			UserID:      1,
			Name:        "Updated Collection",
			Description: "Updated collection description",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	collection := &Collection{
		Name:        "Updated Collection",
		Description: "Updated collection description",
	}

	result, err := client.UpdateCollection(context.Background(), 1, collection)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Updated Collection", result.Name)
}

func TestDeleteCollection(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/api/collections/1", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	err := client.DeleteCollection(context.Background(), 1)
	assert.NoError(t, err)
}

func TestListCollectionCards(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/collections/1/cards", r.URL.Path)

		// Create mock response
		response := ListResponse[CollectionCard]{
			Items: []CollectionCard{
				{
					ID:           1,
					CollectionID: 1,
					CardID:       1,
					Quantity:     1,
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
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
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	result, err := client.ListCollectionCards(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result.Items))
	assert.Equal(t, 1, result.Items[0].ID)
}

func TestAddCardToCollection(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/collections/1/cards", r.URL.Path)

		// Create mock response
		response := CollectionCard{
			ID:           1,
			CollectionID: 1,
			CardID:       1,
			Quantity:     1,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	card := &CollectionCard{
		CardID:   1,
		Quantity: 1,
	}

	result, err := client.AddCardToCollection(context.Background(), 1, card)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
}

func TestUpdateCollectionCard(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "/api/collections/1/cards/1", r.URL.Path)

		// Create mock response
		response := CollectionCard{
			ID:           1,
			CollectionID: 1,
			CardID:       2,
			Quantity:     2,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	card := &CollectionCard{
		Quantity: 2,
	}

	result, err := client.UpdateCollectionCard(context.Background(), 1, 1, card)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.Quantity)
}

func TestRemoveCardFromCollection(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/api/collections/1/cards/1", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	err := client.RemoveCardFromCollection(context.Background(), 1, 1)
	assert.NoError(t, err)
}

func TestListCollectionsWithParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/collections", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		// Verify query parameters
		assert.Equal(t, "1", r.URL.Query().Get("userId"))
		assert.Equal(t, "test", r.URL.Query().Get("name"))
		assert.Equal(t, "true", r.URL.Query().Get("isPublic"))
		assert.Equal(t, "2", r.URL.Query().Get("page"))
		assert.Equal(t, "10", r.URL.Query().Get("pageSize"))

		response := ListResponse[Collection]{
			Items: []Collection{
				{
					ID:        1,
					UserID:    1,
					Name:      "test",
					IsPublic:  true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			ItemCount:      1,
			TotalItemCount: 1,
			Page:           2,
			PageCount:      3,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	userID := 1
	name := "test"
	isPublic := true
	page := 2
	pageSize := 10
	params := &ListCollectionsParams{
		UserID:   &userID,
		Name:     &name,
		IsPublic: &isPublic,
		Page:     &page,
		PageSize: &pageSize,
	}
	result, err := client.ListCollections(context.Background(), params)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Items, 1)
	assert.Equal(t, 1, result.Items[0].ID)
	assert.Equal(t, "test", result.Items[0].Name)
	assert.True(t, result.Items[0].IsPublic)
}

func TestListCollectionsError(t *testing.T) {
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
	result, err := client.ListCollections(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Internal server error")
}

func TestListCollectionsInvalidResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"invalid json`))
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	result, err := client.ListCollections(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to decode response")
}

func TestGetCollectionError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Collection not found",
			Code:    "NOT_FOUND",
		})
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	result, err := client.GetCollection(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Collection not found")
}
