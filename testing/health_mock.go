package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/services/health"
)

// MockHealthService implements health.ServiceInterface for testing
type MockHealthService struct {
	GetFunc func(ctx context.Context) (*health.Health, error)
}

func (m *MockHealthService) Get(ctx context.Context) (*health.Health, error) {
	if m.GetFunc != nil {
		return m.GetFunc(ctx)
	}
	return &health.Health{}, nil
}

// Compile-time interface compliance check
var _ health.ServiceInterface = (*MockHealthService)(nil)
