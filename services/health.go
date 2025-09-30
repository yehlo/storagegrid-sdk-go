package services

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
)

const (
	healthEndpoint string = "/grid/health"
)

// HealthServiceInterface defines the contract for health service operations
type HealthServiceInterface interface {
	Get(ctx context.Context) (*models.Health, error)
}

type HealthService struct {
	client HTTPClient
}

func NewHealthService(client HTTPClient) *HealthService {
	return &HealthService{client: client}
}

func (s *HealthService) Get(ctx context.Context) (*models.Health, error) {
	var response models.Response[*models.Health]
	if err := s.client.DoParsed(ctx, "GET", healthEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
