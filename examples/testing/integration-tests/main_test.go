package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/yehlo/storagegrid-sdk-go/client"
	"github.com/yehlo/storagegrid-sdk-go/models"
)

// Integration tests that require a real StorageGRID environment
// These tests are disabled by default and require environment variables to run

func TestMain(m *testing.M) {
	// Check if integration tests should run
	if os.Getenv("STORAGEGRID_INTEGRATION_TESTS") != "true" {
		fmt.Println("Skipping integration tests. Set STORAGEGRID_INTEGRATION_TESTS=true to enable.")
		os.Exit(0)
	}

	// Validate required environment variables
	requiredVars := []string{
		"STORAGEGRID_HOST",
		"STORAGEGRID_USERNAME",
		"STORAGEGRID_PASSWORD",
	}

	for _, env := range requiredVars {
		if os.Getenv(env) == "" {
			log.Fatalf("Required environment variable %s is not set", env)
		}
	}

	// Run tests
	code := m.Run()
	os.Exit(code)
}

func setupGridClient(t *testing.T) *client.GridClient {
	t.Helper()

	host := os.Getenv("STORAGEGRID_HOST")
	username := os.Getenv("STORAGEGRID_USERNAME")
	password := os.Getenv("STORAGEGRID_PASSWORD")

	credentials := &models.Credentials{
		Username: username,
		Password: password,
	}

	gridClient, err := client.NewGridClient(
		client.WithEndpoint(host),
		client.WithCredentials(credentials),
		client.WithSkipSSL(), // For testing environments
	)
	if err != nil {
		t.Fatalf("Failed to create grid client: %v", err)
	}

	return gridClient
}

func setupTenantClient(t *testing.T, tenantID string) *client.TenantClient {
	t.Helper()

	host := os.Getenv("STORAGEGRID_HOST")
	username := os.Getenv("STORAGEGRID_TENANT_USERNAME")
	password := os.Getenv("STORAGEGRID_TENANT_PASSWORD")

	if username == "" || password == "" {
		t.Skip("Tenant credentials not provided, skipping tenant client tests")
	}

	credentials := &models.Credentials{
		Username:  username,
		Password:  password,
		AccountId: &tenantID,
	}

	tenantClient, err := client.NewTenantClient(
		client.WithEndpoint(host),
		client.WithCredentials(credentials),
		client.WithSkipSSL(), // For testing environments
	)
	if err != nil {
		t.Fatalf("Failed to create tenant client: %v", err)
	}

	return tenantClient
}

func TestGridClient_HealthCheck(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	gridClient := setupGridClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test health check
	health, err := gridClient.Health().Get(ctx)
	if err != nil {
		t.Fatalf("Failed to get grid health: %v", err)
	}

	if health == nil {
		t.Fatal("Health response is nil")
	}

	t.Logf("Grid health status: %v", health)
}

func TestGridClient_TenantManagement(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	gridClient := setupGridClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Generate unique tenant name for this test
	tenantName := fmt.Sprintf("test-tenant-%d", time.Now().Unix())

	// Create tenant
	tenant := &models.Tenant{
		Name:         &tenantName,
		Capabilities: []string{"s3", "management"},
	}

	t.Logf("Creating tenant: %s", tenantName)
	createdTenant, err := gridClient.Tenant().Create(ctx, tenant)
	if err != nil {
		t.Fatalf("Failed to create tenant: %v", err)
	}

	if createdTenant.Id == "" {
		t.Fatal("Created tenant has no ID")
	}

	t.Logf("Created tenant with ID: %s", createdTenant.Id)

	// Cleanup: Delete the tenant
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cleanupCancel()

		err := gridClient.Tenant().Delete(cleanupCtx, createdTenant.Id)
		if err != nil {
			t.Logf("Warning: Failed to cleanup tenant %s: %v", createdTenant.Id, err)
		} else {
			t.Logf("Successfully cleaned up tenant %s", createdTenant.Id)
		}
	})

	// List tenants to verify creation
	tenants, err := gridClient.Tenant().List(ctx)
	if err != nil {
		t.Fatalf("Failed to list tenants: %v", err)
	}

	found := false
	for _, tenant := range *tenants {
		if tenant.Id == createdTenant.Id {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Created tenant not found in tenant list")
	}

	// Get tenant by ID
	retrievedTenant, err := gridClient.Tenant().GetById(ctx, createdTenant.Id)
	if err != nil {
		t.Fatalf("Failed to get tenant by ID: %v", err)
	}

	if retrievedTenant.Id != createdTenant.Id {
		t.Fatalf("Retrieved tenant ID mismatch: expected %s, got %s", createdTenant.Id, retrievedTenant.Id)
	}

	if retrievedTenant.Name == nil || *retrievedTenant.Name != tenantName {
		t.Fatalf("Retrieved tenant name mismatch: expected %s, got %v", tenantName, retrievedTenant.Name)
	}
}

