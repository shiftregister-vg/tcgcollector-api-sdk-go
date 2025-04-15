package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ListCardsParams represents the parameters for listing cards
type ListCardsParams struct {
	SetID    *int
	Name     *string
	Number   *string
	Rarity   *string
	Page     *int
	PageSize *int
}

// ListCards lists cards with optional filtering
func (c *Client) ListCards(ctx context.Context, params *ListCardsParams) (*ListResponse[Card], error) {
	query := url.Values{}
	if params != nil {
		if params.SetID != nil {
			query.Add("setId", fmt.Sprintf("%d", *params.SetID))
		}
		if params.Name != nil {
			query.Add("name", *params.Name)
		}
		if params.Number != nil {
			query.Add("number", *params.Number)
		}
		if params.Rarity != nil {
			query.Add("rarity", *params.Rarity)
		}
		if params.Page != nil {
			query.Add("page", fmt.Sprintf("%d", *params.Page))
		}
		if params.PageSize != nil {
			query.Add("pageSize", fmt.Sprintf("%d", *params.PageSize))
		}
	}

	path := "/api/cards"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	var result ListResponse[Card]
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCard gets a single card by ID
func (c *Client) GetCard(ctx context.Context, id int) (*Card, error) {
	var result Card
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/cards/%d", id), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCardPrices gets the price history for a card
func (c *Client) GetCardPrices(ctx context.Context, cardID int) (*ListResponse[CardPrice], error) {
	var result ListResponse[CardPrice]
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/cards/%d/prices", cardID), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RecalculateCachedValues recalculates cached values for all cards
func (c *Client) RecalculateCachedValues(ctx context.Context) error {
	return c.doRequest(ctx, http.MethodPost, "/api/cards/recalculate-cached-values", nil, nil)
}

// RegenerateSlugs regenerates slugs for all cards
func (c *Client) RegenerateSlugs(ctx context.Context) error {
	return c.doRequest(ctx, http.MethodPost, "/api/cards/regenerate-slugs", nil, nil)
}

// RegenerateSurrogateNumbersAndFullNames regenerates surrogate numbers and full names for all cards
func (c *Client) RegenerateSurrogateNumbersAndFullNames(ctx context.Context) error {
	return c.doRequest(ctx, http.MethodPost, "/api/cards/regenerate-surrogate-numbers-and-full-names", nil, nil)
}
