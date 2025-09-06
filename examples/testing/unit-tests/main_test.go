package main

import (
	"context"
	"errors"
	"testing"

	"github.com/yehlo/storagegrid-sdk-go/models"
	"github.com/yehlo/storagegrid-sdk-go/services"
	sgTesting "github.com/yehlo/storagegrid-sdk-go/testing"
)

// Example service that depends on StorageGRID SDK
type TenantManager struct {
	tenantService services.TenantServiceInterface
	bucketService services.BucketServiceInterface
}

func NewTenantManager(tenantService services.TenantServiceInterface, bucketService services.BucketServiceInterface) *TenantManager {
	return &TenantManager{
		tenantService: tenantService,
		bucketService: bucketService,
	}
}

func (tm *TenantManager) CreateTenantWithBucket(ctx context.Context, tenantName, bucketName string) error {
	// Create tenant
	tenant := &models.Tenant{
		Name:         &tenantName,
		Capabilities: []string{"s3", "management"},
	}

	createdTenant, err := tm.tenantService.Create(ctx, tenant)
	if err != nil {
		return err
	}

	// Validate tenant was created
	if createdTenant.Id == "" {
		return errors.New("tenant creation failed: no ID returned")
	}

	// Create bucket (this would normally require a tenant client)
	bucket := &models.Bucket{
		Name:   bucketName,
		Region: "us-east-1",
	}

	_, err = tm.bucketService.Create(ctx, bucket)
	return err
}

func (tm *TenantManager) GetTenantUsageSummary(ctx context.Context, tenantID string) (*UsageSummary, error) {
	// Get tenant usage
	usage, err := tm.tenantService.GetUsage(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	summary := &UsageSummary{
		TenantID:     tenantID,
		TotalBytes:   0,
		TotalObjects: 0,
		BucketCount:  len(usage.Buckets),
	}

	if usage.DataBytes != nil {
		summary.TotalBytes = *usage.DataBytes
	}
	if usage.ObjectCount != nil {
		summary.TotalObjects = *usage.ObjectCount
	}

	return summary, nil
}

type UsageSummary struct {
	TenantID     string
	TotalBytes   int64
	TotalObjects int64
	BucketCount  int
}

// Test functions demonstrating mock usage
func TestTenantManager_CreateTenantWithBucket(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name          string
		tenantName    string
		bucketName    string
		setupMocks    func(*sgTesting.MockTenantService, *sgTesting.MockBucketService)
		expectError   bool
		errorContains string
	}{
		{
			name:       "successful creation",
			tenantName: "test-tenant",
			bucketName: "test-bucket",
			setupMocks: func(mockTenant *sgTesting.MockTenantService, mockBucket *sgTesting.MockBucketService) {
				// Mock successful tenant creation
				mockTenant.CreateFunc = func(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error) {
					tenant.Id = "tenant-123"
					return tenant, nil
				}

				// Mock successful bucket creation
				mockBucket.CreateFunc = func(ctx context.Context, bucket *models.Bucket) (*models.Bucket, error) {
					return bucket, nil
				}
			},
			expectError: false,
		},
		{
			name:       "tenant creation fails",
			tenantName: "failing-tenant",
			bucketName: "test-bucket",
			setupMocks: func(mockTenant *sgTesting.MockTenantService, mockBucket *sgTesting.MockBucketService) {
				// Mock tenant creation failure
				mockTenant.CreateFunc = func(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error) {
					return nil, errors.New("tenant creation failed")
				}
			},
			expectError:   true,
			errorContains: "tenant creation failed",
		},
		{
			name:       "bucket creation fails",
			tenantName: "test-tenant",
			bucketName: "failing-bucket",
			setupMocks: func(mockTenant *sgTesting.MockTenantService, mockBucket *sgTesting.MockBucketService) {
				// Mock successful tenant creation
				mockTenant.CreateFunc = func(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error) {
					tenant.Id = "tenant-123"
					return tenant, nil
				}

				// Mock bucket creation failure
				mockBucket.CreateFunc = func(ctx context.Context, bucket *models.Bucket) (*models.Bucket, error) {
					return nil, errors.New("bucket creation failed")
				}
			},
			expectError:   true,
			errorContains: "bucket creation failed",
		},
		{
			name:       "tenant created without ID",
			tenantName: "no-id-tenant",
			bucketName: "test-bucket",
			setupMocks: func(mockTenant *sgTesting.MockTenantService, mockBucket *sgTesting.MockBucketService) {
				// Mock tenant creation returning empty ID
				mockTenant.CreateFunc = func(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error) {
					// Don't set ID - simulates invalid response
					return tenant, nil
				}
			},
			expectError:   true,
			errorContains: "no ID returned",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock services
			mockTenantService := &sgTesting.MockTenantService{}
			mockBucketService := &sgTesting.MockBucketService{}

			// Setup mocks for this test case
			tt.setupMocks(mockTenantService, mockBucketService)

			// Create service under test
			tm := NewTenantManager(mockTenantService, mockBucketService)

			// Execute the function
			err := tm.CreateTenantWithBucket(ctx, tt.tenantName, tt.bucketName)

			// Verify results
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else if tt.errorContains != "" && !contains(err.Error(), tt.errorContains) {
					t.Errorf("Expected error to contain '%s', got '%s'", tt.errorContains, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			}
		})
	}
}

