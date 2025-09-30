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
	List(ctx context.Context) ([]models.TenantGroup, error)
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

func (s *TenantGroupService) List(ctx context.Context) ([]models.TenantGroup, error) {
	var response models.Response[[]models.TenantGroup]
	if err := s.client.DoParsed(ctx, "GET", tenantGroupEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *TenantGroupService) GetById(ctx context.Context, id string) (*models.TenantGroup, error) {
	var response models.Response[*models.TenantGroup]
	if err := s.client.DoParsed(ctx, "GET", tenantGroupEndpoint+"/"+id, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *TenantGroupService) GetByName(ctx context.Context, name string) (*models.TenantGroup, error) {
	var response models.Response[*models.TenantGroup]
	if err := s.client.DoParsed(ctx, "GET", tenantGroupEndpoint+"/group/"+name, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *TenantGroupService) Create(ctx context.Context, group *models.TenantGroup) (*models.TenantGroup, error) {
	// enforce group/ prefix on group.uniqueName if manually created
	if !strings.HasPrefix(group.UniqueName, "group/") {
		group.UniqueName = "group/" + group.UniqueName
	}

	var response models.Response[*models.TenantGroup]
	if err := s.client.DoParsed(ctx, "POST", tenantGroupEndpoint, group, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *TenantGroupService) Update(ctx context.Context, group *models.TenantGroup) (*models.TenantGroup, error) {
	var response models.Response[*models.TenantGroup]
	if err := s.client.DoParsed(ctx, "PUT", tenantGroupEndpoint+"/"+*group.Id, group, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *TenantGroupService) Delete(ctx context.Context, id string) error {
	return s.client.DoParsed(ctx, "DELETE", tenantGroupEndpoint+"/"+id, nil, nil)
}
