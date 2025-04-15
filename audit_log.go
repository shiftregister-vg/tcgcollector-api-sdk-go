package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ListAuditLogEntriesParams represents the parameters for listing audit log entries
type ListAuditLogEntriesParams struct {
	EventTypeID *int
	UserID      *int
	StartDate   *time.Time
	EndDate     *time.Time
	Page        *int
	PageSize    *int
}

// ListAuditLogEntries lists audit log entries with optional filtering
func (c *Client) ListAuditLogEntries(ctx context.Context, params *ListAuditLogEntriesParams) (*ListResponse[AuditLogEntry], error) {
	// Build query parameters
	query := url.Values{}
	if params != nil {
		if params.EventTypeID != nil {
			query.Add("eventTypeId", fmt.Sprintf("%d", *params.EventTypeID))
		}
		if params.UserID != nil {
			query.Add("userId", fmt.Sprintf("%d", *params.UserID))
		}
		if params.StartDate != nil {
			query.Add("startDate", params.StartDate.Format(time.RFC3339))
		}
		if params.EndDate != nil {
			query.Add("endDate", params.EndDate.Format(time.RFC3339))
		}
		if params.Page != nil {
			query.Add("page", fmt.Sprintf("%d", *params.Page))
		}
		if params.PageSize != nil {
			query.Add("pageSize", fmt.Sprintf("%d", *params.PageSize))
		}
	}

	// Build path with query parameters
	path := "/api/audit-log"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	var result ListResponse[AuditLogEntry]
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAuditLogEntry gets a single audit log entry by ID
func (c *Client) GetAuditLogEntry(ctx context.Context, id int) (*AuditLogEntry, error) {
	var result AuditLogEntry
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/audit-log/%d", id), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
