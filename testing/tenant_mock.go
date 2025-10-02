package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/services/tenant"
	"github.com/yehlo/storagegrid-sdk-go/services/tenantusage"
)

// MockTenantService implements tenant.ServiceInterface for testing
type MockTenantService struct {
	ListFunc     func(ctx context.Context) ([]tenant.Tenant, error)
	GetByIDFunc  func(ctx context.Context, id string) (*tenant.Tenant, error)
	CreateFunc   func(ctx context.Context, tenant *tenant.Tenant) (*tenant.Tenant, error)
	UpdateFunc   func(ctx context.Context, tenant *tenant.Tenant) (*tenant.Tenant, error)
	DeleteFunc   func(ctx context.Context, id string) error
	GetUsageFunc func(ctx context.Context, id string) (*tenantusage.TenantUsage, error)
}

func (m *MockTenantService) List(ctx context.Context) ([]tenant.Tenant, error) {
	if m.ListFunc != nil {
		return m.ListFunc(ctx)
	}
	return []tenant.Tenant{}, nil
}

func (m *MockTenantService) GetByID(ctx context.Context, id string) (*tenant.Tenant, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, id)
	}
	return &tenant.Tenant{ID: id}, nil
}

func (m *MockTenantService) Create(ctx context.Context, tenant *tenant.Tenant) (*tenant.Tenant, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, tenant)
	}
	tenant.ID = "mock-tenant-id"
	return tenant, nil
}

func (m *MockTenantService) Update(ctx context.Context, tenant *tenant.Tenant) (*tenant.Tenant, error) {
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

func (m *MockTenantService) GetUsage(ctx context.Context, id string) (*tenantusage.TenantUsage, error) {
	if m.GetUsageFunc != nil {
		return m.GetUsageFunc(ctx, id)
	}
	return &tenantusage.TenantUsage{}, nil
}

// Compile-time interface compliance check
var _ tenant.ServiceInterface = (*MockTenantService)(nil)
