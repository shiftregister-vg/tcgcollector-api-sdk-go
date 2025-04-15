package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// CardSupertype represents a card supertype
type CardSupertype struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListCardSupertypes retrieves a list of card supertypes
func (c *Client) ListCardSupertypes(ctx context.Context) ([]CardSupertype, error) {
	var response []CardSupertype
	if err := c.doRequest(ctx, http.MethodGet, "/api/card-supertypes", nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetCardSupertype retrieves a single card supertype by ID
func (c *Client) GetCardSupertype(ctx context.Context, id int) (*CardSupertype, error) {
	var response CardSupertype
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-supertypes/%d", id), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
