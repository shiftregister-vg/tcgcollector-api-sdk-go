package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
)

// CardList represents a card list
type CardList struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	CardCount   int    `json:"cardCount"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// CardListEntry represents an entry in a card list
type CardListEntry struct {
	ID         int    `json:"id"`
	CardListID int    `json:"cardListId"`
	CardID     int    `json:"cardId"`
	Quantity   int    `json:"quantity"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

// ListCardLists retrieves a list of card lists
func (c *Client) ListCardLists(ctx context.Context) ([]CardList, error) {
	var response []CardList
	if err := c.doRequest(ctx, http.MethodGet, "/api/card-lists", nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetCardList retrieves a single card list by ID
func (c *Client) GetCardList(ctx context.Context, id int) (*CardList, error) {
	var response CardList
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-lists/%d", id), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListCardListEntries retrieves entries for a card list
func (c *Client) ListCardListEntries(ctx context.Context, cardListID int) ([]CardListEntry, error) {
	var response []CardListEntry
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-lists/%d/entries", cardListID), nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// RecalculateCardCounts recalculates card counts for all card lists
func (c *Client) RecalculateCardCounts(ctx context.Context) error {
	return c.doRequest(ctx, http.MethodPost, "/api/card-lists/recalculate-card-counts", nil, nil)
}

// RegenerateCardListSlugs regenerates slugs for all card lists
func (c *Client) RegenerateCardListSlugs(ctx context.Context) error {
	return c.doRequest(ctx, http.MethodPost, "/api/card-lists/regenerate-slugs", nil, nil)
}

// BulkReplaceCardListEntries replaces all entries in a card list
func (c *Client) BulkReplaceCardListEntries(ctx context.Context, cardListID int, entries []CardListEntry) error {
	return c.doRequest(ctx, http.MethodPost, fmt.Sprintf("/api/card-lists/%d/entries/bulk-replace", cardListID), entries, nil)
}
