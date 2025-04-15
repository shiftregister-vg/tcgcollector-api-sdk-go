package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ListImagesParams contains the parameters for listing images
type ListImagesParams struct {
	Page     *int `json:"page,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
}

// ListImagesResponse represents the response from listing images
type ListImagesResponse struct {
	Items          []Image `json:"items"`
	ItemCount      int     `json:"itemCount"`
	TotalItemCount int     `json:"totalItemCount"`
	Page           int     `json:"page"`
	PageCount      int     `json:"pageCount"`
}

// CreateImageParams contains the parameters for creating an image
type CreateImageParams struct {
	File []byte `json:"file"`
}

// ListImages retrieves a list of images
func (c *Client) ListImages(ctx context.Context, params *ListImagesParams) (*ListImagesResponse, error) {
	path := "/api/images"
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

	var response ListImagesResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetImage retrieves a single image by ID
func (c *Client) GetImage(ctx context.Context, id int) (*Image, error) {
	var response Image
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/images/%d", id), nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateImage creates a new image
func (c *Client) CreateImage(ctx context.Context, params *CreateImageParams) (*Image, error) {
	var response Image
	if err := c.doRequest(ctx, http.MethodPost, "/api/images", params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteImage deletes an image
func (c *Client) DeleteImage(ctx context.Context, id int) error {
	return c.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/images/%d", id), nil, nil)
}
