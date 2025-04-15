package tcgcollector

import (
	"context"
	"net/http"
)

// CardDatabaseLogEntry represents a card database log entry
type CardDatabaseLogEntry struct {
	ID        int    `json:"id"`
	Action    string `json:"action"`
	Details   string `json:"details"`
	CreatedAt string `json:"createdAt"`
}

// ListCardDatabaseLogEntries retrieves a list of card database log entries
func (c *Client) ListCardDatabaseLogEntries(ctx context.Context) ([]CardDatabaseLogEntry, error) {
	var response []CardDatabaseLogEntry
	if err := c.doRequest(ctx, http.MethodGet, "/api/card-database-log", nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// PruneCardDatabaseLog prunes old card database log entries
func (c *Client) PruneCardDatabaseLog(ctx context.Context) error {
	return c.doRequest(ctx, http.MethodPost, "/api/card-database-log/prune", nil, nil)
}
