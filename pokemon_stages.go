package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ListPokemonStagesParams contains the parameters for listing Pokémon stages
type ListPokemonStagesParams struct {
	Page     *int `json:"page,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
}

// ListPokemonStagesResponse represents the response from listing Pokémon stages
type ListPokemonStagesResponse struct {
	Items          []PokemonStage `json:"items"`
	ItemCount      int            `json:"itemCount"`
	TotalItemCount int            `json:"totalItemCount"`
	Page           int            `json:"page"`
	PageCount      int            `json:"pageCount"`
}

// PokemonStage represents a Pokémon stage
type PokemonStage struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListPokemonStages retrieves a list of Pokémon stages
func (c *Client) ListPokemonStages(ctx context.Context, params *ListPokemonStagesParams) (*ListPokemonStagesResponse, error) {
	path := "/api/pokemon-stages"
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

	var response ListPokemonStagesResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetPokemonStage retrieves a single Pokémon stage by ID
func (c *Client) GetPokemonStage(ctx context.Context, id int) (*PokemonStage, error) {
	var response PokemonStage
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/pokemon-stages/%d", id), nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
