package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// CardEffectType represents a card effect type
type CardEffectType struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListCardEffectTypes retrieves a list of card effect types
func (c *Client) ListCardEffectTypes(ctx context.Context) ([]CardEffectType, error) {
	var response []CardEffectType
	if err := c.doRequest(ctx, http.MethodGet, "/api/card-effect-types", nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetCardEffectType retrieves a single card effect type by ID
func (c *Client) GetCardEffectType(ctx context.Context, id int) (*CardEffectType, error) {
	var response CardEffectType
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-effect-types/%d", id), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
