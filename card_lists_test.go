package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCardLists(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-lists", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"name": "My Collection",
				"description": "My personal collection",
				"slug": "my-collection",
				"cardCount": 100,
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			},
			{
				"id": 2,
				"name": "Trade Binder",
				"description": "Cards for trade",
				"slug": "trade-binder",
				"cardCount": 50,
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			}
		]`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	lists, err := client.ListCardLists(context.Background())
	assert.NoError(t, err)
	assert.Len(t, lists, 2)
	assert.Equal(t, 1, lists[0].ID)
	assert.Equal(t, "My Collection", lists[0].Name)
	assert.Equal(t, 2, lists[1].ID)
	assert.Equal(t, "Trade Binder", lists[1].Name)
}

func TestGetCardList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-lists/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "My Collection",
			"description": "My personal collection",
			"slug": "my-collection",
			"cardCount": 100,
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	list, err := client.GetCardList(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, list.ID)
	assert.Equal(t, "My Collection", list.Name)
	assert.Equal(t, "My personal collection", list.Description)
}

func TestListCardListEntries(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-lists/1/entries", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"cardListId": 1,
				"cardId": 1,
				"quantity": 2,
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			},
			{
				"id": 2,
				"cardListId": 1,
				"cardId": 2,
				"quantity": 1,
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			}
		]`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	entries, err := client.ListCardListEntries(context.Background(), 1)
	assert.NoError(t, err)
	assert.Len(t, entries, 2)
	assert.Equal(t, 1, entries[0].ID)
	assert.Equal(t, 1, entries[0].CardID)
	assert.Equal(t, 2, entries[0].Quantity)
	assert.Equal(t, 2, entries[1].ID)
	assert.Equal(t, 2, entries[1].CardID)
	assert.Equal(t, 1, entries[1].Quantity)
}

func TestRecalculateCardCounts(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/card-lists/recalculate-card-counts", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	err := client.RecalculateCardCounts(context.Background())
	assert.NoError(t, err)
}

func TestRegenerateCardListSlugs(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/card-lists/regenerate-slugs", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	err := client.RegenerateCardListSlugs(context.Background())
	assert.NoError(t, err)
}

func TestBulkReplaceCardListEntries(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/card-lists/1/entries/bulk-replace", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	entries := []CardListEntry{
		{
			CardID:   1,
			Quantity: 2,
		},
		{
			CardID:   2,
			Quantity: 1,
		},
	}
	err := client.BulkReplaceCardListEntries(context.Background(), 1, entries)
	assert.NoError(t, err)
}
