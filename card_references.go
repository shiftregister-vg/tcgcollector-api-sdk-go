package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ListCardReferencesParams contains the parameters for listing card references
type ListCardReferencesParams struct {
	Page     *int `json:"page,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
}

// ListCardReferencesResponse represents the response from listing card references
type ListCardReferencesResponse struct {
	Items          []CardReference `json:"items"`
	ItemCount      int             `json:"itemCount"`
	TotalItemCount int             `json:"totalItemCount"`
	Page           int             `json:"page"`
	PageCount      int             `json:"pageCount"`
}

// CardReference represents a reference to a card
type CardReference struct {
	ID          int       `json:"id"`
	CardID      int       `json:"cardId"`
	ReferenceID int       `json:"referenceId"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListCardReferences retrieves a list of card references
func (c *Client) ListCardReferences(ctx context.Context, params *ListCardReferencesParams) (*ListCardReferencesResponse, error) {
	path := "/api/card-references"
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

	var response ListCardReferencesResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCardReference retrieves a single card reference by ID
func (c *Client) GetCardReference(ctx context.Context, id int) (*CardReference, error) {
	var response CardReference
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-references/%d", id), nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
