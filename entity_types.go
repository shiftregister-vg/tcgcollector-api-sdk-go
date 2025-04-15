package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ListEntityTypesParams contains the parameters for listing entity types
type ListEntityTypesParams struct {
	Page     *int `json:"page,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
}

// ListEntityTypesResponse represents the response from listing entity types
type ListEntityTypesResponse struct {
	Items          []EntityType `json:"items"`
	ItemCount      int          `json:"itemCount"`
	TotalItemCount int          `json:"totalItemCount"`
	Page           int          `json:"page"`
	PageCount      int          `json:"pageCount"`
}

// EntityType represents an entity type
type EntityType struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListEntityTypes retrieves a list of entity types
func (c *Client) ListEntityTypes(ctx context.Context, params *ListEntityTypesParams) (*ListEntityTypesResponse, error) {
	path := "/api/entity-types"
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

	var response ListEntityTypesResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetEntityType retrieves a single entity type by ID
func (c *Client) GetEntityType(ctx context.Context, id int) (*EntityType, error) {
	var response EntityType
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/entity-types/%d", id), nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
