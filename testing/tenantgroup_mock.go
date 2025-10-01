package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/services/tenantgroup"
)

// MockTenantGroupService implements tenantgroup.ServiceInterface for testing
type MockTenantGroupService struct {
	ListFunc      func(ctx context.Context) ([]tenantgroup.TenantGroup, error)
	GetByIDFunc   func(ctx context.Context, id string) (*tenantgroup.TenantGroup, error)
	GetByNameFunc func(ctx context.Context, name string) (*tenantgroup.TenantGroup, error)
	CreateFunc    func(ctx context.Context, group *tenantgroup.TenantGroup) (*tenantgroup.TenantGroup, error)
	UpdateFunc    func(ctx context.Context, group *tenantgroup.TenantGroup) (*tenantgroup.TenantGroup, error)
	DeleteFunc    func(ctx context.Context, id string) error
}

func (m *MockTenantGroupService) List(ctx context.Context) ([]tenantgroup.TenantGroup, error) {
	if m.ListFunc != nil {
		return m.ListFunc(ctx)
	}
	return []tenantgroup.TenantGroup{}, nil
}

func (m *MockTenantGroupService) GetByID(ctx context.Context, id string) (*tenantgroup.TenantGroup, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, id)
	}
	mockID := id
	return &tenantgroup.TenantGroup{ID: &mockID}, nil
}

func (m *MockTenantGroupService) GetByName(ctx context.Context, name string) (*tenantgroup.TenantGroup, error) {
	if m.GetByNameFunc != nil {
		return m.GetByNameFunc(ctx, name)
	}
	mockID := "mock-group-id"
	return &tenantgroup.TenantGroup{ID: &mockID, UniqueName: name}, nil
}

func (m *MockTenantGroupService) Create(ctx context.Context, group *tenantgroup.TenantGroup) (*tenantgroup.TenantGroup, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, group)
	}
	mockID := "mock-group-id"
	group.ID = &mockID
	return group, nil
}

func (m *MockTenantGroupService) Update(ctx context.Context, group *tenantgroup.TenantGroup) (*tenantgroup.TenantGroup, error) {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(ctx, group)
	}
	return group, nil
}

func (m *MockTenantGroupService) Delete(ctx context.Context, id string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(ctx, id)
	}
	return nil
}

// Compile-time interface compliance check
var _ tenantgroup.ServiceInterface = (*MockTenantGroupService)(nil)
