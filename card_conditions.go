package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// CardCondition represents a card condition
type CardCondition struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListCardConditions retrieves a list of card conditions
func (c *Client) ListCardConditions(ctx context.Context) ([]CardCondition, error) {
	var response []CardCondition
	if err := c.doRequest(ctx, http.MethodGet, "/api/card-conditions", nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetCardCondition retrieves a single card condition by ID
func (c *Client) GetCardCondition(ctx context.Context, id int) (*CardCondition, error) {
	var response CardCondition
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-conditions/%d", id), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
