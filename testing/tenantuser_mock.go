package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
)

// MockTenantUserService implements services.TenantUserServiceInterface for testing
type MockTenantUserService struct {
	ListFunc        func(ctx context.Context) (*[]models.User, error)
	GetByIdFunc     func(ctx context.Context, id string) (*models.User, error)
	GetByNameFunc   func(ctx context.Context, name string) (*models.User, error)
	CreateFunc      func(ctx context.Context, user *models.User) (*models.User, error)
	UpdateFunc      func(ctx context.Context, user *models.User) (*models.User, error)
	DeleteFunc      func(ctx context.Context, id string) error
	SetPasswordFunc func(ctx context.Context, id string, password string) error
}

func (m *MockTenantUserService) List(ctx context.Context) (*[]models.User, error) {
	if m.ListFunc != nil {
		return m.ListFunc(ctx)
	}
	return &[]models.User{}, nil
}

func (m *MockTenantUserService) GetById(ctx context.Context, id string) (*models.User, error) {
	if m.GetByIdFunc != nil {
		return m.GetByIdFunc(ctx, id)
	}
	mockId := id
	return &models.User{Id: &mockId}, nil
}

func (m *MockTenantUserService) GetByName(ctx context.Context, name string) (*models.User, error) {
	if m.GetByNameFunc != nil {
		return m.GetByNameFunc(ctx, name)
	}
	mockId := "mock-user-id"
	return &models.User{Id: &mockId, UniqueName: name}, nil
}

func (m *MockTenantUserService) Create(ctx context.Context, user *models.User) (*models.User, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, user)
	}
	mockId := "mock-user-id"
	user.Id = &mockId
	return user, nil
}

func (m *MockTenantUserService) Update(ctx context.Context, user *models.User) (*models.User, error) {
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
var _ services.TenantUserServiceInterface = (*MockTenantUserService)(nil)
