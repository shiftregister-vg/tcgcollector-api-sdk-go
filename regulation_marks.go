package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ListRegulationMarksParams contains the parameters for listing regulation marks
type ListRegulationMarksParams struct {
	Page     *int `json:"page,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
}

// ListRegulationMarksResponse represents the response from listing regulation marks
type ListRegulationMarksResponse struct {
	Items          []RegulationMark `json:"items"`
	ItemCount      int              `json:"itemCount"`
	TotalItemCount int              `json:"totalItemCount"`
	Page           int              `json:"page"`
	PageCount      int              `json:"pageCount"`
}

// RegulationMark represents a regulation mark
type RegulationMark struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListRegulationMarks retrieves a list of regulation marks
func (c *Client) ListRegulationMarks(ctx context.Context, params *ListRegulationMarksParams) (*ListRegulationMarksResponse, error) {
	path := "/api/regulation-marks"
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

	var response ListRegulationMarksResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetRegulationMark retrieves a single regulation mark by ID
func (c *Client) GetRegulationMark(ctx context.Context, id int) (*RegulationMark, error) {
	var response RegulationMark
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/regulation-marks/%d", id), nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
