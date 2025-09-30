package services

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
)

const (
	hagroupEndpoint string = "/private/ha-groups"
)

// HAGroupServiceInterface defines the contract for HA group service operations
type HAGroupServiceInterface interface {
	List(ctx context.Context) ([]models.HAGroup, error)
	GetById(ctx context.Context, id string) (*models.HAGroup, error)
	Create(ctx context.Context, hagroup *models.HAGroup) (*models.HAGroup, error)
	Update(ctx context.Context, hagroup *models.HAGroup) (*models.HAGroup, error)
	Delete(ctx context.Context, id string) error
}

type HAGroupService struct {
	client HTTPClient
}

func NewHAGroupService(client HTTPClient) *HAGroupService {
	return &HAGroupService{client: client}
}

func (s *HAGroupService) List(ctx context.Context) ([]models.HAGroup, error) {
	var response models.Response[[]models.HAGroup]
	if err := s.client.DoParsed(ctx, "GET", hagroupEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *HAGroupService) GetById(ctx context.Context, id string) (*models.HAGroup, error) {
	var response models.Response[*models.HAGroup]
	if err := s.client.DoParsed(ctx, "GET", hagroupEndpoint+"/"+id, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *HAGroupService) Create(ctx context.Context, hagroup *models.HAGroup) (*models.HAGroup, error) {
	var response models.Response[*models.HAGroup]
	if err := s.client.DoParsed(ctx, "POST", hagroupEndpoint, hagroup, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *HAGroupService) Update(ctx context.Context, hagroup *models.HAGroup) (*models.HAGroup, error) {
	var response models.Response[*models.HAGroup]
	if err := s.client.DoParsed(ctx, "PUT", hagroupEndpoint+"/"+hagroup.Id, hagroup, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *HAGroupService) Delete(ctx context.Context, id string) error {
	return s.client.DoParsed(ctx, "DELETE", hagroupEndpoint+"/"+id, nil, nil)
}
