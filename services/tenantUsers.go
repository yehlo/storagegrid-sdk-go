package services

import (
	"context"
	"strings"

	models "github.com/yehlo/storagegrid-sdk-go/models"
)

const (
	tenantUserEndpoint string = "/org/users"
)

type TenantUserService struct {
	client HTTPClient
}

func NewTenantUserService(client HTTPClient) *TenantUserService {
	return &TenantUserService{client: client}
}

func (s *TenantUserService) List(ctx context.Context) (*[]models.User, error) {
	response := models.Response{}
	response.Data = &[]models.User{}
	err := s.client.DoParsed(ctx, "GET", tenantUserEndpoint, nil, &response)
	if err != nil {
		return nil, err
	}

	users := response.Data.(*[]models.User)

	return users, nil
}

func (s *TenantUserService) GetById(ctx context.Context, id string) (*models.User, error) {
	response := models.Response{}
	response.Data = &models.User{}
	err := s.client.DoParsed(ctx, "GET", tenantUserEndpoint+"/"+id, nil, &response)
	if err != nil {
		return nil, err
	}

	user := response.Data.(*models.User)

	return user, nil
}

func (s *TenantUserService) GetByName(ctx context.Context, name string) (*models.User, error) {
	response := models.Response{}
	response.Data = &models.User{}
	err := s.client.DoParsed(ctx, "GET", tenantUserEndpoint+"/user/"+name, nil, &response)
	if err != nil {
		return nil, err
	}

	user := response.Data.(*models.User)

	return user, nil
}

func (s *TenantUserService) Create(ctx context.Context, user *models.User) (*models.User, error) {
	// enforce user/ prefix on user.uniqueName if manually created
	if !strings.HasPrefix(user.UniqueName, "user/") {
		user.UniqueName = "user/" + user.UniqueName
	}

	response := models.Response{}
	response.Data = &models.User{}
	err := s.client.DoParsed(ctx, "POST", tenantUserEndpoint, user, &response)
	if err != nil {
		return nil, err
	}

	user = response.Data.(*models.User)

	return user, nil
}

func (s *TenantUserService) Update(ctx context.Context, user *models.User) (*models.User, error) {
	response := models.Response{}
	response.Data = &models.User{}
	err := s.client.DoParsed(ctx, "PUT", tenantUserEndpoint+"/"+*user.Id, user, &response)
	if err != nil {
		return nil, err
	}

	user = response.Data.(*models.User)

	return user, nil
}

func (s *TenantUserService) Delete(ctx context.Context, id string) error {
	err := s.client.DoParsed(ctx, "DELETE", tenantUserEndpoint+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
