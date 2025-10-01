package trafficclasses

import (
	"context"
	"net/http"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
)

const (
	endpoint = "/grid/traffic-classes/policies"
)

// ServiceInterface defines the contract for traffic class service operations
type ServiceInterface interface {
	CreatePolicy(ctx context.Context, tenant *Policy) (*Policy, error)
	// Delete(ctx context.Context, id string) error
	// GetByID(ctx context.Context, id string) (*tenant.Tenant, error)
	List(context.Context) ([]TrafficClass, error)
	// Update(ctx context.Context, tenant *tenant.Tenant) (*tenant.Tenant, error)
}

type Service struct {
	client services.HTTPClient
}

func NewService(client services.HTTPClient) *Service {
	return &Service{client: client}
}

func (s *Service) List(ctx context.Context) ([]TrafficClass, error) {
	var response models.Response[[]TrafficClass]
	if err := s.client.DoParsed(ctx, http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) CreatePolicy(ctx context.Context, tenant *Policy) (*Policy, error) {
	var response models.Response[*Policy]
	if err := s.client.DoParsed(ctx, http.MethodPost, endpoint, tenant, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// Compile-time interface compliance check
var _ ServiceInterface = (*Service)(nil)
