package services

import (
	"context"
	"net/http"

	"github.com/yehlo/storagegrid-sdk-go/models"
)

const (
	trafficClassEndpoint = "/grid/traffic-classes/policies"
)

// TrafficClassServiceInterface defines the contract for tenant service operations
type TrafficClassServiceInterface interface {
	CreatePolicy(ctx context.Context, tenant *models.TrafficClassPolicy) (*models.TrafficClassPolicy, error)
	// Delete(ctx context.Context, id string) error
	// GetByID(ctx context.Context, id string) (*models.Tenant, error)
	List(context.Context) ([]models.TrafficClass, error)
	// Update(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error)
}

type TrafficClassService struct {
	client HTTPClient
}

func NewTrafficClassService(client HTTPClient) *TrafficClassService {
	return &TrafficClassService{client: client}
}

func (s *TrafficClassService) List(ctx context.Context) ([]models.TrafficClass, error) {
	var response models.Response[[]models.TrafficClass]
	if err := s.client.DoParsed(ctx, http.MethodGet, trafficClassEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *TrafficClassService) CreatePolicy(ctx context.Context, tenant *models.TrafficClassPolicy) (*models.TrafficClassPolicy, error) {
	var response models.Response[*models.TrafficClassPolicy]
	if err := s.client.DoParsed(ctx, http.MethodPost, trafficClassEndpoint, tenant, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// Compile-time interface compliance check
var _ TrafficClassServiceInterface = (*TrafficClassService)(nil)
