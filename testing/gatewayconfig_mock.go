package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/services/gateway"
)

// MockGatewayConfigService implements gatewayconfig.ServiceInterface for testing
type MockGatewayConfigService struct {
	ListGatewayConfigsFunc        func(ctx context.Context) ([]gateway.Config, error)
	GetGatewayConfigByIDFunc      func(ctx context.Context, id string) (*gateway.Config, error)
	CreateGatewayConfigFunc       func(ctx context.Context, gatewayConfig *gateway.Config) (*gateway.Config, error)
	UpdateGatewayConfigFunc       func(ctx context.Context, gatewayConfig *gateway.Config) (*gateway.Config, error)
	DeleteGatewayConfigFunc       func(ctx context.Context, id string) error
	GetGatewayServerConfigFunc    func(ctx context.Context, gatewayID string) (*gateway.ServerConfig, error)
	UpdateGatewayServerConfigFunc func(ctx context.Context, gatewayID string, gatewayServerConfig *gateway.ServerConfig) (*gateway.ServerConfig, error)
}

func (m *MockGatewayConfigService) ListConfig(ctx context.Context) ([]gateway.Config, error) {
	if m.ListGatewayConfigsFunc != nil {
		return m.ListGatewayConfigsFunc(ctx)
	}
	return []gateway.Config{}, nil
}

func (m *MockGatewayConfigService) GetConfigByID(ctx context.Context, id string) (*gateway.Config, error) {
	if m.GetGatewayConfigByIDFunc != nil {
		return m.GetGatewayConfigByIDFunc(ctx, id)
	}
	return &gateway.Config{ID: id}, nil
}

func (m *MockGatewayConfigService) CreateConfig(ctx context.Context, gatewayConfig *gateway.Config) (*gateway.Config, error) {
	if m.CreateGatewayConfigFunc != nil {
		return m.CreateGatewayConfigFunc(ctx, gatewayConfig)
	}
	gatewayConfig.ID = "mock-gateway-id"
	return gatewayConfig, nil
}

func (m *MockGatewayConfigService) UpdateConfig(ctx context.Context, gatewayConfig *gateway.Config) (*gateway.Config, error) {
	if m.UpdateGatewayConfigFunc != nil {
		return m.UpdateGatewayConfigFunc(ctx, gatewayConfig)
	}
	return gatewayConfig, nil
}

func (m *MockGatewayConfigService) DeleteConfig(ctx context.Context, id string) error {
	if m.DeleteGatewayConfigFunc != nil {
		return m.DeleteGatewayConfigFunc(ctx, id)
	}
	return nil
}

func (m *MockGatewayConfigService) GetServerConfig(ctx context.Context, gatewayID string) (*gateway.ServerConfig, error) {
	if m.GetGatewayServerConfigFunc != nil {
		return m.GetGatewayServerConfigFunc(ctx, gatewayID)
	}
	return &gateway.ServerConfig{}, nil
}

func (m *MockGatewayConfigService) UpdateServerConfig(ctx context.Context, gatewayID string, gatewayServerConfig *gateway.ServerConfig) (*gateway.ServerConfig, error) {
	if m.UpdateGatewayServerConfigFunc != nil {
		return m.UpdateGatewayServerConfigFunc(ctx, gatewayID, gatewayServerConfig)
	}
	return gatewayServerConfig, nil
}

// Compile-time interface compliance check
var _ gateway.ServiceInterface = (*MockGatewayConfigService)(nil)
