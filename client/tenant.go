package client

import (
	"net/url"

	"github.com/yehlo/storagegrid-sdk-go/services/accesskeys"
	"github.com/yehlo/storagegrid-sdk-go/services/bucket"
	"github.com/yehlo/storagegrid-sdk-go/services/region"
	"github.com/yehlo/storagegrid-sdk-go/services/tenantgroup"
	"github.com/yehlo/storagegrid-sdk-go/services/tenantuser"
)

const (
	tenantAPI = "api/v4"
)

// TenantClient is the base struct used to interface with the storagegrid tenant API
type TenantClient struct {
	client *Client

	// Services

	bucket       bucket.ServiceInterface
	s3AccessKeys accesskeys.ServiceInterface
	users        tenantuser.ServiceInterface
	groups       tenantgroup.ServiceInterface
	region       region.ServiceInterface
}

// NewTenantClient returns a new tenant client according to the options provided
func NewTenantClient(options ...Option) (*TenantClient, error) {
	c, err := newClient(options...)
	if err != nil {
		return nil, err
	}

	c.baseURL = c.baseURL.ResolveReference(&url.URL{Path: tenantAPI})

	return &TenantClient{
		client:       c,
		bucket:       bucket.NewService(c),
		s3AccessKeys: accesskeys.NewService(c),
		users:        tenantuser.NewService(c),
		groups:       tenantgroup.NewService(c),
		region:       region.NewTenantService(c),
	}, nil
}

// Service getters return interfaces to enable testing with mocks

// Bucket returns the bucket service interface
func (tc *TenantClient) Bucket() bucket.ServiceInterface {
	return tc.bucket
}

// S3AccessKeys returns the S3 access key service interface
func (tc *TenantClient) S3AccessKeys() accesskeys.ServiceInterface {
	return tc.s3AccessKeys
}

// Users returns the tenant user service interface
func (tc *TenantClient) Users() tenantuser.ServiceInterface {
	return tc.users
}

// Groups returns the tenant group service interface
func (tc *TenantClient) Groups() tenantgroup.ServiceInterface {
	return tc.groups
}

// Region returns the region service interface
func (tc *TenantClient) Region() region.ServiceInterface {
	return tc.region
}
