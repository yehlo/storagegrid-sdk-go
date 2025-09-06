package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/services"
)

// MockRegionService implements services.RegionServiceInterface for testing
type MockRegionService struct {
	ListFunc func(ctx context.Context) (*[]string, error)
}

func (m *MockRegionService) List(ctx context.Context) (*[]string, error) {
	if m.ListFunc != nil {
		return m.ListFunc(ctx)
	}
	return &[]string{"us-east-1", "us-west-2"}, nil
}

// Compile-time interface compliance check
var _ services.RegionServiceInterface = (*MockRegionService)(nil)
