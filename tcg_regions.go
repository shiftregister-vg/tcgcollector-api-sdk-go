package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ListTCGRegionsParams contains the parameters for listing TCG regions
type ListTCGRegionsParams struct {
	Page     *int `json:"page,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
}

// ListTCGRegionsResponse represents the response from listing TCG regions
type ListTCGRegionsResponse struct {
	Items          []TCGRegion `json:"items"`
	ItemCount      int         `json:"itemCount"`
	TotalItemCount int         `json:"totalItemCount"`
	Page           int         `json:"page"`
	PageCount      int         `json:"pageCount"`
}

// TCGRegion represents a TCG region
type TCGRegion struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Code        string    `json:"code"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListTCGRegions retrieves a list of TCG regions
func (c *Client) ListTCGRegions(ctx context.Context, params *ListTCGRegionsParams) (*ListTCGRegionsResponse, error) {
	path := "/api/tcg-regions"
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

	var response ListTCGRegionsResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetTCGRegion retrieves a single TCG region by ID
func (c *Client) GetTCGRegion(ctx context.Context, id int) (*TCGRegion, error) {
	var response TCGRegion
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/tcg-regions/%d", id), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
