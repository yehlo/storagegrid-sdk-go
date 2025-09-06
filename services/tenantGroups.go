package services

import (
	"context"
	"strings"

	"github.com/yehlo/storagegrid-sdk-go/models"
)

const (
	tenantGroupEndpoint string = "/org/groups"
)

// TenantGroupServiceInterface defines the contract for tenant group service operations
type TenantGroupServiceInterface interface {
	List(ctx context.Context) (*[]models.TenantGroup, error)
	GetById(ctx context.Context, id string) (*models.TenantGroup, error)
	GetByName(ctx context.Context, name string) (*models.TenantGroup, error)
	Create(ctx context.Context, group *models.TenantGroup) (*models.TenantGroup, error)
	Update(ctx context.Context, group *models.TenantGroup) (*models.TenantGroup, error)
	Delete(ctx context.Context, id string) error
}

type TenantGroupService struct {
	client HTTPClient
}

func NewTenantGroupService(client HTTPClient) *TenantGroupService {
	return &TenantGroupService{client: client}
}

func (s *TenantGroupService) List(ctx context.Context) (*[]models.TenantGroup, error) {
	response := models.Response{}
	response.Data = &[]models.TenantGroup{}
	err := s.client.DoParsed(ctx, "GET", tenantGroupEndpoint, nil, &response)
	if err != nil {
		return nil, err
	}

	groups := response.Data.(*[]models.TenantGroup)

	return groups, nil
}

func (s *TenantGroupService) GetById(ctx context.Context, id string) (*models.TenantGroup, error) {
	response := models.Response{}
	response.Data = &models.TenantGroup{}
	err := s.client.DoParsed(ctx, "GET", tenantGroupEndpoint+"/"+id, nil, &response)
	if err != nil {
		return nil, err
	}

	group := response.Data.(*models.TenantGroup)

	return group, nil
}

func (s *TenantGroupService) GetByName(ctx context.Context, name string) (*models.TenantGroup, error) {
	response := models.Response{}
	response.Data = &models.TenantGroup{}
	err := s.client.DoParsed(ctx, "GET", tenantGroupEndpoint+"/group/"+name, nil, &response)
	if err != nil {
		return nil, err
	}

	group := response.Data.(*models.TenantGroup)

	return group, nil
}

func (s *TenantGroupService) Create(ctx context.Context, group *models.TenantGroup) (*models.TenantGroup, error) {
	// enforce group/ prefix on group.uniqueName if manually created
	if !strings.HasPrefix(group.UniqueName, "group/") {
		group.UniqueName = "group/" + group.UniqueName
	}

	response := models.Response{}
	response.Data = &models.TenantGroup{}
	err := s.client.DoParsed(ctx, "POST", tenantGroupEndpoint, group, &response)
	if err != nil {
		return nil, err
	}

	group = response.Data.(*models.TenantGroup)

	return group, nil
}

func (s *TenantGroupService) Update(ctx context.Context, group *models.TenantGroup) (*models.TenantGroup, error) {
	response := models.Response{}
	response.Data = &models.TenantGroup{}
	err := s.client.DoParsed(ctx, "PUT", tenantGroupEndpoint+"/"+*group.Id, group, &response)
	if err != nil {
		return nil, err
	}

	group = response.Data.(*models.TenantGroup)

	return group, nil
}

func (s *TenantGroupService) Delete(ctx context.Context, id string) error {
	err := s.client.DoParsed(ctx, "DELETE", tenantGroupEndpoint+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
