package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ListCardVariantsParams represents the parameters for listing card variants
type ListCardVariantsParams struct {
	CardID   *int
	TypeID   *int
	Page     *int
	PageSize *int
}

// ListCardVariants lists card variants with optional filtering
func (c *Client) ListCardVariants(ctx context.Context, params *ListCardVariantsParams) (*ListResponse[CardVariant], error) {
	query := url.Values{}
	if params != nil {
		if params.CardID != nil {
			query.Add("cardId", fmt.Sprintf("%d", *params.CardID))
		}
		if params.TypeID != nil {
			query.Add("typeId", fmt.Sprintf("%d", *params.TypeID))
		}
		if params.Page != nil {
			query.Add("page", fmt.Sprintf("%d", *params.Page))
		}
		if params.PageSize != nil {
			query.Add("pageSize", fmt.Sprintf("%d", *params.PageSize))
		}
	}

	path := "/api/card-variants"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	var result ListResponse[CardVariant]
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCardVariant gets a single card variant by ID
func (c *Client) GetCardVariant(ctx context.Context, id int) (*CardVariant, error) {
	var result CardVariant
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-variants/%d", id), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateCardVariant creates a new card variant
func (c *Client) CreateCardVariant(ctx context.Context, variant *CardVariant) (*CardVariant, error) {
	var result CardVariant
	if err := c.doRequest(ctx, http.MethodPost, "/api/card-variants", variant, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateCardVariant updates an existing card variant
func (c *Client) UpdateCardVariant(ctx context.Context, id int, variant *CardVariant) (*CardVariant, error) {
	var result CardVariant
	if err := c.doRequest(ctx, http.MethodPut, fmt.Sprintf("/api/card-variants/%d", id), variant, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteCardVariant deletes a card variant
func (c *Client) DeleteCardVariant(ctx context.Context, id int) error {
	return c.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/card-variants/%d", id), nil, nil)
}

// GetCardVariantPrices gets the price history for a card variant
func (c *Client) GetCardVariantPrices(ctx context.Context, variantID int) (*ListResponse[CardVariantPrice], error) {
	var result ListResponse[CardVariantPrice]
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-variants/%d/prices", variantID), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RecalculateComputedAndCachedValues recalculates computed and cached values for all card variants
func (c *Client) RecalculateComputedAndCachedValues(ctx context.Context) error {
	return c.doRequest(ctx, http.MethodPost, "/api/card-variants/recalculate-computed-and-cached-values", nil, nil)
}
