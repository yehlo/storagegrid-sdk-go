package services

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
)

const (
	tenantEndpoint = "/grid/accounts"
)

// TenantServiceInterface defines the contract for tenant service operations
type TenantServiceInterface interface {
	List(ctx context.Context) ([]models.Tenant, error)
	GetByID(ctx context.Context, id string) (*models.Tenant, error)
	Create(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error)
	Update(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error)
	Delete(ctx context.Context, id string) error
	GetUsage(ctx context.Context, id string) (*models.TenantUsage, error)
}

type TenantService struct {
	client HTTPClient
}

func NewTenantService(client HTTPClient) *TenantService {
	return &TenantService{client: client}
}

func (s *TenantService) List(ctx context.Context) ([]models.Tenant, error) {
	var response models.Response[[]models.Tenant]
	err := s.client.DoParsed(ctx, "GET", tenantEndpoint, nil, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *TenantService) GetByID(ctx context.Context, id string) (*models.Tenant, error) {
	var response models.Response[*models.Tenant]
	err := s.client.DoParsed(ctx, "GET", tenantEndpoint+"/"+id, nil, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *TenantService) Create(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error) {
	var response models.Response[*models.Tenant]
	err := s.client.DoParsed(ctx, "POST", tenantEndpoint, tenant, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *TenantService) Update(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error) {
	var response models.Response[*models.Tenant]
	err := s.client.DoParsed(ctx, "PUT", tenantEndpoint+"/"+tenant.ID, tenant, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *TenantService) Delete(ctx context.Context, id string) error {
	err := s.client.DoParsed(ctx, "DELETE", tenantEndpoint+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *TenantService) GetUsage(ctx context.Context, id string) (*models.TenantUsage, error) {
	var response models.Response[*models.TenantUsage]

	err := s.client.DoParsed(ctx, "GET", tenantEndpoint+"/"+id+"/usage", nil, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
