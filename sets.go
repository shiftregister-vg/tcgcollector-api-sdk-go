package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ListSetsParams represents the parameters for listing sets
type ListSetsParams struct {
	Name        *string
	Code        *string
	ReleaseDate *time.Time
	Page        *int
	PageSize    *int
}

// ListSets lists sets with optional filtering
func (c *Client) ListSets(ctx context.Context, params *ListSetsParams) (*ListResponse[Set], error) {
	query := url.Values{}
	if params != nil {
		if params.Name != nil {
			query.Add("name", *params.Name)
		}
		if params.Code != nil {
			query.Add("code", *params.Code)
		}
		if params.ReleaseDate != nil {
			query.Add("releaseDate", params.ReleaseDate.Format(time.RFC3339))
		}
		if params.Page != nil {
			query.Add("page", fmt.Sprintf("%d", *params.Page))
		}
		if params.PageSize != nil {
			query.Add("pageSize", fmt.Sprintf("%d", *params.PageSize))
		}
	}

	path := "/api/sets"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	var result ListResponse[Set]
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSet gets a single set by ID
func (c *Client) GetSet(ctx context.Context, id int) (*Set, error) {
	var result Set
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/sets/%d", id), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSetCards gets all cards in a set
func (c *Client) GetSetCards(ctx context.Context, setID int) (*ListResponse[Card], error) {
	var result ListResponse[Card]
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/sets/%d/cards", setID), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
