package tcgcollector

import (
	"context"
	"net/http"
)

// GetStatistics retrieves the statistics for the API
func (c *Client) GetStatistics(ctx context.Context) (*UserStatistics, error) {
	var response UserStatistics
	if err := c.doRequest(ctx, http.MethodGet, "/api/statistics", nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
