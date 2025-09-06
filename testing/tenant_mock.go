package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
)

// MockTenantService implements services.TenantServiceInterface for testing
type MockTenantService struct {
	ListFunc     func(ctx context.Context) (*[]models.Tenant, error)
	GetByIdFunc  func(ctx context.Context, id string) (*models.Tenant, error)
	CreateFunc   func(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error)
	UpdateFunc   func(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error)
	DeleteFunc   func(ctx context.Context, id string) error
	GetUsageFunc func(ctx context.Context, id string) (*models.TenantUsage, error)
}

func (m *MockTenantService) List(ctx context.Context) (*[]models.Tenant, error) {
	if m.ListFunc != nil {
		return m.ListFunc(ctx)
	}
	return &[]models.Tenant{}, nil
}

func (m *MockTenantService) GetById(ctx context.Context, id string) (*models.Tenant, error) {
	if m.GetByIdFunc != nil {
		return m.GetByIdFunc(ctx, id)
	}
	return &models.Tenant{Id: id}, nil
}

func (m *MockTenantService) Create(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, tenant)
	}
	tenant.Id = "mock-tenant-id"
	return tenant, nil
}

func (m *MockTenantService) Update(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error) {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(ctx, tenant)
	}
	return tenant, nil
}

func (m *MockTenantService) Delete(ctx context.Context, id string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(ctx, id)
	}
	return nil
}

func (m *MockTenantService) GetUsage(ctx context.Context, id string) (*models.TenantUsage, error) {
	if m.GetUsageFunc != nil {
		return m.GetUsageFunc(ctx, id)
	}
	return &models.TenantUsage{}, nil
}

// Compile-time interface compliance check
var _ services.TenantServiceInterface = (*MockTenantService)(nil)
