package client

import (
	"net/url"

	"github.com/yehlo/storagegrid-sdk-go/services/gateway"
	"github.com/yehlo/storagegrid-sdk-go/services/hagroup"
	"github.com/yehlo/storagegrid-sdk-go/services/health"
	"github.com/yehlo/storagegrid-sdk-go/services/region"
	"github.com/yehlo/storagegrid-sdk-go/services/tenant"
	"github.com/yehlo/storagegrid-sdk-go/services/trafficclass"
)

const (
	gridAPI = "api/v4"
)

// GridClient is the base struct used to interface with the storagegrid API
type GridClient struct {
	client *Client

	// Services
	tenant       tenant.ServiceInterface
	health       health.ServiceInterface
	region       region.ServiceInterface
	haGroup      hagroup.ServiceInterface
	gateway      gateway.ServiceInterface
	trafficClass trafficclass.ServiceInterface
}

func NewGridClient(options ...Option) (*GridClient, error) {
	c, err := newClient(options...)
	if err != nil {
		return nil, err
	}

	c.baseURL = c.baseURL.ResolveReference(&url.URL{Path: gridAPI})

	return &GridClient{
		client:       c,
		gateway:      gateway.NewService(c),
		haGroup:      hagroup.NewService(c),
		health:       health.NewService(c),
		region:       region.NewGridService(c),
		tenant:       tenant.NewService(c),
		trafficClass: trafficclass.NewService(c),
	}, nil
}

// Service getters return interfaces to enable testing with mocks
func (gc *GridClient) TrafficClass() trafficclass.ServiceInterface {
	return gc.trafficClass
}

func (gc *GridClient) Tenant() tenant.ServiceInterface {
	return gc.tenant
}

func (gc *GridClient) Health() health.ServiceInterface {
	return gc.health
}

func (gc *GridClient) Region() region.ServiceInterface {
	return gc.region
}

func (gc *GridClient) HAGroup() hagroup.ServiceInterface {
	return gc.haGroup
}

func (gc *GridClient) Gateway() gateway.ServiceInterface {
	return gc.gateway
}
