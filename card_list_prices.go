package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ListCardListPricesParams contains the parameters for listing card list prices
type ListCardListPricesParams struct {
	Page     *int `json:"page,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
}

// ListCardListPricesResponse represents the response from listing card list prices
type ListCardListPricesResponse struct {
	Items          []CardListPrice `json:"items"`
	ItemCount      int             `json:"itemCount"`
	TotalItemCount int             `json:"totalItemCount"`
	Page           int             `json:"page"`
	PageCount      int             `json:"pageCount"`
}

// CardListPrice represents a price for a card list
type CardListPrice struct {
	ID         int       `json:"id"`
	CardListID int       `json:"cardListId"`
	Price      float64   `json:"price"`
	Currency   string    `json:"currency"`
	Source     string    `json:"source"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// ListCardListPrices retrieves a list of card list prices
func (c *Client) ListCardListPrices(ctx context.Context, params *ListCardListPricesParams) (*ListCardListPricesResponse, error) {
	path := "/api/card-list-prices"
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

	var response ListCardListPricesResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCardListPrice retrieves a single card list price by ID
func (c *Client) GetCardListPrice(ctx context.Context, id int) (*CardListPrice, error) {
	var response CardListPrice
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-list-prices/%d", id), nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
