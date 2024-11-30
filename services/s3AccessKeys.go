package services

import (
	"context"
	"fmt"

	models "github.com/yehlo/storagegrid-sdk-go/models"
)

const (
	currentUsers3AccessKeyEndpoint string = "/org/users/current-user/s3-access-keys"
	users3AccessKeyEndpoint        string = "/org/users/%s/s3-access-keys"
)

func getUsers3AccessKeyEndpoint(userId string) string {
	return fmt.Sprintf(users3AccessKeyEndpoint, userId)
}

type s3AccessKeyService struct {
	client HTTPClient
}

func News3AccessKeyService(client HTTPClient) *s3AccessKeyService {
	return &s3AccessKeyService{client: client}
}

func (s *s3AccessKeyService) ListForCurrentUser(ctx context.Context) (*[]models.S3AccessKey, error) {
	response := models.Response{}
	response.Data = &[]models.S3AccessKey{}
	err := s.client.DoParsed(ctx, "GET", currentUsers3AccessKeyEndpoint, nil, &response)
	if err != nil {
		return nil, err
	}

	s3AccessKeys := response.Data.(*[]models.S3AccessKey)

	return s3AccessKeys, nil
}

func (s *s3AccessKeyService) ListForUser(ctx context.Context, userId string) (*[]models.S3AccessKey, error) {
	response := models.Response{}
	response.Data = &[]models.S3AccessKey{}
	err := s.client.DoParsed(ctx, "GET", getUsers3AccessKeyEndpoint(userId), nil, &response)
	if err != nil {
		return nil, err
	}

	s3AccessKeys := response.Data.(*[]models.S3AccessKey)

	return s3AccessKeys, nil
}

func (s *s3AccessKeyService) GetByIdForCurrentUser(ctx context.Context, id string) (*models.S3AccessKey, error) {
	response := models.Response{}
	response.Data = &models.S3AccessKey{}
	err := s.client.DoParsed(ctx, "GET", currentUsers3AccessKeyEndpoint+"/"+id, nil, &response)
	if err != nil {
		return nil, err
	}

	s3AccessKey := response.Data.(*models.S3AccessKey)

	return s3AccessKey, nil
}

func (s *s3AccessKeyService) GetByIdForUser(ctx context.Context, userId string, id string) (*models.S3AccessKey, error) {
	response := models.Response{}
	response.Data = &models.S3AccessKey{}
	err := s.client.DoParsed(ctx, "GET", getUsers3AccessKeyEndpoint(userId)+"/"+id, nil, &response)
	if err != nil {
		return nil, err
	}

	s3AccessKey := response.Data.(*models.S3AccessKey)

	return s3AccessKey, nil
}

func (s *s3AccessKeyService) CreateForCurrentUser(ctx context.Context, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error) {
	response := models.Response{}
	response.Data = &models.S3AccessKey{}
	err := s.client.DoParsed(ctx, "POST", currentUsers3AccessKeyEndpoint, s3AccessKey, &response)
	if err != nil {
		return nil, err
	}

	s3AccessKey = response.Data.(*models.S3AccessKey)

	return s3AccessKey, nil
}

func (s *s3AccessKeyService) CreateForUser(ctx context.Context, userId string, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error) {
	response := models.Response{}
	response.Data = &models.S3AccessKey{}
	err := s.client.DoParsed(ctx, "POST", getUsers3AccessKeyEndpoint(userId), s3AccessKey, &response)
	if err != nil {
		return nil, err
	}

	s3AccessKey = response.Data.(*models.S3AccessKey)

	return s3AccessKey, nil
}

func (s *s3AccessKeyService) DeleteForCurrentUser(ctx context.Context, id string) error {
	err := s.client.DoParsed(ctx, "DELETE", currentUsers3AccessKeyEndpoint+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *s3AccessKeyService) DeleteForUser(ctx context.Context, userId string, id string) error {
	err := s.client.DoParsed(ctx, "DELETE", getUsers3AccessKeyEndpoint(userId)+"/"+id, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
