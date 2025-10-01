package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/services/gatewayconfig"
)

// MockGatewayConfigService implements gatewayconfig.ServiceInterface for testing
type MockGatewayConfigService struct {
	ListGatewayConfigsFunc        func(ctx context.Context) ([]gatewayconfig.GatewayConfig, error)
	GetGatewayConfigByIDFunc      func(ctx context.Context, id string) (*gatewayconfig.GatewayConfig, error)
	CreateGatewayConfigFunc       func(ctx context.Context, gatewayConfig *gatewayconfig.GatewayConfig) (*gatewayconfig.GatewayConfig, error)
	UpdateGatewayConfigFunc       func(ctx context.Context, gatewayConfig *gatewayconfig.GatewayConfig) (*gatewayconfig.GatewayConfig, error)
	DeleteGatewayConfigFunc       func(ctx context.Context, id string) error
	GetGatewayServerConfigFunc    func(ctx context.Context, gatewayID string) (*gatewayconfig.GWServerConfig, error)
	UpdateGatewayServerConfigFunc func(ctx context.Context, gatewayID string, gatewayServerConfig *gatewayconfig.GWServerConfig) (*gatewayconfig.GWServerConfig, error)
}

func (m *MockGatewayConfigService) List(ctx context.Context) ([]gatewayconfig.GatewayConfig, error) {
	if m.ListGatewayConfigsFunc != nil {
		return m.ListGatewayConfigsFunc(ctx)
	}
	return []gatewayconfig.GatewayConfig{}, nil
}

func (m *MockGatewayConfigService) GetByID(ctx context.Context, id string) (*gatewayconfig.GatewayConfig, error) {
	if m.GetGatewayConfigByIDFunc != nil {
		return m.GetGatewayConfigByIDFunc(ctx, id)
	}
	return &gatewayconfig.GatewayConfig{ID: id}, nil
}

func (m *MockGatewayConfigService) Create(ctx context.Context, gatewayConfig *gatewayconfig.GatewayConfig) (*gatewayconfig.GatewayConfig, error) {
	if m.CreateGatewayConfigFunc != nil {
		return m.CreateGatewayConfigFunc(ctx, gatewayConfig)
	}
	gatewayConfig.ID = "mock-gateway-id"
	return gatewayConfig, nil
}

func (m *MockGatewayConfigService) Update(ctx context.Context, gatewayConfig *gatewayconfig.GatewayConfig) (*gatewayconfig.GatewayConfig, error) {
	if m.UpdateGatewayConfigFunc != nil {
		return m.UpdateGatewayConfigFunc(ctx, gatewayConfig)
	}
	return gatewayConfig, nil
}

func (m *MockGatewayConfigService) Delete(ctx context.Context, id string) error {
	if m.DeleteGatewayConfigFunc != nil {
		return m.DeleteGatewayConfigFunc(ctx, id)
	}
	return nil
}

func (m *MockGatewayConfigService) GetServerConfig(ctx context.Context, gatewayID string) (*gatewayconfig.GWServerConfig, error) {
	if m.GetGatewayServerConfigFunc != nil {
		return m.GetGatewayServerConfigFunc(ctx, gatewayID)
	}
	return &gatewayconfig.GWServerConfig{}, nil
}

func (m *MockGatewayConfigService) UpdateServerConfig(ctx context.Context, gatewayID string, gatewayServerConfig *gatewayconfig.GWServerConfig) (*gatewayconfig.GWServerConfig, error) {
	if m.UpdateGatewayServerConfigFunc != nil {
		return m.UpdateGatewayServerConfigFunc(ctx, gatewayID, gatewayServerConfig)
	}
	return gatewayServerConfig, nil
}

// Compile-time interface compliance check
var _ gatewayconfig.ServiceInterface = (*MockGatewayConfigService)(nil)
