package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ListTCGPriceSourcesParams contains the parameters for listing TCG price sources
type ListTCGPriceSourcesParams struct {
	Page     *int `json:"page,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
}

// ListTCGPriceSourcesResponse represents the response from listing TCG price sources
type ListTCGPriceSourcesResponse struct {
	Items          []TCGPriceSource `json:"items"`
	ItemCount      int              `json:"itemCount"`
	TotalItemCount int              `json:"totalItemCount"`
	Page           int              `json:"page"`
	PageCount      int              `json:"pageCount"`
}

// TCGPriceSource represents a TCG price source
type TCGPriceSource struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListTCGPriceSources retrieves a list of TCG price sources
func (c *Client) ListTCGPriceSources(ctx context.Context, params *ListTCGPriceSourcesParams) (*ListTCGPriceSourcesResponse, error) {
	path := "/api/tcg-price-sources"
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

	var response ListTCGPriceSourcesResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetTCGPriceSource retrieves a single TCG price source by ID
func (c *Client) GetTCGPriceSource(ctx context.Context, id int) (*TCGPriceSource, error) {
	var response TCGPriceSource
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/tcg-price-sources/%d", id), nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
