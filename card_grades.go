package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// CardGrade represents a graded card
type CardGrade struct {
	ID             int       `json:"id"`
	CardID         int       `json:"cardId"`
	GradeCompanyID int       `json:"gradeCompanyId"`
	GradeValue     string    `json:"gradeValue"`
	CertificateID  string    `json:"certificateId"`
	GradedAt       time.Time `json:"gradedAt"`
	Notes          string    `json:"notes,omitempty"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// ListCardGradesParams contains the parameters for listing card grades
type ListCardGradesParams struct {
	CardID         *int    `json:"cardId,omitempty"`
	GradeCompanyID *int    `json:"gradeCompanyId,omitempty"`
	GradeValue     *string `json:"gradeValue,omitempty"`
	Page           *int    `json:"page,omitempty"`
	PageSize       *int    `json:"pageSize,omitempty"`
}

// ListCardGrades retrieves a list of card grades with optional filtering
func (c *Client) ListCardGrades(ctx context.Context, params *ListCardGradesParams) (*ListResponse[CardGrade], error) {
	query := url.Values{}
	if params != nil {
		if params.CardID != nil {
			query.Add("cardId", fmt.Sprintf("%d", *params.CardID))
		}
		if params.GradeCompanyID != nil {
			query.Add("gradeCompanyId", fmt.Sprintf("%d", *params.GradeCompanyID))
		}
		if params.GradeValue != nil {
			query.Add("gradeValue", *params.GradeValue)
		}
		if params.Page != nil {
			query.Add("page", fmt.Sprintf("%d", *params.Page))
		}
		if params.PageSize != nil {
			query.Add("pageSize", fmt.Sprintf("%d", *params.PageSize))
		}
	}

	path := "/api/card-grades"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	var result ListResponse[CardGrade]
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCardGrade retrieves a single card grade by ID
func (c *Client) GetCardGrade(ctx context.Context, id int) (*CardGrade, error) {
	var result CardGrade
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-grades/%d", id), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateCardGrade creates a new card grade
func (c *Client) CreateCardGrade(ctx context.Context, grade *CardGrade) (*CardGrade, error) {
	var result CardGrade
	if err := c.doRequest(ctx, http.MethodPost, "/api/card-grades", grade, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateCardGrade updates an existing card grade
func (c *Client) UpdateCardGrade(ctx context.Context, id int, grade *CardGrade) (*CardGrade, error) {
	var result CardGrade
	if err := c.doRequest(ctx, http.MethodPut, fmt.Sprintf("/api/card-grades/%d", id), grade, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteCardGrade deletes a card grade
func (c *Client) DeleteCardGrade(ctx context.Context, id int) error {
	return c.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/card-grades/%d", id), nil, nil)
}
