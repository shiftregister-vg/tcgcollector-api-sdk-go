package tcgcollector

import (
	"context"
	"net/http"
)

// AllowedExternalAccountHosts represents the response from getting allowed external account hosts
type AllowedExternalAccountHosts struct {
	Hosts []string `json:"hosts"`
}

// BaseTCGCurrency represents the response from getting base TCG currency
type BaseTCGCurrency struct {
	Currency string `json:"currency"`
}

// GetAllowedExternalAccountHosts retrieves the list of allowed external account hosts
func (c *Client) GetAllowedExternalAccountHosts(ctx context.Context) (*AllowedExternalAccountHosts, error) {
	var response AllowedExternalAccountHosts
	if err := c.doRequest(ctx, http.MethodGet, "/api/configuration/allowed-external-account-hosts", nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetBaseTCGCurrency retrieves the base TCG currency
func (c *Client) GetBaseTCGCurrency(ctx context.Context) (*BaseTCGCurrency, error) {
	var response BaseTCGCurrency
	if err := c.doRequest(ctx, http.MethodGet, "/api/configuration/base-tcg-currency", nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
