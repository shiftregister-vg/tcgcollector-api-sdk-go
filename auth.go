package tcgcollector

import (
	"context"
	"net/http"
)

// Login authenticates a user and returns a JWT token
func (c *Client) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	var response LoginResponse
	if err := c.doRequest(ctx, http.MethodPost, "/api/auth/login", request, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Register creates a new user account
func (c *Client) Register(ctx context.Context, request *RegisterRequest) (*RegisterResponse, error) {
	var response RegisterResponse
	if err := c.doRequest(ctx, http.MethodPost, "/api/auth/register", request, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Logout invalidates the current JWT token
func (c *Client) Logout(ctx context.Context) error {
	return c.doRequest(ctx, http.MethodPost, "/api/auth/logout", nil, nil)
}

// RefreshToken refreshes the current JWT token
func (c *Client) RefreshToken(ctx context.Context) (*LoginResponse, error) {
	var response LoginResponse
	if err := c.doRequest(ctx, http.MethodPost, "/api/auth/refresh", nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
