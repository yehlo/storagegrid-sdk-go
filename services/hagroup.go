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
	List(ctx context.Context) (*[]models.HAGroup, error)
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

func (s *HAGroupService) List(ctx context.Context) (*[]models.HAGroup, error) {
	response := models.Response{}
	response.Data = &[]models.HAGroup{}
	err := s.client.DoParsed(ctx, "GET", hagroupEndpoint, nil, &response)
	if err != nil {
		return nil, err
	}

	hagroups := response.Data.(*[]models.HAGroup)

	return hagroups, nil
}

func (s *HAGroupService) GetById(ctx context.Context, id string) (*models.HAGroup, error) {
	response := models.Response{}
	response.Data = &models.HAGroup{}
	err := s.client.DoParsed(ctx, "GET", hagroupEndpoint+"/"+id, nil, &response)
	if err != nil {
		return nil, err
	}

	hagroup := response.Data.(*models.HAGroup)

	return hagroup, nil
}

func (s *HAGroupService) Create(ctx context.Context, hagroup *models.HAGroup) (*models.HAGroup, error) {
	response := models.Response{}
	response.Data = &models.HAGroup{}
	err := s.client.DoParsed(ctx, "POST", hagroupEndpoint, hagroup, &response)
	if err != nil {
		return nil, err
	}

	hagroup = response.Data.(*models.HAGroup)

	return hagroup, nil
}

func (s *HAGroupService) Update(ctx context.Context, hagroup *models.HAGroup) (*models.HAGroup, error) {
	response := models.Response{}
	response.Data = &models.HAGroup{}
	err := s.client.DoParsed(ctx, "PUT", hagroupEndpoint+"/"+hagroup.Id, hagroup, &response)
	if err != nil {
		return nil, err
	}

	hagroup = response.Data.(*models.HAGroup)

	return hagroup, nil
}

func (s *HAGroupService) Delete(ctx context.Context, id string) error {
	err := s.client.DoParsed(ctx, "DELETE", hagroupEndpoint+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
