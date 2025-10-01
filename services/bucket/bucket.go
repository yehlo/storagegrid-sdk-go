package bucket

import (
	"context"
	"fmt"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
	"github.com/yehlo/storagegrid-sdk-go/services/tenantusage"
)

const (
	bucketEndpoint      string = "/org/containers"
	tenantUsageEndpoint string = "/org/usage"
)

// ServiceInterface defines the contract for bucket service operations
type ServiceInterface interface {
	List(ctx context.Context) ([]Bucket, error)
	GetByName(ctx context.Context, name string) (*Bucket, error)
	Create(ctx context.Context, bucket *Bucket) (*Bucket, error)
	GetUsage(ctx context.Context, name string) (*tenantusage.BucketStats, error)
	Delete(ctx context.Context, name string) error
	Drain(ctx context.Context, name string) (*DeleteObjectStatus, error)
	DrainStatus(ctx context.Context, name string) (*DeleteObjectStatus, error)
}

type Service struct {
	client services.HTTPClient
}

func NewService(client services.HTTPClient) *Service {
	return &Service{client: client}
}

func (s *Service) List(ctx context.Context) ([]Bucket, error) {
	var response models.Response[[]Bucket]
	err := s.client.DoParsed(ctx, "GET", bucketEndpoint, nil, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) GetByName(ctx context.Context, name string) (*Bucket, error) {
	// the bucket endpoint doesn't have a simple get by name, so we have to list all buckets and find the one we want
	buckets, err := s.List(ctx)
	if err != nil {
		return nil, err
	}

	for _, bucket := range buckets {
		if bucket.Name == name {
			return &bucket, nil
		}
	}

	return nil, fmt.Errorf("bucket with name %s not found", name)
}

func (s *Service) Create(ctx context.Context, bucket *Bucket) (*Bucket, error) {
	var response models.Response[*Bucket]
	if err := s.client.DoParsed(ctx, "POST", bucketEndpoint, bucket, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) GetUsage(ctx context.Context, name string) (*tenantusage.BucketStats, error) {
	var response models.Response[tenantusage.TenantUsage]
	if err := s.client.DoParsed(ctx, "GET", tenantUsageEndpoint, nil, &response); err != nil {
		return nil, err
	}

	for _, bucket := range response.Data.Buckets {
		if bucket.Name != nil && *bucket.Name == name {
			return &bucket, nil
		}
	}

	return nil, fmt.Errorf("usage for bucket with name %s not found", name)
}

func (s *Service) Delete(ctx context.Context, name string) error {
	return s.client.DoParsed(ctx, "DELETE", bucketEndpoint+"/"+name, nil, nil)
}

// Drain a bucket by name. This will delete all objects in the bucket but leave the bucket itself intact.
func (s *Service) Drain(ctx context.Context, name string) (*DeleteObjectStatus, error) {
	var response models.Response[*DeleteObjectStatus]
	body := map[string]string{"deleteObjects": "true"}

	if err := s.client.DoParsed(ctx, "POST", bucketEndpoint+"/"+name+"/delete-objects", body, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Service) DrainStatus(ctx context.Context, name string) (*DeleteObjectStatus, error) {
	var response models.Response[*DeleteObjectStatus]

	if err := s.client.DoParsed(ctx, "GET", bucketEndpoint+"/"+name+"/delete-objects", nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
