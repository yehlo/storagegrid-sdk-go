package services

import (
	"context"
	"fmt"

	"github.com/yehlo/storagegrid-sdk-go/models"
)

const (
	bucketEndpoint      string = "/org/containers"
	tenantUsageEndpoint string = "/org/usage"
)

// BucketServiceInterface defines the contract for bucket service operations
type BucketServiceInterface interface {
	List(ctx context.Context) (*[]models.Bucket, error)
	GetByName(ctx context.Context, name string) (*models.Bucket, error)
	Create(ctx context.Context, bucket *models.Bucket) (*models.Bucket, error)
	GetUsage(ctx context.Context, name string) (*models.BucketStats, error)
	Delete(ctx context.Context, name string) error
	Drain(ctx context.Context, name string) (*models.BucketDeleteObjectStatus, error)
	DrainStatus(ctx context.Context, name string) (*models.BucketDeleteObjectStatus, error)
}

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
	// the bucket endpoint doesn't have a simple get by name, so we have to list all buckets and find the one we want
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

func (s *BucketService) GetUsage(ctx context.Context, name string) (*models.BucketStats, error) {
	response := models.Response{}
	response.Data = &models.TenantUsage{}
	err := s.client.DoParsed(ctx, "GET", tenantUsageEndpoint, nil, &response)
	if err != nil {
		return nil, err
	}

	tenantUsage := response.Data.(*models.TenantUsage)

	for _, bucket := range tenantUsage.Buckets {
		if *bucket.Name == name {
			return bucket, nil
		}
	}

	return nil, fmt.Errorf("usage for bucket with name %s not found", name)
}

func (s *BucketService) Delete(ctx context.Context, name string) error {
	err := s.client.DoParsed(ctx, "DELETE", bucketEndpoint+"/"+name, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// Drain a bucket by name. This will delete all objects in the bucket but leave the bucket itself intact.
func (s *BucketService) Drain(ctx context.Context, name string) (*models.BucketDeleteObjectStatus, error) {
	response := models.Response{}
	response.Data = &models.BucketDeleteObjectStatus{}
	body := map[string]string{"deleteObjects": "true"}

	err := s.client.DoParsed(ctx, "POST", bucketEndpoint+"/"+name+"/delete-objects", body, &response)
	if err != nil {
		return nil, err
	}

	deleteObjectStatus := response.Data.(*models.BucketDeleteObjectStatus)

	return deleteObjectStatus, nil
}

func (s *BucketService) DrainStatus(ctx context.Context, name string) (*models.BucketDeleteObjectStatus, error) {
	response := models.Response{}
	response.Data = &models.BucketDeleteObjectStatus{}

	err := s.client.DoParsed(ctx, "GET", bucketEndpoint+"/"+name+"/delete-objects", nil, &response)
	if err != nil {
		return nil, err
	}

	deleteObjectStatus := response.Data.(*models.BucketDeleteObjectStatus)

	return deleteObjectStatus, nil
}
