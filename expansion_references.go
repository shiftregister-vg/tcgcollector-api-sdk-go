package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ListExpansionReferencesParams contains the parameters for listing expansion references
type ListExpansionReferencesParams struct {
	Page     *int `json:"page,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
}

// ListExpansionReferencesResponse represents the response from listing expansion references
type ListExpansionReferencesResponse struct {
	Items          []ExpansionReference `json:"items"`
	ItemCount      int                  `json:"itemCount"`
	TotalItemCount int                  `json:"totalItemCount"`
	Page           int                  `json:"page"`
	PageCount      int                  `json:"pageCount"`
}

// ExpansionReference represents a reference to an expansion
type ExpansionReference struct {
	ID          int       `json:"id"`
	ExpansionID int       `json:"expansionId"`
	ReferenceID int       `json:"referenceId"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListExpansionReferences retrieves a list of expansion references
func (c *Client) ListExpansionReferences(ctx context.Context, params *ListExpansionReferencesParams) (*ListExpansionReferencesResponse, error) {
	path := "/api/expansion-references"
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

	var response ListExpansionReferencesResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetExpansionReference retrieves a single expansion reference by ID
func (c *Client) GetExpansionReference(ctx context.Context, id int) (*ExpansionReference, error) {
	var response ExpansionReference
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/expansion-references/%d", id), nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
