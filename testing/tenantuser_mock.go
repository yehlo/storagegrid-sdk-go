package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/services/tenantuser"
)

// MockTenantUserService implements tenantuser.ServiceInterface for testing
type MockTenantUserService struct {
	ListFunc        func(ctx context.Context) ([]tenantuser.User, error)
	GetByIDFunc     func(ctx context.Context, id string) (*tenantuser.User, error)
	GetByNameFunc   func(ctx context.Context, name string) (*tenantuser.User, error)
	CreateFunc      func(ctx context.Context, user *tenantuser.User) (*tenantuser.User, error)
	UpdateFunc      func(ctx context.Context, user *tenantuser.User) (*tenantuser.User, error)
	DeleteFunc      func(ctx context.Context, id string) error
	SetPasswordFunc func(ctx context.Context, id string, password string) error
}

func (m *MockTenantUserService) List(ctx context.Context) ([]tenantuser.User, error) {
	if m.ListFunc != nil {
		return m.ListFunc(ctx)
	}
	return []tenantuser.User{}, nil
}

func (m *MockTenantUserService) GetByID(ctx context.Context, id string) (*tenantuser.User, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, id)
	}
	mockID := id
	return &tenantuser.User{ID: &mockID}, nil
}

func (m *MockTenantUserService) GetByName(ctx context.Context, name string) (*tenantuser.User, error) {
	if m.GetByNameFunc != nil {
		return m.GetByNameFunc(ctx, name)
	}
	mockID := "mock-user-id"
	return &tenantuser.User{ID: &mockID, UniqueName: name}, nil
}

func (m *MockTenantUserService) Create(ctx context.Context, user *tenantuser.User) (*tenantuser.User, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, user)
	}
	mockID := "mock-user-id"
	user.ID = &mockID
	return user, nil
}

func (m *MockTenantUserService) Update(ctx context.Context, user *tenantuser.User) (*tenantuser.User, error) {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(ctx, user)
	}
	return user, nil
}

func (m *MockTenantUserService) Delete(ctx context.Context, id string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(ctx, id)
	}
	return nil
}

func (m *MockTenantUserService) SetPassword(ctx context.Context, id string, password string) error {
	if m.SetPasswordFunc != nil {
		return m.SetPasswordFunc(ctx, id, password)
	}
	return nil
}

// Compile-time interface compliance check
var _ tenantuser.ServiceInterface = (*MockTenantUserService)(nil)
