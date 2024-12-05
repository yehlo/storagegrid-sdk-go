package services

import (
	"context"

	models "github.com/yehlo/storagegrid-sdk-go/models"
)

const (
	healthEndpoint string = "/grid/health"
)

type HealthService struct {
	client HTTPClient
}

func NewHealthService(client HTTPClient) *HealthService {
	return &HealthService{client: client}
}

func (s *HealthService) Get(ctx context.Context) (*models.Health, error) {
	response := models.Response{}
	response.Data = &models.Health{}
	err := s.client.DoParsed(ctx, "GET", healthEndpoint, nil, &response)
	if err != nil {
		return nil, err
	}

	health := response.Data.(*models.Health)

	return health, nil
}
