package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ListExpansionPricesParams contains the parameters for listing expansion prices
type ListExpansionPricesParams struct {
	Page     *int `json:"page,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
}

// ListExpansionPricesResponse represents the response from listing expansion prices
type ListExpansionPricesResponse struct {
	Items          []ExpansionPrice `json:"items"`
	ItemCount      int              `json:"itemCount"`
	TotalItemCount int              `json:"totalItemCount"`
	Page           int              `json:"page"`
	PageCount      int              `json:"pageCount"`
}

// ExpansionPrice represents a price for an expansion
type ExpansionPrice struct {
	ID          int       `json:"id"`
	ExpansionID int       `json:"expansionId"`
	Price       float64   `json:"price"`
	Currency    string    `json:"currency"`
	Source      string    `json:"source"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListExpansionPrices retrieves a list of expansion prices
func (c *Client) ListExpansionPrices(ctx context.Context, params *ListExpansionPricesParams) (*ListExpansionPricesResponse, error) {
	path := "/api/expansion-prices"
	if params != nil {
		query := url.Values{}
		if params.Page != nil {
			query.Set("page", fmt.Sprintf("%d", *params.Page))
		}
		if params.PageSize != nil {
			query.Set("pageSize", fmt.Sprintf("%d", *params.PageSize))
		}
		if len(query) > 0 {
			path = fmt.Sprintf("%s?%s", path, query.Encode())
		}
	}

	var response ListExpansionPricesResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetExpansionPrice retrieves a single expansion price by ID
func (c *Client) GetExpansionPrice(ctx context.Context, id int) (*ExpansionPrice, error) {
	var response ExpansionPrice
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/expansion-prices/%d", id), nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
