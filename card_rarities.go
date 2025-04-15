package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
)

// CardRarity represents a card rarity
type CardRarity struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// ListCardRarities retrieves a list of card rarities
func (c *Client) ListCardRarities(ctx context.Context) ([]CardRarity, error) {
	var response []CardRarity
	if err := c.doRequest(ctx, http.MethodGet, "/api/card-rarities", nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetCardRarity retrieves a single card rarity by ID
func (c *Client) GetCardRarity(ctx context.Context, id int) (*CardRarity, error) {
	var response CardRarity
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-rarities/%d", id), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
