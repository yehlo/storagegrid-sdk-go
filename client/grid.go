package client

import (
	"net/url"

	services "github.com/yehlo/storagegrid-sdk-go/services"
)

const (
	gridAPI = "api/v4"
)

type GridClient struct {
	client *Client

	// Services
	Tenant *services.TenantService
	Health *services.HealthService
	Region *services.RegionService
}

func NewGridClient(options ...ClientOption) (*GridClient, error) {
	c, err := newClient(options...)
	if err != nil {
		return nil, err
	}

	c.baseURL = c.baseURL.ResolveReference(&url.URL{Path: gridAPI})

	return &GridClient{
		client: c,
		Tenant: services.NewTenantService(c),
		Health: services.NewHealthService(c),
		Region: services.NewRegionGridService(c),
	}, nil
}
