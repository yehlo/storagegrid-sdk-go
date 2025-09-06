package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
)

// MockHAGroupService implements services.HAGroupServiceInterface for testing
type MockHAGroupService struct {
	ListFunc    func(ctx context.Context) (*[]models.HAGroup, error)
	GetByIdFunc func(ctx context.Context, id string) (*models.HAGroup, error)
	CreateFunc  func(ctx context.Context, hagroup *models.HAGroup) (*models.HAGroup, error)
	UpdateFunc  func(ctx context.Context, hagroup *models.HAGroup) (*models.HAGroup, error)
	DeleteFunc  func(ctx context.Context, id string) error
}

func (m *MockHAGroupService) List(ctx context.Context) (*[]models.HAGroup, error) {
	if m.ListFunc != nil {
		return m.ListFunc(ctx)
	}
	return &[]models.HAGroup{}, nil
}

func (m *MockHAGroupService) GetById(ctx context.Context, id string) (*models.HAGroup, error) {
	if m.GetByIdFunc != nil {
		return m.GetByIdFunc(ctx, id)
	}
	return &models.HAGroup{Id: id}, nil
}

func (m *MockHAGroupService) Create(ctx context.Context, hagroup *models.HAGroup) (*models.HAGroup, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, hagroup)
	}
	hagroup.Id = "mock-hagroup-id"
	return hagroup, nil
}

func (m *MockHAGroupService) Update(ctx context.Context, hagroup *models.HAGroup) (*models.HAGroup, error) {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(ctx, hagroup)
	}
	return hagroup, nil
}

func (m *MockHAGroupService) Delete(ctx context.Context, id string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(ctx, id)
	}
	return nil
}

// Compile-time interface compliance check
var _ services.HAGroupServiceInterface = (*MockHAGroupService)(nil)
