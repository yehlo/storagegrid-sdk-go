package client

import (
	"net/url"

	"github.com/yehlo/storagegrid-sdk-go/services"
)

const (
	tenantAPI = "api/v4"
)

type TenantClient struct {
	client *Client

	// Services
	bucket       services.BucketServiceInterface
	s3AccessKeys services.S3AccessKeyServiceInterface
	users        services.TenantUserServiceInterface
	groups       services.TenantGroupServiceInterface
	region       services.RegionServiceInterface
}

func NewTenantClient(options ...ClientOption) (*TenantClient, error) {
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

func (tc *TenantClient) Bucket() services.BucketServiceInterface {
	return tc.bucket
}

func (tc *TenantClient) S3AccessKeys() services.S3AccessKeyServiceInterface {
	return tc.s3AccessKeys
}

func (tc *TenantClient) Users() services.TenantUserServiceInterface {
	return tc.users
}

func (tc *TenantClient) Groups() services.TenantGroupServiceInterface {
	return tc.groups
}

func (tc *TenantClient) Region() services.RegionServiceInterface {
	return tc.region
}
