package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ListUsersParams contains the parameters for listing users
type ListUsersParams struct {
	Page     *int    `json:"page,omitempty"`
	PageSize *int    `json:"pageSize,omitempty"`
	Search   *string `json:"search,omitempty"`
}

// ListUsersResponse represents the response from listing users
type ListUsersResponse struct {
	Items          []User `json:"items"`
	ItemCount      int    `json:"itemCount"`
	TotalItemCount int    `json:"totalItemCount"`
	Page           int    `json:"page"`
	PageCount      int    `json:"pageCount"`
}

// CreateUserParams contains the parameters for creating a user
type CreateUserParams struct {
	DisplayName  string `json:"displayName"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

// UpdateUserParams contains the parameters for updating a user
type UpdateUserParams struct {
	DisplayName  *string `json:"displayName,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
	Password     *string `json:"password,omitempty"`
}

// ListUsers retrieves a list of users
func (c *Client) ListUsers(ctx context.Context, params *ListUsersParams) (*ListUsersResponse, error) {
	path := "/api/users"
	if params != nil {
		query := url.Values{}
		if params.Page != nil {
			query.Set("page", fmt.Sprintf("%d", *params.Page))
		}
		if params.PageSize != nil {
			query.Set("pageSize", fmt.Sprintf("%d", *params.PageSize))
		}
		if params.Search != nil {
			query.Set("search", *params.Search)
		}
		if len(query) > 0 {
			path = fmt.Sprintf("%s?%s", path, query.Encode())
		}
	}

	var response ListUsersResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetUser retrieves a single user by ID
func (c *Client) GetUser(ctx context.Context, id int) (*User, error) {
	var response User
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/users/%d", id), nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateUser creates a new user
func (c *Client) CreateUser(ctx context.Context, params *CreateUserParams) (*User, error) {
	var response User
	if err := c.doRequest(ctx, http.MethodPost, "/api/users", params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateUser updates an existing user
func (c *Client) UpdateUser(ctx context.Context, id int, params *UpdateUserParams) (*User, error) {
	var response User
	if err := c.doRequest(ctx, http.MethodPut, fmt.Sprintf("/api/users/%d", id), params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteUser deletes a user
func (c *Client) DeleteUser(ctx context.Context, id int) error {
	return c.doRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/users/%d", id), nil, nil)
}

// GetCurrentUser retrieves the currently authenticated user
func (c *Client) GetCurrentUser(ctx context.Context) (*User, error) {
	var response User
	if err := c.doRequest(ctx, http.MethodGet, "/api/users/me", nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateCurrentUser updates the currently authenticated user
func (c *Client) UpdateCurrentUser(ctx context.Context, params *UpdateUserParams) (*User, error) {
	var response User
	if err := c.doRequest(ctx, http.MethodPut, "/api/users/me", params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteCurrentUser deletes the currently authenticated user
func (c *Client) DeleteCurrentUser(ctx context.Context) error {
	return c.doRequest(ctx, http.MethodDelete, "/api/users/me", nil, nil)
}

// GetUserPreferences gets a user's preferences
func (c *Client) GetUserPreferences(ctx context.Context, userID int) (*UserPreferences, error) {
	var result UserPreferences
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/users/%d/preferences", userID), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateUserPreferences updates a user's preferences
func (c *Client) UpdateUserPreferences(ctx context.Context, userID int, preferences *UserPreferences) (*UserPreferences, error) {
	var result UserPreferences
	if err := c.doRequest(ctx, http.MethodPut, fmt.Sprintf("/api/users/%d/preferences", userID), preferences, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserCount retrieves the total count of users
func (c *Client) GetUserCount(ctx context.Context) (int, error) {
	var response struct {
		Count int `json:"count"`
	}
	if err := c.doRequest(ctx, http.MethodGet, "/api/users/count", nil, &response); err != nil {
		return 0, err
	}
	return response.Count, nil
}

// PruneActivityLogs prunes user activity logs
func (c *Client) PruneActivityLogs(ctx context.Context) error {
	return c.doRequest(ctx, http.MethodPost, "/api/users/prune-activity-logs", nil, nil)
}

// DisableUserPremium disables premium features for a user
func (c *Client) DisableUserPremium(ctx context.Context, userID int) error {
	return c.doRequest(ctx, http.MethodPost, fmt.Sprintf("/api/users/%d/disable-premium", userID), nil, nil)
}

// EnableUserPremiumWithoutSubscription enables premium features for a user without requiring a subscription
func (c *Client) EnableUserPremiumWithoutSubscription(ctx context.Context, userID int) error {
	return c.doRequest(ctx, http.MethodPost, fmt.Sprintf("/api/users/%d/enable-premium-without-subscription", userID), nil, nil)
}

// GenerateAPIAccessToken generates a new API access token for a user
func (c *Client) GenerateAPIAccessToken(ctx context.Context, userID int) (string, error) {
	var response struct {
		Token string `json:"token"`
	}
	if err := c.doRequest(ctx, http.MethodPost, fmt.Sprintf("/api/users/%d/generate-api-access-token", userID), nil, &response); err != nil {
		return "", err
	}
	return response.Token, nil
}

// GetUserPermissions retrieves the permissions for a user
func (c *Client) GetUserPermissions(ctx context.Context, userID int) ([]string, error) {
	var response struct {
		Permissions []string `json:"permissions"`
	}
	if err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/api/users/%d/permissions", userID), nil, &response); err != nil {
		return nil, err
	}
	return response.Permissions, nil
}

// RevokeAPIAccessToken revokes the API access token for a user
func (c *Client) RevokeAPIAccessToken(ctx context.Context, userID int) error {
	return c.doRequest(ctx, http.MethodPost, fmt.Sprintf("/api/users/%d/revoke-api-access-token", userID), nil, nil)
}
