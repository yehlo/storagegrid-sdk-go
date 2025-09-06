package services

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
)

const (
	tenantEndpoint string = "/grid/accounts"
)

// TenantServiceInterface defines the contract for tenant service operations
type TenantServiceInterface interface {
	List(ctx context.Context) (*[]models.Tenant, error)
	GetById(ctx context.Context, id string) (*models.Tenant, error)
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

func (s *TenantService) List(ctx context.Context) (*[]models.Tenant, error) {
	response := models.Response{}
	response.Data = &[]models.Tenant{}
	err := s.client.DoParsed(ctx, "GET", tenantEndpoint, nil, &response)
	if err != nil {
		return nil, err
	}

	tenants := response.Data.(*[]models.Tenant)

	return tenants, nil
}

func (s *TenantService) GetById(ctx context.Context, id string) (*models.Tenant, error) {
	response := models.Response{}
	response.Data = &models.Tenant{}
	err := s.client.DoParsed(ctx, "GET", tenantEndpoint+"/"+id, nil, &response)
	if err != nil {
		return nil, err
	}

	tenant := response.Data.(*models.Tenant)

	return tenant, nil
}

func (s *TenantService) Create(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error) {
	response := models.Response{}
	response.Data = &models.Tenant{}
	err := s.client.DoParsed(ctx, "POST", tenantEndpoint, tenant, &response)
	if err != nil {
		return nil, err
	}

	tenant = response.Data.(*models.Tenant)

	return tenant, nil
}

func (s *TenantService) Update(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error) {
	response := models.Response{}
	response.Data = &models.Tenant{}
	err := s.client.DoParsed(ctx, "PUT", tenantEndpoint+"/"+tenant.Id, tenant, &response)
	if err != nil {
		return nil, err
	}

	tenant = response.Data.(*models.Tenant)

	return tenant, nil
}

func (s *TenantService) Delete(ctx context.Context, id string) error {
	err := s.client.DoParsed(ctx, "DELETE", tenantEndpoint+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *TenantService) GetUsage(ctx context.Context, id string) (*models.TenantUsage, error) {
	response := models.Response{}
	response.Data = &models.TenantUsage{}

	err := s.client.DoParsed(ctx, "GET", tenantEndpoint+"/"+id+"/usage", nil, &response)
	if err != nil {
		return nil, err
	}

	usage := response.Data.(*models.TenantUsage)
	return usage, nil
}
