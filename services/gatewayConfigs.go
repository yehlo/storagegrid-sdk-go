package services

import (
	"context"
	"fmt"

	"github.com/yehlo/storagegrid-sdk-go/models"
)

const (
	gatewayConfigEndpoint string = "/private/gateway-configs"
	serverConfigEndpoint  string = gatewayConfigEndpoint + "/%s/server-config"
)

// GatewayConfigServiceInterface defines the contract for gateway config service operations
type GatewayConfigServiceInterface interface {
	ListGatewayConfigs(ctx context.Context) ([]models.GatewayConfig, error)
	GetGatewayConfigByID(ctx context.Context, id string) (*models.GatewayConfig, error)
	CreateGatewayConfig(ctx context.Context, gatewayConfig *models.GatewayConfig) (*models.GatewayConfig, error)
	UpdateGatewayConfig(ctx context.Context, gatewayConfig *models.GatewayConfig) (*models.GatewayConfig, error)
	DeleteGatewayConfig(ctx context.Context, id string) error
	GetGatewayServerConfig(ctx context.Context, gatewayID string) (*models.GWServerConfig, error)
	UpdateGatewayServerConfig(ctx context.Context, gatewayID string, gatewayServerConfig *models.GWServerConfig) (*models.GWServerConfig, error)
}

func getServerConfigEndpoint(id string) string {
	return fmt.Sprintf(serverConfigEndpoint, id)
}

type GatewayConfigService struct {
	client HTTPClient
}

func NewGatewayConfigService(client HTTPClient) *GatewayConfigService {
	return &GatewayConfigService{client: client}
}

func (s *GatewayConfigService) ListGatewayConfigs(ctx context.Context) ([]models.GatewayConfig, error) {
	var response models.Response[[]models.GatewayConfig]
	if err := s.client.DoParsed(ctx, "GET", gatewayConfigEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *GatewayConfigService) GetGatewayConfigByID(ctx context.Context, id string) (*models.GatewayConfig, error) {
	var response models.Response[*models.GatewayConfig]
	if err := s.client.DoParsed(ctx, "GET", gatewayConfigEndpoint+"/"+id, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *GatewayConfigService) CreateGatewayConfig(ctx context.Context, gatewayConfig *models.GatewayConfig) (*models.GatewayConfig, error) {
	var response models.Response[*models.GatewayConfig]
	if err := s.client.DoParsed(ctx, "POST", gatewayConfigEndpoint, gatewayConfig, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *GatewayConfigService) UpdateGatewayConfig(ctx context.Context, gatewayConfig *models.GatewayConfig) (*models.GatewayConfig, error) {
	var response models.Response[*models.GatewayConfig]
	if err := s.client.DoParsed(ctx, "PUT", gatewayConfigEndpoint+"/"+gatewayConfig.ID, gatewayConfig, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *GatewayConfigService) DeleteGatewayConfig(ctx context.Context, id string) error {
	return s.client.DoParsed(ctx, "DELETE", gatewayConfigEndpoint+"/"+id, nil, nil)
}

func (s *GatewayConfigService) GetGatewayServerConfig(ctx context.Context, gatewayID string) (*models.GWServerConfig, error) {
	var response models.Response[*models.GWServerConfig]
	if err := s.client.DoParsed(ctx, "GET", getServerConfigEndpoint(gatewayID), nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *GatewayConfigService) UpdateGatewayServerConfig(ctx context.Context, gatewayID string, gatewayServerConfig *models.GWServerConfig) (*models.GWServerConfig, error) {
	var response models.Response[*models.GWServerConfig]
	if err := s.client.DoParsed(ctx, "PUT", getServerConfigEndpoint(gatewayID), gatewayServerConfig, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
