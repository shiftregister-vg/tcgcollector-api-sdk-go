package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// EnergyType represents an energy type
type EnergyType struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Symbol      string    `json:"symbol"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListEnergyTypes retrieves a list of energy types
func (c *Client) ListEnergyTypes(ctx context.Context) ([]EnergyType, error) {
	var response []EnergyType
	if err := c.doRequest(ctx, http.MethodGet, "/api/energy-types", nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetEnergyType retrieves a single energy type by ID
func (c *Client) GetEnergyType(ctx context.Context, id int) (*EnergyType, error) {
	var response EnergyType
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/energy-types/%d", id), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
