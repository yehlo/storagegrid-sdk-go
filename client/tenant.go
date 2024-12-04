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
	Bucket       *services.BucketService
	S3AccessKeys *services.S3AccessKeyService
	Users        *services.TenantUserService
	Groups       *services.TenantGroupService
}

func NewTenantClient(options ...ClientOption) (*TenantClient, error) {
	c, err := newClient(options...)
	if err != nil {
		return nil, err
	}

	c.baseURL = c.baseURL.ResolveReference(&url.URL{Path: tenantAPI})

	return &TenantClient{
		client:       c,
		Bucket:       services.NewBucketService(c),
		S3AccessKeys: services.NewS3AccessKeyService(c),
		Users:        services.NewTenantUserService(c),
		Groups:       services.NewTenantGroupService(c),
	}, nil
}
