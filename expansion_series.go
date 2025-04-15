package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ListExpansionSeriesParams contains the parameters for listing expansion series
type ListExpansionSeriesParams struct {
	Page     *int `json:"page,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
}

// ListExpansionSeriesResponse represents the response from listing expansion series
type ListExpansionSeriesResponse struct {
	Items          []ExpansionSeries `json:"items"`
	ItemCount      int               `json:"itemCount"`
	TotalItemCount int               `json:"totalItemCount"`
	Page           int               `json:"page"`
	PageCount      int               `json:"pageCount"`
}

// ExpansionSeries represents a series of expansions
type ExpansionSeries struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListExpansionSeries retrieves a list of expansion series
func (c *Client) ListExpansionSeries(ctx context.Context, params *ListExpansionSeriesParams) (*ListExpansionSeriesResponse, error) {
	path := "/api/expansion-series"
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

	var response ListExpansionSeriesResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetExpansionSeries retrieves a single expansion series by ID
func (c *Client) GetExpansionSeries(ctx context.Context, id int) (*ExpansionSeries, error) {
	var response ExpansionSeries
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/expansion-series/%d", id), nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
