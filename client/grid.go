package client

import (
	"net/url"

	"github.com/yehlo/storagegrid-sdk-go/services"
)

const (
	gridAPI = "api/v4"
)

type GridClient struct {
	client *Client

	// Services
	tenant  services.TenantServiceInterface
	health  services.HealthServiceInterface
	region  services.RegionServiceInterface
	haGroup services.HAGroupServiceInterface
	gateway services.GatewayConfigServiceInterface
}

func NewGridClient(options ...ClientOption) (*GridClient, error) {
	c, err := newClient(options...)
	if err != nil {
		return nil, err
	}

	c.baseURL = c.baseURL.ResolveReference(&url.URL{Path: gridAPI})

	return &GridClient{
		client:  c,
		tenant:  services.NewTenantService(c),
		health:  services.NewHealthService(c),
		region:  services.NewRegionGridService(c),
		haGroup: services.NewHAGroupService(c),
		gateway: services.NewGatewayConfigService(c),
	}, nil
}

// Service getters return interfaces to enable testing with mocks

func (gc *GridClient) Tenant() services.TenantServiceInterface {
	return gc.tenant
}

func (gc *GridClient) Health() services.HealthServiceInterface {
	return gc.health
}

func (gc *GridClient) Region() services.RegionServiceInterface {
	return gc.region
}

func (gc *GridClient) HAGroup() services.HAGroupServiceInterface {
	return gc.haGroup
}

func (gc *GridClient) Gateway() services.GatewayConfigServiceInterface {
	return gc.gateway
}
