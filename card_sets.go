package tcgcollector

import (
	"context"
	"fmt"
)

// ListCardSets retrieves a list of all card sets
func (c *Client) ListCardSets(ctx context.Context) ([]Set, error) {
	var sets []Set
	if err := c.doRequest(ctx, "GET", "/api/card-sets", nil, &sets); err != nil {
		return nil, fmt.Errorf("failed to list card sets: %w", err)
	}
	return sets, nil
}

// GetCardSet retrieves a specific card set by ID
func (c *Client) GetCardSet(ctx context.Context, id int) (*Set, error) {
	var set Set
	if err := c.doRequest(ctx, "GET", fmt.Sprintf("/api/card-sets/%d", id), nil, &set); err != nil {
		return nil, fmt.Errorf("failed to get card set: %w", err)
	}
	return &set, nil
}
