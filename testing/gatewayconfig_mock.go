package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
)

// MockGatewayConfigService implements services.GatewayConfigServiceInterface for testing
type MockGatewayConfigService struct {
	ListGatewayConfigsFunc        func(ctx context.Context) (*[]models.GatewayConfig, error)
	GetGatewayConfigByIdFunc      func(ctx context.Context, id string) (*models.GatewayConfig, error)
	CreateGatewayConfigFunc       func(ctx context.Context, gatewayConfig *models.GatewayConfig) (*models.GatewayConfig, error)
	UpdateGatewayConfigFunc       func(ctx context.Context, gatewayConfig *models.GatewayConfig) (*models.GatewayConfig, error)
	DeleteGatewayConfigFunc       func(ctx context.Context, id string) error
	GetGatewayServerConfigFunc    func(ctx context.Context, gatewayID string) (*models.GWServerConfig, error)
	UpdateGatewayServerConfigFunc func(ctx context.Context, gatewayID string, gatewayServerConfig *models.GWServerConfig) (*models.GWServerConfig, error)
}

func (m *MockGatewayConfigService) ListGatewayConfigs(ctx context.Context) (*[]models.GatewayConfig, error) {
	if m.ListGatewayConfigsFunc != nil {
		return m.ListGatewayConfigsFunc(ctx)
	}
	return &[]models.GatewayConfig{}, nil
}

func (m *MockGatewayConfigService) GetGatewayConfigById(ctx context.Context, id string) (*models.GatewayConfig, error) {
	if m.GetGatewayConfigByIdFunc != nil {
		return m.GetGatewayConfigByIdFunc(ctx, id)
	}
	return &models.GatewayConfig{Id: id}, nil
}

func (m *MockGatewayConfigService) CreateGatewayConfig(ctx context.Context, gatewayConfig *models.GatewayConfig) (*models.GatewayConfig, error) {
	if m.CreateGatewayConfigFunc != nil {
		return m.CreateGatewayConfigFunc(ctx, gatewayConfig)
	}
	gatewayConfig.Id = "mock-gateway-id"
	return gatewayConfig, nil
}

func (m *MockGatewayConfigService) UpdateGatewayConfig(ctx context.Context, gatewayConfig *models.GatewayConfig) (*models.GatewayConfig, error) {
	if m.UpdateGatewayConfigFunc != nil {
		return m.UpdateGatewayConfigFunc(ctx, gatewayConfig)
	}
	return gatewayConfig, nil
}

func (m *MockGatewayConfigService) DeleteGatewayConfig(ctx context.Context, id string) error {
	if m.DeleteGatewayConfigFunc != nil {
		return m.DeleteGatewayConfigFunc(ctx, id)
	}
	return nil
}

func (m *MockGatewayConfigService) GetGatewayServerConfig(ctx context.Context, gatewayID string) (*models.GWServerConfig, error) {
	if m.GetGatewayServerConfigFunc != nil {
		return m.GetGatewayServerConfigFunc(ctx, gatewayID)
	}
	return &models.GWServerConfig{}, nil
}

func (m *MockGatewayConfigService) UpdateGatewayServerConfig(ctx context.Context, gatewayID string, gatewayServerConfig *models.GWServerConfig) (*models.GWServerConfig, error) {
	if m.UpdateGatewayServerConfigFunc != nil {
		return m.UpdateGatewayServerConfigFunc(ctx, gatewayID, gatewayServerConfig)
	}
	return gatewayServerConfig, nil
}

// Compile-time interface compliance check
var _ services.GatewayConfigServiceInterface = (*MockGatewayConfigService)(nil)
