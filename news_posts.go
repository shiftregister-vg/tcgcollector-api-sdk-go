package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ListNewsPostsParams contains the parameters for listing news posts
type ListNewsPostsParams struct {
	Page     int `url:"page,omitempty"`
	PageSize int `url:"pageSize,omitempty"`
}

// ListNewsPostsResponse contains the response data for listing news posts
type ListNewsPostsResponse struct {
	Items []NewsPost `json:"items"`
	Total int        `json:"total"`
}

// ListNewsPosts retrieves a list of news posts with optional pagination
func (c *Client) ListNewsPosts(ctx context.Context, params *ListNewsPostsParams) (*ListNewsPostsResponse, error) {
	path := "/api/news-posts"
	if params != nil {
		query := url.Values{}
		if params.Page > 0 {
			query.Set("page", strconv.Itoa(params.Page))
		}
		if params.PageSize > 0 {
			query.Set("pageSize", strconv.Itoa(params.PageSize))
		}
		if len(query) > 0 {
			path = fmt.Sprintf("%s?%s", path, query.Encode())
		}
	}

	var result ListNewsPostsResponse
	err := c.doRequest(ctx, http.MethodGet, path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetNewsPost retrieves a specific news post by ID
func (c *Client) GetNewsPost(ctx context.Context, id int) (*NewsPost, error) {
	var result NewsPost
	err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/news-posts/%d", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateNewsPostRequest contains the parameters for creating a news post
type CreateNewsPostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// CreateNewsPost creates a new news post
func (c *Client) CreateNewsPost(ctx context.Context, request *CreateNewsPostRequest) (*NewsPost, error) {
	var result NewsPost
	err := c.doRequest(ctx, http.MethodPost, "/api/news-posts", request, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateNewsPostRequest contains the parameters for updating a news post
type UpdateNewsPostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// UpdateNewsPost updates an existing news post
func (c *Client) UpdateNewsPost(ctx context.Context, id int, request *UpdateNewsPostRequest) (*NewsPost, error) {
	var result NewsPost
	err := c.doRequest(ctx, http.MethodPut, fmt.Sprintf("/api/news-posts/%d", id), request, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteNewsPost deletes a news post by ID
func (c *Client) DeleteNewsPost(ctx context.Context, id int) error {
	return c.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/news-posts/%d", id), nil, nil)
}
