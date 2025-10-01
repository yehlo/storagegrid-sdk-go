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
	ListForCurrentUser(ctx context.Context) ([]S3AccessKey, error)
	ListForUser(ctx context.Context, userID string) ([]S3AccessKey, error)
	GetByIDForCurrentUser(ctx context.Context, id string) (*S3AccessKey, error)
	GetByIDForUser(ctx context.Context, userID string, id string) (*S3AccessKey, error)
	CreateForCurrentUser(ctx context.Context, s3AccessKey *S3AccessKey) (*S3AccessKey, error)
	CreateForUser(ctx context.Context, userID string, s3AccessKey *S3AccessKey) (*S3AccessKey, error)
	DeleteForCurrentUser(ctx context.Context, id string) error
	DeleteForUser(ctx context.Context, userID string, id string) error
}

func getUsersEndpoint(userID string) string {
	return fmt.Sprintf(usersEndpoint, userID)
}

type Service struct {
	client services.HTTPClient
}

func NewService(client services.HTTPClient) *Service {
	return &Service{client: client}
}

func (s *Service) ListForCurrentUser(ctx context.Context) ([]S3AccessKey, error) {
	var response models.Response[[]S3AccessKey]
	if err := s.client.DoParsed(ctx, "GET", currentUsersEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) ListForUser(ctx context.Context, userID string) ([]S3AccessKey, error) {
	var response models.Response[[]S3AccessKey]
	if err := s.client.DoParsed(ctx, "GET", getUsersEndpoint(userID), nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) GetByIDForCurrentUser(ctx context.Context, id string) (*S3AccessKey, error) {
	var response models.Response[*S3AccessKey]
	if err := s.client.DoParsed(ctx, "GET", currentUsersEndpoint+"/"+id, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) GetByIDForUser(ctx context.Context, userID string, id string) (*S3AccessKey, error) {
	var response models.Response[*S3AccessKey]
	if err := s.client.DoParsed(ctx, "GET", getUsersEndpoint(userID)+"/"+id, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) CreateForCurrentUser(ctx context.Context, s3AccessKey *S3AccessKey) (*S3AccessKey, error) {
	var response models.Response[*S3AccessKey]
	if err := s.client.DoParsed(ctx, "POST", currentUsersEndpoint, s3AccessKey, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) CreateForUser(ctx context.Context, userID string, s3AccessKey *S3AccessKey) (*S3AccessKey, error) {
	var response models.Response[*S3AccessKey]
	if err := s.client.DoParsed(ctx, "POST", getUsersEndpoint(userID), s3AccessKey, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) DeleteForCurrentUser(ctx context.Context, id string) error {
	return s.client.DoParsed(ctx, "DELETE", currentUsersEndpoint+"/"+id, nil, nil)
}

func (s *Service) DeleteForUser(ctx context.Context, userID string, id string) error {
	return s.client.DoParsed(ctx, "DELETE", getUsersEndpoint(userID)+"/"+id, nil, nil)
}
