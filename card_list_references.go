package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ListCardListReferencesParams contains the parameters for listing card list references
type ListCardListReferencesParams struct {
	Page     *int `json:"page,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
}

// ListCardListReferencesResponse represents the response from listing card list references
type ListCardListReferencesResponse struct {
	Items          []CardListReference `json:"items"`
	ItemCount      int                 `json:"itemCount"`
	TotalItemCount int                 `json:"totalItemCount"`
	Page           int                 `json:"page"`
	PageCount      int                 `json:"pageCount"`
}

// CardListReference represents a reference to a card list
type CardListReference struct {
	ID          int       `json:"id"`
	CardListID  int       `json:"cardListId"`
	ReferenceID int       `json:"referenceId"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListCardListReferences retrieves a list of card list references
func (c *Client) ListCardListReferences(ctx context.Context, params *ListCardListReferencesParams) (*ListCardListReferencesResponse, error) {
	path := "/api/card-list-references"
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

	var response ListCardListReferencesResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCardListReference retrieves a single card list reference by ID
func (c *Client) GetCardListReference(ctx context.Context, id int) (*CardListReference, error) {
	var response CardListReference
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-list-references/%d", id), nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
