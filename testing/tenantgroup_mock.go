package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
)

// MockTenantGroupService implements services.TenantGroupServiceInterface for testing
type MockTenantGroupService struct {
	ListFunc      func(ctx context.Context) ([]models.TenantGroup, error)
	GetByIDFunc   func(ctx context.Context, id string) (*models.TenantGroup, error)
	GetByNameFunc func(ctx context.Context, name string) (*models.TenantGroup, error)
	CreateFunc    func(ctx context.Context, group *models.TenantGroup) (*models.TenantGroup, error)
	UpdateFunc    func(ctx context.Context, group *models.TenantGroup) (*models.TenantGroup, error)
	DeleteFunc    func(ctx context.Context, id string) error
}

func (m *MockTenantGroupService) List(ctx context.Context) ([]models.TenantGroup, error) {
	if m.ListFunc != nil {
		return m.ListFunc(ctx)
	}
	return []models.TenantGroup{}, nil
}

func (m *MockTenantGroupService) GetByID(ctx context.Context, id string) (*models.TenantGroup, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, id)
	}
	mockID := id
	return &models.TenantGroup{ID: &mockID}, nil
}

func (m *MockTenantGroupService) GetByName(ctx context.Context, name string) (*models.TenantGroup, error) {
	if m.GetByNameFunc != nil {
		return m.GetByNameFunc(ctx, name)
	}
	mockID := "mock-group-id"
	return &models.TenantGroup{ID: &mockID, UniqueName: name}, nil
}

func (m *MockTenantGroupService) Create(ctx context.Context, group *models.TenantGroup) (*models.TenantGroup, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, group)
	}
	mockID := "mock-group-id"
	group.ID = &mockID
	return group, nil
}

func (m *MockTenantGroupService) Update(ctx context.Context, group *models.TenantGroup) (*models.TenantGroup, error) {
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
var _ services.TenantGroupServiceInterface = (*MockTenantGroupService)(nil)
