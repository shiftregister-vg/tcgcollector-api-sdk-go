package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// CardIllustrator represents a card illustrator
type CardIllustrator struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"imageUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListCardIllustrators retrieves a list of card illustrators
func (c *Client) ListCardIllustrators(ctx context.Context) ([]CardIllustrator, error) {
	var response []CardIllustrator
	if err := c.doRequest(ctx, http.MethodGet, "/api/card-illustrators", nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetCardIllustrator retrieves a single card illustrator by ID
func (c *Client) GetCardIllustrator(ctx context.Context, id int) (*CardIllustrator, error) {
	var response CardIllustrator
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-illustrators/%d", id), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
