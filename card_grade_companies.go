package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// CardGradeCompany represents a card grading company
type CardGradeCompany struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Website     string    `json:"website"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ListCardGradeCompanies retrieves a list of card grading companies
func (c *Client) ListCardGradeCompanies(ctx context.Context) ([]CardGradeCompany, error) {
	var response []CardGradeCompany
	if err := c.doRequest(ctx, http.MethodGet, "/api/card-grade-companies", nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetCardGradeCompany retrieves a single card grading company by ID
func (c *Client) GetCardGradeCompany(ctx context.Context, id int) (*CardGradeCompany, error) {
	var response CardGradeCompany
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-grade-companies/%d", id), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
