package client

import (
	"net/url"
)

const (
	tenantAPI = "api/v4"
)

type TenantClient struct {
	client *Client

	// Services
	// Tenant *services.TenantService
}

func NewTenantClient(options ...ClientOption) (*TenantClient, error) {
	c, err := newClient(options...)
	if err != nil {
		return nil, err
	}

	c.baseURL = c.baseURL.ResolveReference(&url.URL{Path: tenantAPI})

	return &TenantClient{
		client: c,
	}, nil
}
