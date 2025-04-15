package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// CardFormat represents a card format
type CardFormat struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListCardFormats retrieves a list of card formats
func (c *Client) ListCardFormats(ctx context.Context) ([]CardFormat, error) {
	var response []CardFormat
	if err := c.doRequest(ctx, http.MethodGet, "/api/card-formats", nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetCardFormat retrieves a single card format by ID
func (c *Client) GetCardFormat(ctx context.Context, id int) (*CardFormat, error) {
	var response CardFormat
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-formats/%d", id), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