func TestTenantClient_BucketOperations(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// This test requires a pre-existing tenant and credentials
	tenantID := os.Getenv("STORAGEGRID_TEST_TENANT_ID")
	if tenantID == "" {
		t.Skip("STORAGEGRID_TEST_TENANT_ID not set, skipping tenant client tests")
	}

	tenantClient := setupTenantClient(t, tenantID)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Generate unique bucket name
	bucketName := fmt.Sprintf("test-bucket-%d", time.Now().Unix())

	// Create bucket
	bucket := &models.Bucket{
		Name:   bucketName,
		Region: "us-east-1",
	}

	t.Logf("Creating bucket: %s", bucketName)
	createdBucket, err := tenantClient.Bucket().Create(ctx, bucket)
	if err != nil {
		t.Fatalf("Failed to create bucket: %v", err)
	}

	if createdBucket.Name != bucketName {
		t.Fatalf("Created bucket name mismatch: expected %s, got %s", bucketName, createdBucket.Name)
	}

	t.Logf("Created bucket: %s", createdBucket.Name)

	// Cleanup: Delete the bucket
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cleanupCancel()

		err := tenantClient.Bucket().Delete(cleanupCtx, bucketName)
		if err != nil {
			t.Logf("Warning: Failed to cleanup bucket %s: %v", bucketName, err)
		} else {
			t.Logf("Successfully cleaned up bucket %s", bucketName)
		}
	})

	// List buckets to verify creation
	buckets, err := tenantClient.Bucket().List(ctx)
	if err != nil {
		t.Fatalf("Failed to list buckets: %v", err)
	}

	found := false
	for _, bucket := range *buckets {
		if bucket.Name == bucketName {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Created bucket not found in bucket list")
	}

	// Get bucket details
	retrievedBucket, err := tenantClient.Bucket().GetByName(ctx, bucketName)
	if err != nil {
		t.Fatalf("Failed to get bucket details: %v", err)
	}

	if retrievedBucket.Name != bucketName {
		t.Fatalf("Retrieved bucket name mismatch: expected %s, got %s", bucketName, retrievedBucket.Name)
	}
}

func TestGridClient_HAGroupOperations(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	gridClient := setupGridClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// List HA groups (read-only test)
	haGroups, err := gridClient.HAGroup().List(ctx)
	if err != nil {
		t.Fatalf("Failed to list HA groups: %v", err)
	}

	t.Logf("Found %d HA groups", len(*haGroups))

	// If there are HA groups, test getting details of the first one
	if len(*haGroups) > 0 {
		haGroup := (*haGroups)[0]
		if haGroup.Id != "" {
			details, err := gridClient.HAGroup().GetById(ctx, haGroup.Id)
			if err != nil {
				t.Fatalf("Failed to get HA group details: %v", err)
			}

			if details.Id != haGroup.Id {
				t.Fatalf("HA group ID mismatch: expected %s, got %s", haGroup.Id, details.Id)
			}

			t.Logf("Successfully retrieved HA group: %s", details.Id)
		}
	}
}

// Benchmark test for health check performance
func BenchmarkGridClient_HealthCheck(b *testing.B) {
	if os.Getenv("STORAGEGRID_INTEGRATION_TESTS") != "true" {
		b.Skip("Integration tests not enabled")
	}

	host := os.Getenv("STORAGEGRID_HOST")
	username := os.Getenv("STORAGEGRID_USERNAME")
	password := os.Getenv("STORAGEGRID_PASSWORD")

	credentials := &models.Credentials{
		Username: username,
		Password: password,
	}

	gridClient, err := client.NewGridClient(
		client.WithEndpoint(host),
		client.WithCredentials(credentials),
		client.WithSkipSSL(),
	)
	if err != nil {
		b.Fatalf("Failed to create grid client: %v", err)
	}

	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := gridClient.Health().Get(ctx)
		if err != nil {
			b.Fatalf("Health check failed: %v", err)
		}
	}
}
