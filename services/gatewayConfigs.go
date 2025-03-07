package services

import (
	"context"
	"fmt"

	models "github.com/yehlo/storagegrid-sdk-go/models"
)

const (
	gatewayConfigEndpoint string = "/private/gateway-configs"
	serverConfigEndpoint  string = gatewayConfigEndpoint + "/%s/server-config"
)

func getServerConfigEndpoint(id string) string {
	return fmt.Sprintf(serverConfigEndpoint, id)
}

type GatewayConfigService struct {
	client HTTPClient
}

func NewGatewayConfigService(client HTTPClient) *GatewayConfigService {
	return &GatewayConfigService{client: client}
}

func (s *GatewayConfigService) ListGatewayConfigs(ctx context.Context) (*[]models.GatewayConfig, error) {
	response := models.Response{}
	response.Data = &[]models.GatewayConfig{}
	err := s.client.DoParsed(ctx, "GET", gatewayConfigEndpoint, nil, &response)
	if err != nil {
		return nil, err
	}

	gatewayConfigs := response.Data.(*[]models.GatewayConfig)

	return gatewayConfigs, nil
}

func (s *GatewayConfigService) GetGatewayConfigById(ctx context.Context, id string) (*models.GatewayConfig, error) {
	response := models.Response{}
	response.Data = &models.GatewayConfig{}
	err := s.client.DoParsed(ctx, "GET", gatewayConfigEndpoint+"/"+id, nil, &response)
	if err != nil {
		return nil, err
	}

	gatewayConfig := response.Data.(*models.GatewayConfig)

	return gatewayConfig, nil
}

func (s *GatewayConfigService) CreateGatewayConfig(ctx context.Context, gatewayConfig *models.GatewayConfig) (*models.GatewayConfig, error) {
	response := models.Response{}
	response.Data = &models.GatewayConfig{}
	err := s.client.DoParsed(ctx, "POST", gatewayConfigEndpoint, gatewayConfig, &response)
	if err != nil {
		return nil, err
	}

	gatewayConfig = response.Data.(*models.GatewayConfig)

	return gatewayConfig, nil
}

func (s *GatewayConfigService) UpdateGatewayConfig(ctx context.Context, gatewayConfig *models.GatewayConfig) (*models.GatewayConfig, error) {
	response := models.Response{}
	response.Data = &models.GatewayConfig{}
	err := s.client.DoParsed(ctx, "PUT", gatewayConfigEndpoint+"/"+gatewayConfig.Id, gatewayConfig, &response)
	if err != nil {
		return nil, err
	}

	gatewayConfig = response.Data.(*models.GatewayConfig)

	return gatewayConfig, nil
}

func (s *GatewayConfigService) DeleteGatewayConfig(ctx context.Context, id string) error {
	err := s.client.DoParsed(ctx, "DELETE", gatewayConfigEndpoint+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *GatewayConfigService) GetGatewayServerConfig(ctx context.Context, id string) (*models.GWServerConfig, error) {
	response := models.Response{}
	response.Data = &models.GWServerConfig{}
	err := s.client.DoParsed(ctx, "GET", getServerConfigEndpoint(id), nil, &response)
	if err != nil {
		return nil, err
	}

	gatewayServerConfig := response.Data.(*models.GWServerConfig)

	return gatewayServerConfig, nil
}

func (s *GatewayConfigService) UpdateGatewayServerConfig(ctx context.Context, gatewayServerConfig *models.GWServerConfig) (*models.GWServerConfig, error) {
	response := models.Response{}
	response.Data = &models.GWServerConfig{}
	err := s.client.DoParsed(ctx, "PUT", gatewayConfigEndpoint, gatewayServerConfig, &response)
	if err != nil {
		return nil, err
	}

	gatewayServerConfig = response.Data.(*models.GWServerConfig)

	return gatewayServerConfig, nil
}