func TestTenantManager_GetTenantUsageSummary(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name            string
		tenantID        string
		setupMock       func(*sgTesting.MockTenantService)
		expectedSummary *UsageSummary
		expectError     bool
	}{
		{
			name:     "successful usage retrieval",
			tenantID: "tenant-123",
			setupMock: func(mockTenant *sgTesting.MockTenantService) {
				mockTenant.GetUsageFunc = func(ctx context.Context, id string) (*models.TenantUsage, error) {
					dataBytes := int64(1024 * 1024 * 1024) // 1GB
					objectCount := int64(1000)

					return &models.TenantUsage{
						DataBytes:   &dataBytes,
						ObjectCount: &objectCount,
						Buckets: []*models.BucketStats{
							{Name: stringPtr("bucket1")},
							{Name: stringPtr("bucket2")},
						},
					}, nil
				}
			},
			expectedSummary: &UsageSummary{
				TenantID:     "tenant-123",
				TotalBytes:   1024 * 1024 * 1024,
				TotalObjects: 1000,
				BucketCount:  2,
			},
			expectError: false,
		},
		{
			name:     "usage retrieval fails",
			tenantID: "nonexistent-tenant",
			setupMock: func(mockTenant *sgTesting.MockTenantService) {
				mockTenant.GetUsageFunc = func(ctx context.Context, id string) (*models.TenantUsage, error) {
					return nil, errors.New("tenant not found")
				}
			},
			expectError: true,
		},
		{
			name:     "empty usage data",
			tenantID: "empty-tenant",
			setupMock: func(mockTenant *sgTesting.MockTenantService) {
				mockTenant.GetUsageFunc = func(ctx context.Context, id string) (*models.TenantUsage, error) {
					return &models.TenantUsage{
						Buckets: []*models.BucketStats{},
					}, nil
				}
			},
			expectedSummary: &UsageSummary{
				TenantID:     "empty-tenant",
				TotalBytes:   0,
				TotalObjects: 0,
				BucketCount:  0,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock service
			mockTenantService := &sgTesting.MockTenantService{}
			mockBucketService := &sgTesting.MockBucketService{}

			// Setup mock for this test case
			tt.setupMock(mockTenantService)

			// Create service under test
			tm := NewTenantManager(mockTenantService, mockBucketService)

			// Execute the function
			summary, err := tm.GetTenantUsageSummary(ctx, tt.tenantID)

			// Verify results
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}

				if summary == nil {
					t.Errorf("Expected summary but got nil")
					return
				}

				// Verify summary fields
				if summary.TenantID != tt.expectedSummary.TenantID {
					t.Errorf("Expected TenantID %s, got %s", tt.expectedSummary.TenantID, summary.TenantID)
				}
				if summary.TotalBytes != tt.expectedSummary.TotalBytes {
					t.Errorf("Expected TotalBytes %d, got %d", tt.expectedSummary.TotalBytes, summary.TotalBytes)
				}
				if summary.TotalObjects != tt.expectedSummary.TotalObjects {
					t.Errorf("Expected TotalObjects %d, got %d", tt.expectedSummary.TotalObjects, summary.TotalObjects)
				}
				if summary.BucketCount != tt.expectedSummary.BucketCount {
					t.Errorf("Expected BucketCount %d, got %d", tt.expectedSummary.BucketCount, summary.BucketCount)
				}
			}
		})
	}
}

// Example benchmark test
func BenchmarkTenantManager_GetTenantUsageSummary(b *testing.B) {
	ctx := context.Background()

	// Setup mock
	mockTenantService := &sgTesting.MockTenantService{}
	mockTenantService.GetUsageFunc = func(ctx context.Context, id string) (*models.TenantUsage, error) {
		dataBytes := int64(1024 * 1024 * 1024)
		objectCount := int64(1000)

		return &models.TenantUsage{
			DataBytes:   &dataBytes,
			ObjectCount: &objectCount,
			Buckets:     []*models.BucketStats{},
		}, nil
	}

	tm := NewTenantManager(mockTenantService, &sgTesting.MockBucketService{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tm.GetTenantUsageSummary(ctx, "tenant-123")
		if err != nil {
			b.Fatalf("Unexpected error: %v", err)
		}
	}
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func contains(str, substr string) bool {
	return len(str) >= len(substr) && (str == substr || (len(str) > len(substr) &&
		(str[:len(substr)] == substr || str[len(str)-len(substr):] == substr ||
			containsInMiddle(str, substr))))
}

func containsInMiddle(str, substr string) bool {
	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
