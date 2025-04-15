package tcgcollector

import (
	"context"
	"net/http"
)

// HealthStatus represents the health status of the API
type HealthStatus struct {
	Status    string `json:"status"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}

// GetHealth retrieves the health status of the API
func (c *Client) GetHealth(ctx context.Context) (*HealthStatus, error) {
	var response HealthStatus
	if err := c.doRequest(ctx, http.MethodGet, "/api/health", nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
