package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
)

// MockHealthService implements services.HealthServiceInterface for testing
type MockHealthService struct {
	GetFunc func(ctx context.Context) (*models.Health, error)
}

func (m *MockHealthService) Get(ctx context.Context) (*models.Health, error) {
	if m.GetFunc != nil {
		return m.GetFunc(ctx)
	}
	return &models.Health{}, nil
}

// Compile-time interface compliance check
var _ services.HealthServiceInterface = (*MockHealthService)(nil)
