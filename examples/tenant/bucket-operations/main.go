package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yehlo/storagegrid-sdk-go/client"
	"github.com/yehlo/storagegrid-sdk-go/models"
)

func main() {
	// Get configuration from environment
	endpoint := os.Getenv("STORAGEGRID_ENDPOINT")
	accountID := os.Getenv("STORAGEGRID_ACCOUNT_ID")
	username := os.Getenv("STORAGEGRID_TENANT_USERNAME")
	password := os.Getenv("STORAGEGRID_TENANT_PASSWORD")
	skipSSL := os.Getenv("STORAGEGRID_SKIP_SSL") == "true"

	if endpoint == "" || accountID == "" || username == "" || password == "" {
		log.Fatal("Required environment variables: STORAGEGRID_ENDPOINT, STORAGEGRID_ACCOUNT_ID, STORAGEGRID_TENANT_USERNAME, STORAGEGRID_TENANT_PASSWORD")
	}

	ctx := context.Background()

	// Configure client options
	opts := []client.ClientOption{
		client.WithEndpoint(endpoint),
		client.WithCredentials(&models.Credentials{
			Username:  username,
			Password:  password,
			AccountId: &accountID,
		}),
	}

	if skipSSL {
		opts = append(opts, client.WithSkipSSL())
	}

	// Create tenant client
	tenantClient, err := client.NewTenantClient(opts...)
	if err != nil {
		log.Fatalf("Failed to create tenant client: %v", err)
	}

	fmt.Println("🪣 StorageGRID Bucket Operations Example")
	fmt.Println("=======================================")

	// List existing buckets
	if err := listBuckets(ctx, tenantClient); err != nil {
		log.Printf("Failed to list buckets: %v", err)
	}

	// Create example buckets
	if err := createExampleBuckets(ctx, tenantClient); err != nil {
		log.Printf("Failed to create buckets: %v", err)
	}

	// Monitor bucket usage
	if err := monitorBucketUsage(ctx, tenantClient); err != nil {
		log.Printf("Failed to monitor usage: %v", err)
	}

	// Demonstrate bucket management
	if err := manageBucketLifecycle(ctx, tenantClient); err != nil {
		log.Printf("Failed to manage bucket lifecycle: %v", err)
	}
}

func listBuckets(ctx context.Context, client *client.TenantClient) error {
	fmt.Println("\n📋 Listing existing buckets...")

	buckets, err := client.Bucket().List(ctx)
	if err != nil {
		return fmt.Errorf("failed to list buckets: %w", err)
	}

	if len(*buckets) == 0 {
		fmt.Println("  No buckets found")
		return nil
	}

	fmt.Printf("  Found %d bucket(s):\n", len(*buckets))
	for _, bucket := range *buckets {
		fmt.Printf("    • %s\n", bucket.Name)
		fmt.Printf("      Region: %s\n", bucket.Region)
		fmt.Printf("      Created: %s\n", bucket.CreationTime.Format("2006-01-02 15:04:05"))

		if bucket.EnableVersioning != nil {
			fmt.Printf("      Versioning: %v\n", *bucket.EnableVersioning)
		}

		if bucket.S3ObjectLock != nil && bucket.S3ObjectLock.Enabled != nil {
			fmt.Printf("      Object Lock: %v\n", *bucket.S3ObjectLock.Enabled)
		}
		fmt.Println()
	}

	return nil
}

func createExampleBuckets(ctx context.Context, client *client.TenantClient) error {
	fmt.Println("🏗️  Creating example buckets...")

	// Define bucket configurations
	bucketConfigs := []struct {
		name        string
		versioning  bool
		objectLock  bool
		description string
	}{
		{
			name:        "app-data-" + time.Now().Format("20060102"),
			versioning:  false,
			objectLock:  false,
			description: "Application data storage",
		},
		{
			name:        "app-backups-" + time.Now().Format("20060102"),
			versioning:  true,
			objectLock:  false,
			description: "Application backups with versioning",
		},
		{
			name:        "compliance-data-" + time.Now().Format("20060102"),
			versioning:  true,
			objectLock:  true,
			description: "Compliance data with object lock",
		},
	}

	for _, config := range bucketConfigs {
		// Check if bucket already exists
		existing, err := client.Bucket().GetByName(ctx, config.name)
		if err == nil && existing != nil {
			fmt.Printf("  Bucket %s already exists, skipping\n", config.name)
			continue
		}

		// Create bucket
		bucket := &models.Bucket{
			Name:             config.name,
			Region:           "us-east-1",
			EnableVersioning: &config.versioning,
		}

		// Configure S3 Object Lock if requested
		if config.objectLock {
			bucket.S3ObjectLock = &models.BucketS3ObjectLockSettings{
				Enabled: &config.objectLock,
				DefaultRetentionSetting: &models.BucketS3ObjectLockDefaultRetentionSettings{
					Mode:  "COMPLIANCE",
					Years: 1,
				},
			}
		}

		created, err := client.Bucket().Create(ctx, bucket)
		if err != nil {
			fmt.Printf("  ❌ Failed to create %s: %v\n", config.name, err)
			continue
		}

		fmt.Printf("  ✅ Created %s\n", created.Name)
		fmt.Printf("     %s\n", config.description)
		fmt.Printf("     Versioning: %v, Object Lock: %v\n",
			config.versioning, config.objectLock)
	}

	return nil
}

