package accesskeys

import (
	"context"
	"fmt"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
)

const (
	currentUsersEndpoint string = "/org/users/current-user/s3-access-keys"
	usersEndpoint        string = "/org/users/%s/s3-access-keys"
)

// ServiceInterface defines the contract for S3 access key service operations
type ServiceInterface interface {
	ListForCurrentUser(ctx context.Context) ([]AccessKey, error)
	ListForUser(ctx context.Context, userID string) ([]AccessKey, error)
	GetByIDForCurrentUser(ctx context.Context, id string) (*AccessKey, error)
	GetByIDForUser(ctx context.Context, userID string, id string) (*AccessKey, error)
	CreateForCurrentUser(ctx context.Context, s3AccessKey *AccessKey) (*AccessKey, error)
	CreateForUser(ctx context.Context, userID string, s3AccessKey *AccessKey) (*AccessKey, error)
	DeleteForCurrentUser(ctx context.Context, id string) error
	DeleteForUser(ctx context.Context, userID string, id string) error
}

func getUsersEndpoint(userID string) string {
	return fmt.Sprintf(usersEndpoint, userID)
}

type Service struct {
	services.HTTPClient
}

// NewService returns a new accesskey service using the provided client
func NewService(client services.HTTPClient) *Service {
	return &Service{client}
}

// ListForCurrentUser lists all access keys for the current user
func (s *Service) ListForCurrentUser(ctx context.Context) ([]AccessKey, error) {
	var response models.Response[[]AccessKey]
	if err := s.DoParsed(ctx, "GET", currentUsersEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// ListForUser lists all access keys for the given user id
func (s *Service) ListForUser(ctx context.Context, userID string) ([]AccessKey, error) {
	var response models.Response[[]AccessKey]
	if err := s.DoParsed(ctx, "GET", getUsersEndpoint(userID), nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetByIDForCurrentUser retrieves the access key specified for the current user
func (s *Service) GetByIDForCurrentUser(ctx context.Context, id string) (*AccessKey, error) {
	var response models.Response[*AccessKey]
	if err := s.DoParsed(ctx, "GET", currentUsersEndpoint+"/"+id, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetByIDForUser retrieves the access key specified for the given user id
func (s *Service) GetByIDForUser(ctx context.Context, userID string, id string) (*AccessKey, error) {
	var response models.Response[*AccessKey]
	if err := s.DoParsed(ctx, "GET", getUsersEndpoint(userID)+"/"+id, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// CreateForCurrentUser creates a new access key for the current user
func (s *Service) CreateForCurrentUser(ctx context.Context, s3AccessKey *AccessKey) (*AccessKey, error) {
	var response models.Response[*AccessKey]
	if err := s.DoParsed(ctx, "POST", currentUsersEndpoint, s3AccessKey, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// CreateForUser creates a new access key for the given user id
func (s *Service) CreateForUser(ctx context.Context, userID string, accessKey *AccessKey) (*AccessKey, error) {
	var response models.Response[*AccessKey]
	if err := s.DoParsed(ctx, "POST", getUsersEndpoint(userID), accessKey, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// DeleteForCurrentUser deletes the access key specified for the current user
func (s *Service) DeleteForCurrentUser(ctx context.Context, id string) error {
	return s.DoParsed(ctx, "DELETE", currentUsersEndpoint+"/"+id, nil, nil)
}

// DeleteForUser deletes the access key specified for the given user id
func (s *Service) DeleteForUser(ctx context.Context, userID string, id string) error {
	return s.DoParsed(ctx, "DELETE", getUsersEndpoint(userID)+"/"+id, nil, nil)
}
