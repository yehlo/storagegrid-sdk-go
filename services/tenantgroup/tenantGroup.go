package tenantgroup

import (
	"context"
	"strings"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
)

const (
	tenantGroupEndpoint string = "/org/groups"
)

// ServiceInterface defines the contract for tenant group service operations
type ServiceInterface interface {
	List(ctx context.Context) ([]TenantGroup, error)
	GetByID(ctx context.Context, id string) (*TenantGroup, error)
	GetByName(ctx context.Context, name string) (*TenantGroup, error)
	Create(ctx context.Context, group *TenantGroup) (*TenantGroup, error)
	Update(ctx context.Context, group *TenantGroup) (*TenantGroup, error)
	Delete(ctx context.Context, id string) error
}

type Service struct {
	client services.HTTPClient
}

func NewTenantGroupService(client services.HTTPClient) *Service {
	return &Service{client: client}
}

func (s *Service) List(ctx context.Context) ([]TenantGroup, error) {
	var response models.Response[[]TenantGroup]
	if err := s.client.DoParsed(ctx, "GET", tenantGroupEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*TenantGroup, error) {
	var response models.Response[*TenantGroup]
	if err := s.client.DoParsed(ctx, "GET", tenantGroupEndpoint+"/"+id, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) GetByName(ctx context.Context, name string) (*TenantGroup, error) {
	var response models.Response[*TenantGroup]
	if err := s.client.DoParsed(ctx, "GET", tenantGroupEndpoint+"/group/"+name, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) Create(ctx context.Context, group *TenantGroup) (*TenantGroup, error) {
	// enforce group/ prefix on group.uniqueName if manually created
	if !strings.HasPrefix(group.UniqueName, "group/") {
		group.UniqueName = "group/" + group.UniqueName
	}

	var response models.Response[*TenantGroup]
	if err := s.client.DoParsed(ctx, "POST", tenantGroupEndpoint, group, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) Update(ctx context.Context, group *TenantGroup) (*TenantGroup, error) {
	var response models.Response[*TenantGroup]
	if err := s.client.DoParsed(ctx, "PUT", tenantGroupEndpoint+"/"+*group.ID, group, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.client.DoParsed(ctx, "DELETE", tenantGroupEndpoint+"/"+id, nil, nil)
}
