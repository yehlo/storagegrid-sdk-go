package gateway

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
	ListConfig(ctx context.Context) ([]Config, error)
	GetConfigByID(ctx context.Context, id string) (*Config, error)
	CreateConfig(ctx context.Context, gatewayConfig *Config) (*Config, error)
	UpdateConfig(ctx context.Context, gatewayConfig *Config) (*Config, error)
	DeleteConfig(ctx context.Context, id string) error
	GetServerConfig(ctx context.Context, gatewayID string) (*ServerConfig, error)
	UpdateServerConfig(ctx context.Context, gatewayID string, gatewayServerConfig *ServerConfig) (*ServerConfig, error)
}

func getServerConfigEndpoint(id string) string {
	return fmt.Sprintf(serverConfigEndpoint, id)
}

type Service struct {
	services.HTTPClient
}

// NewService returns a new gateway service using the provided client
func NewService(client services.HTTPClient) *Service {
	return &Service{client}
}

func (s *Service) ListConfig(ctx context.Context) ([]Config, error) {
	var response models.Response[[]Config]
	if err := s.DoParsed(ctx, "GET", gatewayConfigEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) GetConfigByID(ctx context.Context, id string) (*Config, error) {
	var response models.Response[*Config]
	if err := s.DoParsed(ctx, "GET", gatewayConfigEndpoint+"/"+id, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) CreateConfig(ctx context.Context, gatewayConfig *Config) (*Config, error) {
	var response models.Response[*Config]
	if err := s.DoParsed(ctx, "POST", gatewayConfigEndpoint, gatewayConfig, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) UpdateConfig(ctx context.Context, gatewayConfig *Config) (*Config, error) {
	var response models.Response[*Config]
	if err := s.DoParsed(ctx, "PUT", gatewayConfigEndpoint+"/"+gatewayConfig.ID, gatewayConfig, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) DeleteConfig(ctx context.Context, id string) error {
	return s.DoParsed(ctx, "DELETE", gatewayConfigEndpoint+"/"+id, nil, nil)
}

func (s *Service) GetServerConfig(ctx context.Context, gatewayID string) (*ServerConfig, error) {
	var response models.Response[*ServerConfig]
	if err := s.DoParsed(ctx, "GET", getServerConfigEndpoint(gatewayID), nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) UpdateServerConfig(ctx context.Context, gatewayID string, gatewayServerConfig *ServerConfig) (*ServerConfig, error) {
	var response models.Response[*ServerConfig]
	if err := s.DoParsed(ctx, "PUT", getServerConfigEndpoint(gatewayID), gatewayServerConfig, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
