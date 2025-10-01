package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/services/bucket"
	"github.com/yehlo/storagegrid-sdk-go/services/tenantusage"
)

// MockBucketService implements services.BucketServiceInterface for testing
type MockBucketService struct {
	ListFunc        func(ctx context.Context) ([]bucket.Bucket, error)
	GetByNameFunc   func(ctx context.Context, name string) (*bucket.Bucket, error)
	CreateFunc      func(ctx context.Context, bucket *bucket.Bucket) (*bucket.Bucket, error)
	GetUsageFunc    func(ctx context.Context, name string) (*tenantusage.BucketStats, error)
	DeleteFunc      func(ctx context.Context, name string) error
	DrainFunc       func(ctx context.Context, name string) (*bucket.DeleteObjectStatus, error)
	DrainStatusFunc func(ctx context.Context, name string) (*bucket.DeleteObjectStatus, error)
}

func (m *MockBucketService) List(ctx context.Context) ([]bucket.Bucket, error) {
	if m.ListFunc != nil {
		return m.ListFunc(ctx)
	}
	return []bucket.Bucket{}, nil
}

func (m *MockBucketService) GetByName(ctx context.Context, name string) (*bucket.Bucket, error) {
	if m.GetByNameFunc != nil {
		return m.GetByNameFunc(ctx, name)
	}
	return &bucket.Bucket{Name: name}, nil
}

func (m *MockBucketService) Create(ctx context.Context, bucket *bucket.Bucket) (*bucket.Bucket, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, bucket)
	}
	return bucket, nil
}

func (m *MockBucketService) GetUsage(ctx context.Context, name string) (*tenantusage.BucketStats, error) {
	if m.GetUsageFunc != nil {
		return m.GetUsageFunc(ctx, name)
	}
	bucketName := name
	return &tenantusage.BucketStats{Name: &bucketName}, nil
}

func (m *MockBucketService) Delete(ctx context.Context, name string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(ctx, name)
	}
	return nil
}

func (m *MockBucketService) Drain(ctx context.Context, name string) (*bucket.DeleteObjectStatus, error) {
	if m.DrainFunc != nil {
		return m.DrainFunc(ctx, name)
	}
	return &bucket.DeleteObjectStatus{}, nil
}

func (m *MockBucketService) DrainStatus(ctx context.Context, name string) (*bucket.DeleteObjectStatus, error) {
	if m.DrainStatusFunc != nil {
		return m.DrainStatusFunc(ctx, name)
	}
	return &bucket.DeleteObjectStatus{}, nil
}

// Compile-time interface compliance check
var _ bucket.ServiceInterface = (*MockBucketService)(nil)
