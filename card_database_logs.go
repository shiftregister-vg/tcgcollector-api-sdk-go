package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ListCardDatabaseLogsParams contains the parameters for listing card database logs
type ListCardDatabaseLogsParams struct {
	Page     *int `json:"page,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
}

// ListCardDatabaseLogsResponse represents the response from listing card database logs
type ListCardDatabaseLogsResponse struct {
	Items          []CardDatabaseLog `json:"items"`
	ItemCount      int               `json:"itemCount"`
	TotalItemCount int               `json:"totalItemCount"`
	Page           int               `json:"page"`
	PageCount      int               `json:"pageCount"`
}

// CardDatabaseLog represents a log entry in the card database
type CardDatabaseLog struct {
	ID        int       `json:"id"`
	CardID    int       `json:"cardId"`
	UserID    int       `json:"userId"`
	Action    string    `json:"action"`
	Details   string    `json:"details"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ListCardDatabaseLogs retrieves a list of card database logs
func (c *Client) ListCardDatabaseLogs(ctx context.Context, params *ListCardDatabaseLogsParams) (*ListCardDatabaseLogsResponse, error) {
	path := "/api/card-database-logs"
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

	var response ListCardDatabaseLogsResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCardDatabaseLog retrieves a single card database log by ID
func (c *Client) GetCardDatabaseLog(ctx context.Context, id int) (*CardDatabaseLog, error) {
	var response CardDatabaseLog
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/card-database-logs/%d", id), nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
