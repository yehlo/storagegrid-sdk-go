package gatewayconfig

import (
	"context"
	"fmt"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
)

const (
	gatewayConfigEndpoint string = "/private/gateway-configs"
	serverConfigEndpoint  string = gatewayConfigEndpoint + "/%s/server-config"
)

// ServiceInterface defines the contract for gateway config service operations
type ServiceInterface interface {
	List(ctx context.Context) ([]GatewayConfig, error)
	GetByID(ctx context.Context, id string) (*GatewayConfig, error)
	Create(ctx context.Context, gatewayConfig *GatewayConfig) (*GatewayConfig, error)
	Update(ctx context.Context, gatewayConfig *GatewayConfig) (*GatewayConfig, error)
	Delete(ctx context.Context, id string) error
	GetServerConfig(ctx context.Context, gatewayID string) (*GWServerConfig, error)
	UpdateServerConfig(ctx context.Context, gatewayID string, gatewayServerConfig *GWServerConfig) (*GWServerConfig, error)
}

func getServerConfigEndpoint(id string) string {
	return fmt.Sprintf(serverConfigEndpoint, id)
}

type Service struct {
	client services.HTTPClient
}

func NewService(client services.HTTPClient) *Service {
	return &Service{client: client}
}

func (s *Service) List(ctx context.Context) ([]GatewayConfig, error) {
	var response models.Response[[]GatewayConfig]
	if err := s.client.DoParsed(ctx, "GET", gatewayConfigEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*GatewayConfig, error) {
	var response models.Response[*GatewayConfig]
	if err := s.client.DoParsed(ctx, "GET", gatewayConfigEndpoint+"/"+id, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) Create(ctx context.Context, gatewayConfig *GatewayConfig) (*GatewayConfig, error) {
	var response models.Response[*GatewayConfig]
	if err := s.client.DoParsed(ctx, "POST", gatewayConfigEndpoint, gatewayConfig, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) Update(ctx context.Context, gatewayConfig *GatewayConfig) (*GatewayConfig, error) {
	var response models.Response[*GatewayConfig]
	if err := s.client.DoParsed(ctx, "PUT", gatewayConfigEndpoint+"/"+gatewayConfig.ID, gatewayConfig, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.client.DoParsed(ctx, "DELETE", gatewayConfigEndpoint+"/"+id, nil, nil)
}

func (s *Service) GetServerConfig(ctx context.Context, gatewayID string) (*GWServerConfig, error) {
	var response models.Response[*GWServerConfig]
	if err := s.client.DoParsed(ctx, "GET", getServerConfigEndpoint(gatewayID), nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) UpdateServerConfig(ctx context.Context, gatewayID string, gatewayServerConfig *GWServerConfig) (*GWServerConfig, error) {
	var response models.Response[*GWServerConfig]
	if err := s.client.DoParsed(ctx, "PUT", getServerConfigEndpoint(gatewayID), gatewayServerConfig, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
