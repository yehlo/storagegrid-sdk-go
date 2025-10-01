package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/services/hagroup"
)

// MockHAGroupService implements hagroup.ServiceInterface for testing
type MockHAGroupService struct {
	ListFunc    func(ctx context.Context) ([]hagroup.HAGroup, error)
	GetByIDFunc func(ctx context.Context, id string) (*hagroup.HAGroup, error)
	CreateFunc  func(ctx context.Context, hagroup *hagroup.HAGroup) (*hagroup.HAGroup, error)
	UpdateFunc  func(ctx context.Context, hagroup *hagroup.HAGroup) (*hagroup.HAGroup, error)
	DeleteFunc  func(ctx context.Context, id string) error
}

func (m *MockHAGroupService) List(ctx context.Context) ([]hagroup.HAGroup, error) {
	if m.ListFunc != nil {
		return m.ListFunc(ctx)
	}
	return []hagroup.HAGroup{}, nil
}

func (m *MockHAGroupService) GetByID(ctx context.Context, id string) (*hagroup.HAGroup, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, id)
	}
	return &hagroup.HAGroup{ID: id}, nil
}

func (m *MockHAGroupService) Create(ctx context.Context, hagroup *hagroup.HAGroup) (*hagroup.HAGroup, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, hagroup)
	}
	hagroup.ID = "mock-hagroup-id"
	return hagroup, nil
}

func (m *MockHAGroupService) Update(ctx context.Context, hagroup *hagroup.HAGroup) (*hagroup.HAGroup, error) {
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
var _ hagroup.ServiceInterface = (*MockHAGroupService)(nil)
