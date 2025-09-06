package services

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
)

const (
	gridRegionEndpoint   string = "/grid/regions"
	tenantRegionEndpoint string = "/org/regions"
)

// RegionServiceInterface defines the contract for region service operations
type RegionServiceInterface interface {
	List(ctx context.Context) (*[]string, error)
}

type RegionService struct {
	client   HTTPClient
	endpoint string
}

func NewRegionGridService(client HTTPClient) *RegionService {
	return &RegionService{client: client, endpoint: gridRegionEndpoint}
}

func NewRegionTenantService(client HTTPClient) *RegionService {
	return &RegionService{client: client, endpoint: tenantRegionEndpoint}
}

func (s *RegionService) List(ctx context.Context) (*[]string, error) {
	response := models.Response{}
	response.Data = &[]string{}
	err := s.client.DoParsed(ctx, "GET", s.endpoint, nil, &response)
	if err != nil {
		return nil, err
	}

	regions := response.Data.(*[]string)

	return regions, nil
}
