package hagroup

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
)

const (
	endpoint string = "/private/ha-groups"
)

// ServiceInterface defines the contract for HA group service operations
type ServiceInterface interface {
	List(ctx context.Context) ([]HAGroup, error)
	GetByID(ctx context.Context, id string) (*HAGroup, error)
	Create(ctx context.Context, hagroup *HAGroup) (*HAGroup, error)
	Update(ctx context.Context, hagroup *HAGroup) (*HAGroup, error)
	Delete(ctx context.Context, id string) error
}

type Service struct {
	client services.HTTPClient
}

func NewService(client services.HTTPClient) *Service {
	return &Service{client: client}
}

func (s *Service) List(ctx context.Context) ([]HAGroup, error) {
	var response models.Response[[]HAGroup]
	if err := s.client.DoParsed(ctx, "GET", endpoint, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*HAGroup, error) {
	var response models.Response[*HAGroup]
	if err := s.client.DoParsed(ctx, "GET", endpoint+"/"+id, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) Create(ctx context.Context, hagroup *HAGroup) (*HAGroup, error) {
	var response models.Response[*HAGroup]
	if err := s.client.DoParsed(ctx, "POST", endpoint, hagroup, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) Update(ctx context.Context, hagroup *HAGroup) (*HAGroup, error) {
	var response models.Response[*HAGroup]
	if err := s.client.DoParsed(ctx, "PUT", endpoint+"/"+hagroup.ID, hagroup, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.client.DoParsed(ctx, "DELETE", endpoint+"/"+id, nil, nil)
}
