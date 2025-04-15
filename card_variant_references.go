package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ListCardVariantReferencesParams contains the parameters for listing card variant references
type ListCardVariantReferencesParams struct {
	Page     *int `json:"page,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
}

// ListCardVariantReferencesResponse represents the response from listing card variant references
type ListCardVariantReferencesResponse struct {
	Items          []CardVariantReference `json:"items"`
	ItemCount      int                    `json:"itemCount"`
	TotalItemCount int                    `json:"totalItemCount"`
	Page           int                    `json:"page"`
	PageCount      int                    `json:"pageCount"`
}

// CardVariantReference represents a reference to a card variant
type CardVariantReference struct {
	ID            int       `json:"id"`
	CardVariantID int       `json:"cardVariantId"`
	ReferenceID   int       `json:"referenceId"`
	Type          string    `json:"type"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// ListCardVariantReferences retrieves a list of card variant references
func (c *Client) ListCardVariantReferences(ctx context.Context, params *ListCardVariantReferencesParams) (*ListCardVariantReferencesResponse, error) {
	path := "/api/card-variant-references"
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

	var response ListCardVariantReferencesResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCardVariantReference retrieves a single card variant reference by ID
func (c *Client) GetCardVariantReference(ctx context.Context, id int) (*CardVariantReference, error) {
	var response CardVariantReference
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-variant-references/%d", id), nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
