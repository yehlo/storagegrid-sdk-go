package services

import (
	"context"
	"fmt"

	"github.com/yehlo/storagegrid-sdk-go/models"
)

const (
	currentUsers3AccessKeyEndpoint string = "/org/users/current-user/s3-access-keys"
	users3AccessKeyEndpoint        string = "/org/users/%s/s3-access-keys"
)

// S3AccessKeyServiceInterface defines the contract for S3 access key service operations
type S3AccessKeyServiceInterface interface {
	ListForCurrentUser(ctx context.Context) ([]models.S3AccessKey, error)
	ListForUser(ctx context.Context, userID string) ([]models.S3AccessKey, error)
	GetByIDForCurrentUser(ctx context.Context, id string) (*models.S3AccessKey, error)
	GetByIDForUser(ctx context.Context, userID string, id string) (*models.S3AccessKey, error)
	CreateForCurrentUser(ctx context.Context, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error)
	CreateForUser(ctx context.Context, userID string, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error)
	DeleteForCurrentUser(ctx context.Context, id string) error
	DeleteForUser(ctx context.Context, userID string, id string) error
}

func getUsers3AccessKeyEndpoint(userID string) string {
	return fmt.Sprintf(users3AccessKeyEndpoint, userID)
}

type S3AccessKeyService struct {
	client HTTPClient
}

func NewS3AccessKeyService(client HTTPClient) *S3AccessKeyService {
	return &S3AccessKeyService{client: client}
}

func (s *S3AccessKeyService) ListForCurrentUser(ctx context.Context) ([]models.S3AccessKey, error) {
	var response models.Response[[]models.S3AccessKey]
	if err := s.client.DoParsed(ctx, "GET", currentUsers3AccessKeyEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *S3AccessKeyService) ListForUser(ctx context.Context, userID string) ([]models.S3AccessKey, error) {
	var response models.Response[[]models.S3AccessKey]
	if err := s.client.DoParsed(ctx, "GET", getUsers3AccessKeyEndpoint(userID), nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *S3AccessKeyService) GetByIDForCurrentUser(ctx context.Context, id string) (*models.S3AccessKey, error) {
	var response models.Response[*models.S3AccessKey]
	if err := s.client.DoParsed(ctx, "GET", currentUsers3AccessKeyEndpoint+"/"+id, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *S3AccessKeyService) GetByIDForUser(ctx context.Context, userID string, id string) (*models.S3AccessKey, error) {
	var response models.Response[*models.S3AccessKey]
	if err := s.client.DoParsed(ctx, "GET", getUsers3AccessKeyEndpoint(userID)+"/"+id, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *S3AccessKeyService) CreateForCurrentUser(ctx context.Context, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error) {
	var response models.Response[*models.S3AccessKey]
	if err := s.client.DoParsed(ctx, "POST", currentUsers3AccessKeyEndpoint, s3AccessKey, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *S3AccessKeyService) CreateForUser(ctx context.Context, userID string, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error) {
	var response models.Response[*models.S3AccessKey]
	if err := s.client.DoParsed(ctx, "POST", getUsers3AccessKeyEndpoint(userID), s3AccessKey, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *S3AccessKeyService) DeleteForCurrentUser(ctx context.Context, id string) error {
	return s.client.DoParsed(ctx, "DELETE", currentUsers3AccessKeyEndpoint+"/"+id, nil, nil)
}

func (s *S3AccessKeyService) DeleteForUser(ctx context.Context, userID string, id string) error {
	return s.client.DoParsed(ctx, "DELETE", getUsers3AccessKeyEndpoint(userID)+"/"+id, nil, nil)
}
