package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// CardType represents a card type
type CardType struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListCardTypes retrieves a list of card types
func (c *Client) ListCardTypes(ctx context.Context) ([]CardType, error) {
	var response []CardType
	if err := c.doRequest(ctx, http.MethodGet, "/api/card-types", nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetCardType retrieves a single card type by ID
func (c *Client) GetCardType(ctx context.Context, id int) (*CardType, error) {
	var response CardType
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-types/%d", id), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
