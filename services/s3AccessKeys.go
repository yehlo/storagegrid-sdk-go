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
	ListForCurrentUser(ctx context.Context) (*[]models.S3AccessKey, error)
	ListForUser(ctx context.Context, userId string) (*[]models.S3AccessKey, error)
	GetByIdForCurrentUser(ctx context.Context, id string) (*models.S3AccessKey, error)
	GetByIdForUser(ctx context.Context, userId string, id string) (*models.S3AccessKey, error)
	CreateForCurrentUser(ctx context.Context, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error)
	CreateForUser(ctx context.Context, userId string, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error)
	DeleteForCurrentUser(ctx context.Context, id string) error
	DeleteForUser(ctx context.Context, userId string, id string) error
}

func getUsers3AccessKeyEndpoint(userId string) string {
	return fmt.Sprintf(users3AccessKeyEndpoint, userId)
}

type S3AccessKeyService struct {
	client HTTPClient
}

func NewS3AccessKeyService(client HTTPClient) *S3AccessKeyService {
	return &S3AccessKeyService{client: client}
}

func (s *S3AccessKeyService) ListForCurrentUser(ctx context.Context) (*[]models.S3AccessKey, error) {
	response := models.Response{}
	response.Data = &[]models.S3AccessKey{}
	err := s.client.DoParsed(ctx, "GET", currentUsers3AccessKeyEndpoint, nil, &response)
	if err != nil {
		return nil, err
	}

	s3AccessKeys := response.Data.(*[]models.S3AccessKey)

	return s3AccessKeys, nil
}

func (s *S3AccessKeyService) ListForUser(ctx context.Context, userId string) (*[]models.S3AccessKey, error) {
	response := models.Response{}
	response.Data = &[]models.S3AccessKey{}
	err := s.client.DoParsed(ctx, "GET", getUsers3AccessKeyEndpoint(userId), nil, &response)
	if err != nil {
		return nil, err
	}

	s3AccessKeys := response.Data.(*[]models.S3AccessKey)

	return s3AccessKeys, nil
}

func (s *S3AccessKeyService) GetByIdForCurrentUser(ctx context.Context, id string) (*models.S3AccessKey, error) {
	response := models.Response{}
	response.Data = &models.S3AccessKey{}
	err := s.client.DoParsed(ctx, "GET", currentUsers3AccessKeyEndpoint+"/"+id, nil, &response)
	if err != nil {
		return nil, err
	}

	s3AccessKey := response.Data.(*models.S3AccessKey)

	return s3AccessKey, nil
}

func (s *S3AccessKeyService) GetByIdForUser(ctx context.Context, userId string, id string) (*models.S3AccessKey, error) {
	response := models.Response{}
	response.Data = &models.S3AccessKey{}
	err := s.client.DoParsed(ctx, "GET", getUsers3AccessKeyEndpoint(userId)+"/"+id, nil, &response)
	if err != nil {
		return nil, err
	}

	s3AccessKey := response.Data.(*models.S3AccessKey)

	return s3AccessKey, nil
}

func (s *S3AccessKeyService) CreateForCurrentUser(ctx context.Context, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error) {
	response := models.Response{}
	response.Data = &models.S3AccessKey{}
	err := s.client.DoParsed(ctx, "POST", currentUsers3AccessKeyEndpoint, s3AccessKey, &response)
	if err != nil {
		return nil, err
	}

	s3AccessKey = response.Data.(*models.S3AccessKey)

	return s3AccessKey, nil
}

func (s *S3AccessKeyService) CreateForUser(ctx context.Context, userId string, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error) {
	response := models.Response{}
	response.Data = &models.S3AccessKey{}
	err := s.client.DoParsed(ctx, "POST", getUsers3AccessKeyEndpoint(userId), s3AccessKey, &response)
	if err != nil {
		return nil, err
	}

	s3AccessKey = response.Data.(*models.S3AccessKey)

	return s3AccessKey, nil
}

func (s *S3AccessKeyService) DeleteForCurrentUser(ctx context.Context, id string) error {
	err := s.client.DoParsed(ctx, "DELETE", currentUsers3AccessKeyEndpoint+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *S3AccessKeyService) DeleteForUser(ctx context.Context, userId string, id string) error {
	err := s.client.DoParsed(ctx, "DELETE", getUsers3AccessKeyEndpoint(userId)+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
