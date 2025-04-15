package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CardVariantType represents a card variant type in the TCG Collector API
type CardVariantType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// ListCardVariantTypesParams contains the parameters for listing card variant types
type ListCardVariantTypesParams struct {
	Page     *int `json:"page,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
}

// ListCardVariantTypesResponse represents the response from listing card variant types
type ListCardVariantTypesResponse struct {
	Items          []CardVariantType `json:"items"`
	ItemCount      int               `json:"itemCount"`
	TotalItemCount int               `json:"totalItemCount"`
	Page           int               `json:"page"`
	PageCount      int               `json:"pageCount"`
}

// ListCardVariantTypes retrieves a list of card variant types
func (c *Client) ListCardVariantTypes(ctx context.Context, params *ListCardVariantTypesParams) (*ListCardVariantTypesResponse, error) {
	path := "/api/card-variant-types"
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

	var response ListCardVariantTypesResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCardVariantType retrieves a single card variant type by ID
func (c *Client) GetCardVariantType(ctx context.Context, id int) (*CardVariantType, error) {
	var response CardVariantType
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-variant-types/%d", id), nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateCardVariantType creates a new card variant type
func (c *Client) CreateCardVariantType(ctx context.Context, variantType *CardVariantType) (*CardVariantType, error) {
	var response CardVariantType
	if err := c.doRequest(ctx, http.MethodPost, "/api/card-variant-types", variantType, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateCardVariantType updates an existing card variant type
func (c *Client) UpdateCardVariantType(ctx context.Context, id int, variantType *CardVariantType) (*CardVariantType, error) {
	var response CardVariantType
	if err := c.doRequest(ctx, http.MethodPut, fmt.Sprintf("/api/card-variant-types/%d", id), variantType, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteCardVariantType deletes a card variant type
func (c *Client) DeleteCardVariantType(ctx context.Context, id int) error {
	return c.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/card-variant-types/%d", id), nil, nil)
}
