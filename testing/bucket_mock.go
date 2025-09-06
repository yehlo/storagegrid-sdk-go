package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
)

// MockBucketService implements services.BucketServiceInterface for testing
type MockBucketService struct {
	ListFunc        func(ctx context.Context) (*[]models.Bucket, error)
	GetByNameFunc   func(ctx context.Context, name string) (*models.Bucket, error)
	CreateFunc      func(ctx context.Context, bucket *models.Bucket) (*models.Bucket, error)
	GetUsageFunc    func(ctx context.Context, name string) (*models.BucketStats, error)
	DeleteFunc      func(ctx context.Context, name string) error
	DrainFunc       func(ctx context.Context, name string) (*models.BucketDeleteObjectStatus, error)
	DrainStatusFunc func(ctx context.Context, name string) (*models.BucketDeleteObjectStatus, error)
}

func (m *MockBucketService) List(ctx context.Context) (*[]models.Bucket, error) {
	if m.ListFunc != nil {
		return m.ListFunc(ctx)
	}
	return &[]models.Bucket{}, nil
}

func (m *MockBucketService) GetByName(ctx context.Context, name string) (*models.Bucket, error) {
	if m.GetByNameFunc != nil {
		return m.GetByNameFunc(ctx, name)
	}
	return &models.Bucket{Name: name}, nil
}

func (m *MockBucketService) Create(ctx context.Context, bucket *models.Bucket) (*models.Bucket, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, bucket)
	}
	return bucket, nil
}

func (m *MockBucketService) GetUsage(ctx context.Context, name string) (*models.BucketStats, error) {
	if m.GetUsageFunc != nil {
		return m.GetUsageFunc(ctx, name)
	}
	bucketName := name
	return &models.BucketStats{Name: &bucketName}, nil
}

func (m *MockBucketService) Delete(ctx context.Context, name string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(ctx, name)
	}
	return nil
}

func (m *MockBucketService) Drain(ctx context.Context, name string) (*models.BucketDeleteObjectStatus, error) {
	if m.DrainFunc != nil {
		return m.DrainFunc(ctx, name)
	}
	return &models.BucketDeleteObjectStatus{}, nil
}

func (m *MockBucketService) DrainStatus(ctx context.Context, name string) (*models.BucketDeleteObjectStatus, error) {
	if m.DrainStatusFunc != nil {
		return m.DrainStatusFunc(ctx, name)
	}
	return &models.BucketDeleteObjectStatus{}, nil
}

// Compile-time interface compliance check
var _ services.BucketServiceInterface = (*MockBucketService)(nil)
