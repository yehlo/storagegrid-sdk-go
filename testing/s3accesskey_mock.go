package testing

import (
	"context"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
)

// MockS3AccessKeyService implements services.S3AccessKeyServiceInterface for testing
type MockS3AccessKeyService struct {
	ListForCurrentUserFunc    func(ctx context.Context) ([]models.S3AccessKey, error)
	ListForUserFunc           func(ctx context.Context, userID string) ([]models.S3AccessKey, error)
	GetByIDForCurrentUserFunc func(ctx context.Context, id string) (*models.S3AccessKey, error)
	GetByIDForUserFunc        func(ctx context.Context, userID string, id string) (*models.S3AccessKey, error)
	CreateForCurrentUserFunc  func(ctx context.Context, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error)
	CreateForUserFunc         func(ctx context.Context, userID string, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error)
	DeleteForCurrentUserFunc  func(ctx context.Context, id string) error
	DeleteForUserFunc         func(ctx context.Context, userID string, id string) error
}

func (m *MockS3AccessKeyService) ListForCurrentUser(ctx context.Context) ([]models.S3AccessKey, error) {
	if m.ListForCurrentUserFunc != nil {
		return m.ListForCurrentUserFunc(ctx)
	}
	return []models.S3AccessKey{}, nil
}

func (m *MockS3AccessKeyService) ListForUser(ctx context.Context, userID string) ([]models.S3AccessKey, error) {
	if m.ListForUserFunc != nil {
		return m.ListForUserFunc(ctx, userID)
	}
	return []models.S3AccessKey{}, nil
}

func (m *MockS3AccessKeyService) GetByIDForCurrentUser(ctx context.Context, id string) (*models.S3AccessKey, error) {
	if m.GetByIDForCurrentUserFunc != nil {
		return m.GetByIDForCurrentUserFunc(ctx, id)
	}
	mockID := id
	return &models.S3AccessKey{ID: &mockID}, nil
}

func (m *MockS3AccessKeyService) GetByIDForUser(ctx context.Context, userID string, id string) (*models.S3AccessKey, error) {
	if m.GetByIDForUserFunc != nil {
		return m.GetByIDForUserFunc(ctx, userID, id)
	}
	mockID := id
	return &models.S3AccessKey{ID: &mockID}, nil
}

func (m *MockS3AccessKeyService) CreateForCurrentUser(ctx context.Context, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error) {
	if m.CreateForCurrentUserFunc != nil {
		return m.CreateForCurrentUserFunc(ctx, s3AccessKey)
	}
	mockID := "mock-s3-key-id"
	s3AccessKey.ID = &mockID
	return s3AccessKey, nil
}

func (m *MockS3AccessKeyService) CreateForUser(ctx context.Context, userID string, s3AccessKey *models.S3AccessKey) (*models.S3AccessKey, error) {
	if m.CreateForUserFunc != nil {
		return m.CreateForUserFunc(ctx, userID, s3AccessKey)
	}
	mockID := "mock-s3-key-id"
	s3AccessKey.ID = &mockID
	return s3AccessKey, nil
}

func (m *MockS3AccessKeyService) DeleteForCurrentUser(ctx context.Context, id string) error {
	if m.DeleteForCurrentUserFunc != nil {
		return m.DeleteForCurrentUserFunc(ctx, id)
	}
	return nil
}

func (m *MockS3AccessKeyService) DeleteForUser(ctx context.Context, userID string, id string) error {
	if m.DeleteForUserFunc != nil {
		return m.DeleteForUserFunc(ctx, userID, id)
	}
	return nil
}

// Compile-time interface compliance check
var _ services.S3AccessKeyServiceInterface = (*MockS3AccessKeyService)(nil)
