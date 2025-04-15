package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ListCollectionsParams represents the parameters for listing collections
type ListCollectionsParams struct {
	UserID   *int
	Name     *string
	IsPublic *bool
	Page     *int
	PageSize *int
}

// ListCollections lists collections with optional filtering
func (c *Client) ListCollections(ctx context.Context, params *ListCollectionsParams) (*ListResponse[Collection], error) {
	query := url.Values{}
	if params != nil {
		if params.UserID != nil {
			query.Add("userId", fmt.Sprintf("%d", *params.UserID))
		}
		if params.Name != nil {
			query.Add("name", *params.Name)
		}
		if params.IsPublic != nil {
			query.Add("isPublic", fmt.Sprintf("%v", *params.IsPublic))
		}
		if params.Page != nil {
			query.Add("page", fmt.Sprintf("%d", *params.Page))
		}
		if params.PageSize != nil {
			query.Add("pageSize", fmt.Sprintf("%d", *params.PageSize))
		}
	}

	path := "/api/collections"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	var result ListResponse[Collection]
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCollection gets a single collection by ID
func (c *Client) GetCollection(ctx context.Context, id int) (*Collection, error) {
	var result Collection
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/collections/%d", id), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateCollection creates a new collection
func (c *Client) CreateCollection(ctx context.Context, collection *Collection) (*Collection, error) {
	var result Collection
	if err := c.doRequest(ctx, http.MethodPost, "/api/collections", collection, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateCollection updates an existing collection
func (c *Client) UpdateCollection(ctx context.Context, id int, collection *Collection) (*Collection, error) {
	var result Collection
	if err := c.doRequest(ctx, http.MethodPut, fmt.Sprintf("/api/collections/%d", id), collection, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteCollection deletes a collection
func (c *Client) DeleteCollection(ctx context.Context, id int) error {
	return c.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/collections/%d", id), nil, nil)
}

// ListCollectionCards lists all cards in a collection
func (c *Client) ListCollectionCards(ctx context.Context, collectionID int) (*ListResponse[CollectionCard], error) {
	var result ListResponse[CollectionCard]
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/collections/%d/cards", collectionID), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AddCardToCollection adds a card to a collection
func (c *Client) AddCardToCollection(ctx context.Context, collectionID int, card *CollectionCard) (*CollectionCard, error) {
	var result CollectionCard
	if err := c.doRequest(ctx, http.MethodPost, fmt.Sprintf("/api/collections/%d/cards", collectionID), card, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateCollectionCard updates a card in a collection
func (c *Client) UpdateCollectionCard(ctx context.Context, collectionID, cardID int, card *CollectionCard) (*CollectionCard, error) {
	var result CollectionCard
	if err := c.doRequest(ctx, http.MethodPut, fmt.Sprintf("/api/collections/%d/cards/%d", collectionID, cardID), card, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RemoveCardFromCollection removes a card from a collection
func (c *Client) RemoveCardFromCollection(ctx context.Context, collectionID, cardID int) error {
	return c.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/collections/%d/cards/%d", collectionID, cardID), nil, nil)
}

// InvalidateCardListCache invalidates the card list cache
func (c *Client) InvalidateCardListCache(ctx context.Context) error {
	return c.doRequest(ctx, http.MethodPost, "/api/card-collection/invalidate-card-list-cache", nil, nil)
}

// InvalidateExpansionCache invalidates the expansion cache
func (c *Client) InvalidateExpansionCache(ctx context.Context) error {
	return c.doRequest(ctx, http.MethodPost, "/api/card-collection/invalidate-expansion-cache", nil, nil)
}
