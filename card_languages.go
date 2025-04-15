package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// CardLanguage represents a card language
type CardLanguage struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListCardLanguages retrieves a list of card languages
func (c *Client) ListCardLanguages(ctx context.Context) ([]CardLanguage, error) {
	var response []CardLanguage
	if err := c.doRequest(ctx, http.MethodGet, "/api/card-languages", nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetCardLanguage retrieves a single card language by ID
func (c *Client) GetCardLanguage(ctx context.Context, id int) (*CardLanguage, error) {
	var response CardLanguage
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-languages/%d", id), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
