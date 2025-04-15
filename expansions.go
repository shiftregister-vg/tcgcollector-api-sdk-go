package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
)

// Expansion represents a card expansion
type Expansion struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	CardCount   int    `json:"cardCount"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// ListExpansions retrieves a list of expansions
func (c *Client) ListExpansions(ctx context.Context) ([]Expansion, error) {
	var response []Expansion
	if err := c.doRequest(ctx, http.MethodGet, "/api/expansions", nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetExpansion retrieves a single expansion by ID
func (c *Client) GetExpansion(ctx context.Context, id int) (*Expansion, error) {
	var response Expansion
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/expansions/%d", id), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// RecalculateCardCounts recalculates card counts for all expansions
func (c *Client) RecalculateExpansionCardCounts(ctx context.Context) error {
	return c.doRequest(ctx, http.MethodPost, "/api/expansions/recalculate-card-counts", nil, nil)
}

// RegenerateSlugs regenerates slugs for all expansions
func (c *Client) RegenerateExpansionSlugs(ctx context.Context) error {
	return c.doRequest(ctx, http.MethodPost, "/api/expansions/regenerate-slugs", nil, nil)
}
