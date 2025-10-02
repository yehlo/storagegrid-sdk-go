package health

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
)

const (
	healthEndpoint string = "/grid/health"
)

// ServiceInterface defines the contract for health service operations
type ServiceInterface interface {
	Get(ctx context.Context) (*Health, error)
}

type Service struct {
	services.HTTPClient
}

// NewService returns a new health service using the provided client
func NewService(client services.HTTPClient) *Service {
	return &Service{client}
}

func (s *Service) Get(ctx context.Context) (*Health, error) {
	var response models.Response[*Health]
	if err := s.DoParsed(ctx, "GET", healthEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
