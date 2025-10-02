package region

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
)

const (
	gridRegionEndpoint   string = "/grid/regions"
	tenantRegionEndpoint string = "/org/regions"
)

// ServiceInterface defines the contract for region service operations
type ServiceInterface interface {
	List(ctx context.Context) ([]string, error)
}

type Service struct {
	services.HTTPClient
	endpoint string
}

func NewGridService(client services.HTTPClient) *Service {
	return &Service{HTTPClient: client, endpoint: gridRegionEndpoint}
}

func NewTenantService(client services.HTTPClient) *Service {
	return &Service{HTTPClient: client, endpoint: tenantRegionEndpoint}
}

func (s *Service) List(ctx context.Context) ([]string, error) {
	var response models.Response[[]string]
	if err := s.DoParsed(ctx, "GET", s.endpoint, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
