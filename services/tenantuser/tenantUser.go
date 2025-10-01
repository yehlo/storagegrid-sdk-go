package tenantuser

import (
	"context"
	"strings"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
)

const (
	endpoint string = "/org/users"
)

// ServiceInterface defines the contract for tenant user service operations
type ServiceInterface interface {
	List(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	GetByName(ctx context.Context, name string) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id string) error
	SetPassword(ctx context.Context, id string, password string) error
}

type Service struct {
	client services.HTTPClient
}

func NewService(client services.HTTPClient) *Service {
	return &Service{client: client}
}

func (s *Service) List(ctx context.Context) ([]User, error) {
	var response models.Response[[]User]
	if err := s.client.DoParsed(ctx, "GET", endpoint, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*User, error) {
	var response models.Response[*User]
	if err := s.client.DoParsed(ctx, "GET", endpoint+"/"+id, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) GetByName(ctx context.Context, name string) (*User, error) {
	var response models.Response[*User]
	if err := s.client.DoParsed(ctx, "GET", endpoint+"/user/"+name, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) Create(ctx context.Context, user *User) (*User, error) {
	// enforce user/ prefix on user.uniqueName if manually created
	if !strings.HasPrefix(user.UniqueName, "user/") {
		user.UniqueName = "user/" + user.UniqueName
	}

	var response models.Response[*User]
	if err := s.client.DoParsed(ctx, "POST", endpoint, user, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) Update(ctx context.Context, user *User) (*User, error) {
	var response models.Response[*User]
	if err := s.client.DoParsed(ctx, "PUT", endpoint+"/"+*user.ID, user, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.client.DoParsed(ctx, "DELETE", endpoint+"/"+id, nil, nil)
}

func (s *Service) SetPassword(ctx context.Context, id string, password string) error {
	data := map[string]string{"password": password}
	return s.client.DoParsed(ctx, "POST", endpoint+"/"+id+"/change-password", data, nil)
}
