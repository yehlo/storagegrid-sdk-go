package client

import (
	"net/url"

	"github.com/yehlo/storagegrid-sdk-go/services"
)

const (
	tenantAPI = "api/v4"
)

// TenantClient is the base struct used to interface with the storagegrid tenant API
type TenantClient struct {
	client *Client

	// Services

	bucket       services.BucketServiceInterface
	s3AccessKeys services.S3AccessKeyServiceInterface
	users        services.TenantUserServiceInterface
	groups       services.TenantGroupServiceInterface
	region       services.RegionServiceInterface
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
		bucket:       services.NewBucketService(c),
		s3AccessKeys: services.NewS3AccessKeyService(c),
		users:        services.NewTenantUserService(c),
		groups:       services.NewTenantGroupService(c),
		region:       services.NewRegionTenantService(c),
	}, nil
}

// Service getters return interfaces to enable testing with mocks

// Bucket returns the bucket service interface
func (tc *TenantClient) Bucket() services.BucketServiceInterface {
	return tc.bucket
}

// S3AccessKeys returns the S3 access key service interface
func (tc *TenantClient) S3AccessKeys() services.S3AccessKeyServiceInterface {
	return tc.s3AccessKeys
}

// Users returns the tenant user service interface
func (tc *TenantClient) Users() services.TenantUserServiceInterface {
	return tc.users
}

// Groups returns the tenant group service interface
func (tc *TenantClient) Groups() services.TenantGroupServiceInterface {
	return tc.groups
}

// Region returns the region service interface
func (tc *TenantClient) Region() services.RegionServiceInterface {
	return tc.region
}