func monitorBucketUsage(ctx context.Context, client *client.TenantClient) error {
	fmt.Println("\n📊 Monitoring bucket usage...")

	buckets, err := client.Bucket().List(ctx)
	if err != nil {
		return fmt.Errorf("failed to list buckets: %w", err)
	}

	if len(*buckets) == 0 {
		fmt.Println("  No buckets to monitor")
		return nil
	}

	for _, bucket := range *buckets {
		fmt.Printf("\n  📈 Usage for %s:\n", bucket.Name)

		usage, err := client.Bucket().GetUsage(ctx, bucket.Name)
		if err != nil {
			fmt.Printf("    ❌ Failed to get usage: %v\n", err)
			continue
		}

		if usage.ObjectCount != nil {
			fmt.Printf("    Objects: %s\n", formatNumber(int64(*usage.ObjectCount)))
		}
		if usage.DataBytes != nil {
			fmt.Printf("    Data: %s\n", formatBytes(*usage.DataBytes))
		}
		if usage.Region != nil {
			fmt.Printf("    Region: %s\n", *usage.Region)
		}
		if usage.VersioningEnabled != nil {
			fmt.Printf("    Versioning: %v\n", *usage.VersioningEnabled)
		}
		if usage.Encryption != nil {
			fmt.Printf("    Encryption: %s\n", *usage.Encryption)
		}
	}

	return nil
}

func manageBucketLifecycle(ctx context.Context, client *client.TenantClient) error {
	fmt.Println("\n🔄 Demonstrating bucket lifecycle management...")

	// Create a temporary bucket for demonstration
	tempBucketName := "temp-demo-" + time.Now().Format("20060102-150405")

	fmt.Printf("  Creating temporary bucket: %s\n", tempBucketName)
	tempBucket := &models.Bucket{
		Name:             tempBucketName,
		Region:           "us-east-1",
		EnableVersioning: boolPtr(false),
	}

	created, err := client.Bucket().Create(ctx, tempBucket)
	if err != nil {
		return fmt.Errorf("failed to create temp bucket: %w", err)
	}

	fmt.Printf("  ✅ Created: %s\n", created.Name)

	// Wait a moment
	time.Sleep(2 * time.Second)

	// Check bucket status
	fmt.Printf("  Checking bucket status...\n")
	retrieved, err := client.Bucket().GetByName(ctx, tempBucketName)
	if err != nil {
		fmt.Printf("    ❌ Failed to retrieve bucket: %v\n", err)
	} else {
		fmt.Printf("    ✅ Bucket exists and is accessible\n")
		fmt.Printf("    Created: %s\n", retrieved.CreationTime.Format("2006-01-02 15:04:05"))
	}

	// Demonstrate bucket deletion (commented out to avoid accidental deletion)
	fmt.Printf("  \n💡 To delete the bucket, you would use:\n")
	fmt.Printf("    client.Bucket().Delete(ctx, \"%s\")\n", tempBucketName)
	fmt.Printf("  \n⚠️  Note: Bucket must be empty before deletion\n")

	// Demonstrate bucket draining (removing all objects)
	fmt.Printf("  \n💡 To remove all objects from a bucket:\n")
	fmt.Printf("    status, err := client.Bucket().Drain(ctx, \"%s\")\n", tempBucketName)
	fmt.Printf("    // Check drain status with:\n")
	fmt.Printf("    status, err := client.Bucket().DrainStatus(ctx, \"%s\")\n", tempBucketName)

	return nil
}

// Helper functions
func boolPtr(b bool) *bool { return &b }

func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func formatNumber(n int64) string {
	if n == 0 {
		return "0"
	}

	str := fmt.Sprintf("%d", n)
	result := ""
	for i, char := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result += ","
		}
		result += string(char)
	}
	return result
}
