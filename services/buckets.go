package services

import (
	"context"
	"fmt"

	models "github.com/yehlo/storagegrid-sdk-go/models"
)

const (
	bucketEndpoint string = "/org/containers"
)

type BucketService struct {
	client HTTPClient
}

func NewBucketService(client HTTPClient) *BucketService {
	return &BucketService{client: client}
}

func (s *BucketService) List(ctx context.Context) (*[]models.Bucket, error) {
	response := models.Response{}
	response.Data = &[]models.Bucket{}
	err := s.client.DoParsed(ctx, "GET", bucketEndpoint, nil, &response)
	if err != nil {
		return nil, err
	}

	buckets := response.Data.(*[]models.Bucket)

	return buckets, nil
}

func (s *BucketService) GetByName(ctx context.Context, name string) (*models.Bucket, error) {
	buckets, err := s.List(ctx)
	if err != nil {
		return nil, err
	}

	for _, bucket := range *buckets {
		if bucket.Name == name {
			return &bucket, nil
		}
	}

	return nil, fmt.Errorf("bucket with name %s not found", name)
}

func (s *BucketService) Create(ctx context.Context, bucket *models.Bucket) (*models.Bucket, error) {
	response := models.Response{}
	response.Data = &models.Bucket{}
	err := s.client.DoParsed(ctx, "POST", bucketEndpoint, bucket, &response)
	if err != nil {
		return nil, err
	}

	bucket = response.Data.(*models.Bucket)

	return bucket, nil
}

func (s *BucketService) Delete(ctx context.Context, name string) error {
	err := s.client.DoParsed(ctx, "DELETE", bucketEndpoint+"/"+name, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
