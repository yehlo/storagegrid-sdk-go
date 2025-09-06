package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
)

// MockS3AccessKeyService implements services.S3AccessKeyServiceInterface for testing
type MockS3AccessKeyService struct {
	ListForCurrentUserFunc    func(ctx context.Context) (*[]models.S3AccessKey, error)
	ListForUserFunc           func(ctx context.Context, userId string) (*[]models.S3AccessKey, error)
	GetByIdForCurrentUserFunc func(ctx context.Context, id string) (*models.S3AccessKey, error)
	GetByIdForUserFunc        func(ctx context.Context, userId string, id string) (*models.S3AccessKey, error)
	CreateForCurrentUserFunc  func(ctx context.Context, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error)
	CreateForUserFunc         func(ctx context.Context, userId string, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error)
	DeleteForCurrentUserFunc  func(ctx context.Context, id string) error
	DeleteForUserFunc         func(ctx context.Context, userId string, id string) error
}

func (m *MockS3AccessKeyService) ListForCurrentUser(ctx context.Context) (*[]models.S3AccessKey, error) {
	if m.ListForCurrentUserFunc != nil {
		return m.ListForCurrentUserFunc(ctx)
	}
	return &[]models.S3AccessKey{}, nil
}

func (m *MockS3AccessKeyService) ListForUser(ctx context.Context, userId string) (*[]models.S3AccessKey, error) {
	if m.ListForUserFunc != nil {
		return m.ListForUserFunc(ctx, userId)
	}
	return &[]models.S3AccessKey{}, nil
}

func (m *MockS3AccessKeyService) GetByIdForCurrentUser(ctx context.Context, id string) (*models.S3AccessKey, error) {
	if m.GetByIdForCurrentUserFunc != nil {
		return m.GetByIdForCurrentUserFunc(ctx, id)
	}
	mockId := id
	return &models.S3AccessKey{Id: &mockId}, nil
}

func (m *MockS3AccessKeyService) GetByIdForUser(ctx context.Context, userId string, id string) (*models.S3AccessKey, error) {
	if m.GetByIdForUserFunc != nil {
		return m.GetByIdForUserFunc(ctx, userId, id)
	}
	mockId := id
	return &models.S3AccessKey{Id: &mockId}, nil
}

func (m *MockS3AccessKeyService) CreateForCurrentUser(ctx context.Context, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error) {
	if m.CreateForCurrentUserFunc != nil {
		return m.CreateForCurrentUserFunc(ctx, s3AccessKey)
	}
	mockId := "mock-s3-key-id"
	s3AccessKey.Id = &mockId
	return s3AccessKey, nil
}

func (m *MockS3AccessKeyService) CreateForUser(ctx context.Context, userId string, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error) {
	if m.CreateForUserFunc != nil {
		return m.CreateForUserFunc(ctx, userId, s3AccessKey)
	}
	mockId := "mock-s3-key-id"
	s3AccessKey.Id = &mockId
	return s3AccessKey, nil
}

func (m *MockS3AccessKeyService) DeleteForCurrentUser(ctx context.Context, id string) error {
	if m.DeleteForCurrentUserFunc != nil {
		return m.DeleteForCurrentUserFunc(ctx, id)
	}
	return nil
}

func (m *MockS3AccessKeyService) DeleteForUser(ctx context.Context, userId string, id string) error {
	if m.DeleteForUserFunc != nil {
		return m.DeleteForUserFunc(ctx, userId, id)
	}
	return nil
}

// Compile-time interface compliance check
var _ services.S3AccessKeyServiceInterface = (*MockS3AccessKeyService)(nil)
