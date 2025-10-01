package services

import (
	"context"
	"strings"

	"github.com/yehlo/storagegrid-sdk-go/models"
)

const (
	tenantUserEndpoint string = "/org/users"
)

// TenantUserServiceInterface defines the contract for tenant user service operations
type TenantUserServiceInterface interface {
	List(ctx context.Context) ([]models.User, error)
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByName(ctx context.Context, name string) (*models.User, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, id string) error
	SetPassword(ctx context.Context, id string, password string) error
}

type TenantUserService struct {
	client HTTPClient
}

func NewTenantUserService(client HTTPClient) *TenantUserService {
	return &TenantUserService{client: client}
}

func (s *TenantUserService) List(ctx context.Context) ([]models.User, error) {
	var response models.Response[[]models.User]
	if err := s.client.DoParsed(ctx, "GET", tenantUserEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *TenantUserService) GetByID(ctx context.Context, id string) (*models.User, error) {
	var response models.Response[*models.User]
	if err := s.client.DoParsed(ctx, "GET", tenantUserEndpoint+"/"+id, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *TenantUserService) GetByName(ctx context.Context, name string) (*models.User, error) {
	var response models.Response[*models.User]
	if err := s.client.DoParsed(ctx, "GET", tenantUserEndpoint+"/user/"+name, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *TenantUserService) Create(ctx context.Context, user *models.User) (*models.User, error) {
	// enforce user/ prefix on user.uniqueName if manually created
	if !strings.HasPrefix(user.UniqueName, "user/") {
		user.UniqueName = "user/" + user.UniqueName
	}

	var response models.Response[*models.User]
	if err := s.client.DoParsed(ctx, "POST", tenantUserEndpoint, user, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *TenantUserService) Update(ctx context.Context, user *models.User) (*models.User, error) {
	var response models.Response[*models.User]
	if err := s.client.DoParsed(ctx, "PUT", tenantUserEndpoint+"/"+*user.ID, user, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *TenantUserService) Delete(ctx context.Context, id string) error {
	return s.client.DoParsed(ctx, "DELETE", tenantUserEndpoint+"/"+id, nil, nil)
}

func (s *TenantUserService) SetPassword(ctx context.Context, id string, password string) error {
	data := map[string]string{"password": password}
	return s.client.DoParsed(ctx, "POST", tenantUserEndpoint+"/"+id+"/change-password", data, nil)
}
