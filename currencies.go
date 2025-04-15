package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Currency represents a currency
type Currency struct {
	ID        int       `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ListCurrencies retrieves a list of currencies
func (c *Client) ListCurrencies(ctx context.Context) ([]Currency, error) {
	var response []Currency
	if err := c.doRequest(ctx, http.MethodGet, "/api/currencies", nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetCurrency retrieves a single currency by ID
func (c *Client) GetCurrency(ctx context.Context, id int) (*Currency, error) {
	var response Currency
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/currencies/%d", id), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
