package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/yehlo/storagegrid-sdk-go/client"
	"github.com/yehlo/storagegrid-sdk-go/models"
)

func main() {
	// Get configuration from environment
	endpoint := os.Getenv("STORAGEGRID_ENDPOINT")
	username := os.Getenv("STORAGEGRID_USERNAME")
	password := os.Getenv("STORAGEGRID_PASSWORD")
	skipSSL := os.Getenv("STORAGEGRID_SKIP_SSL") == "true"

	if endpoint == "" || username == "" || password == "" {
		log.Fatal("Required environment variables: STORAGEGRID_ENDPOINT, STORAGEGRID_USERNAME, STORAGEGRID_PASSWORD")
	}

	ctx := context.Background()

	// Configure client options
	opts := []client.ClientOption{
		client.WithEndpoint(endpoint),
		client.WithCredentials(&models.Credentials{
			Username: username,
			Password: password,
		}),
	}

	if skipSSL {
		opts = append(opts, client.WithSkipSSL())
	}

	// Create grid client
	gridClient, err := client.NewGridClient(opts...)
	if err != nil {
		log.Fatalf("Failed to create grid client: %v", err)
	}

	fmt.Println("🏢 StorageGRID Tenant Management Example")
	fmt.Println("======================================")

	// List existing tenants
	if err := listTenants(ctx, gridClient); err != nil {
		log.Printf("Failed to list tenants: %v", err)
	}

	// Demonstrate tenant creation
	if err := createExampleTenant(ctx, gridClient); err != nil {
		log.Printf("Failed to create tenant: %v", err)
	}

	// Show tenant usage monitoring
	if err := monitorTenantUsage(ctx, gridClient); err != nil {
		log.Printf("Failed to monitor usage: %v", err)
	}
}

func listTenants(ctx context.Context, client *client.GridClient) error {
	fmt.Println("\n📋 Listing existing tenants...")

	tenants, err := client.Tenant().List(ctx)
	if err != nil {
		return fmt.Errorf("failed to list tenants: %w", err)
	}

	if len(*tenants) == 0 {
		fmt.Println("  No tenants found")
		return nil
	}

	fmt.Printf("  Found %d tenant(s):\n", len(*tenants))
	for _, tenant := range *tenants {
		fmt.Printf("    • %s (ID: %s)\n",
			stringValue(tenant.Name),
			tenant.Id)

		if tenant.Policy != nil && tenant.Policy.QuotaObjectBytes != nil {
			quota := *tenant.Policy.QuotaObjectBytes
			fmt.Printf("      Quota: %s\n", formatBytes(quota))
		}

		if len(tenant.Capabilities) > 0 {
			fmt.Printf("      Capabilities: %s\n", strings.Join(tenant.Capabilities, ", "))
		}
	}

	return nil
}

func createExampleTenant(ctx context.Context, client *client.GridClient) error {
	fmt.Println("\n🏗️  Creating example tenant...")

	// Check if example tenant already exists
	tenants, err := client.Tenant().List(ctx)
	if err != nil {
		return fmt.Errorf("failed to check existing tenants: %w", err)
	}

	for _, tenant := range *tenants {
		if stringValue(tenant.Name) == "example-tenant" {
			fmt.Println("  Example tenant already exists, skipping creation")
			return nil
		}
	}

	// Create new tenant
	newTenant := &models.Tenant{
		Name:         stringPtr("example-tenant"),
		Description:  stringPtr("Example tenant created by SDK"),
		Capabilities: []string{"management", "s3"},
		Policy: &models.TenantPolicy{
			UseAccountIdentitySource: false,
			AllowPlatformServices:    true,
			AllowSelectObjectContent: boolPtr(true),
			QuotaObjectBytes:         int64Ptr(10 * 1024 * 1024 * 1024), // 10GB
		},
	}

	createdTenant, err := client.Tenant().Create(ctx, newTenant)
	if err != nil {
		return fmt.Errorf("failed to create tenant: %w", err)
	}

	fmt.Printf("  ✅ Created tenant: %s\n", stringValue(createdTenant.Name))
	fmt.Printf("     Account ID: %s\n", createdTenant.Id)
	fmt.Printf("     Capabilities: %s\n", strings.Join(createdTenant.Capabilities, ", "))

	if createdTenant.Policy != nil && createdTenant.Policy.QuotaObjectBytes != nil {
		quota := *createdTenant.Policy.QuotaObjectBytes
		fmt.Printf("     Quota: %s\n", formatBytes(quota))
	}

	return nil
}

func monitorTenantUsage(ctx context.Context, client *client.GridClient) error {
	fmt.Println("\n📊 Monitoring tenant usage...")

	tenants, err := client.Tenant().List(ctx)
	if err != nil {
		return fmt.Errorf("failed to list tenants: %w", err)
	}

	if len(*tenants) == 0 {
		fmt.Println("  No tenants to monitor")
		return nil
	}

	for _, tenant := range *tenants {
		fmt.Printf("\n  📈 Usage for %s:\n", stringValue(tenant.Name))

		usage, err := client.Tenant().GetUsage(ctx, tenant.Id)
		if err != nil {
			fmt.Printf("    ❌ Failed to get usage: %v\n", err)
			continue
		}

		// Display overall usage
		if usage.ObjectCount != nil {
			fmt.Printf("    Objects: %s\n", formatNumber(*usage.ObjectCount))
		}
		if usage.DataBytes != nil {
			fmt.Printf("    Data: %s\n", formatBytes(*usage.DataBytes))
		}
		if usage.CalculationTime != nil {
			fmt.Printf("    Last Updated: %s\n", usage.CalculationTime.Format("2006-01-02 15:04:05"))
		}

		// Display bucket-level usage
		if len(usage.Buckets) > 0 {
			fmt.Printf("    Buckets:\n")
			for _, bucket := range usage.Buckets {
				if bucket.Name != nil {
					fmt.Printf("      • %s: ", *bucket.Name)
					if bucket.ObjectCount != nil && bucket.DataBytes != nil {
						fmt.Printf("%d objects, %s\n",
							*bucket.ObjectCount,
							formatBytes(*bucket.DataBytes))
					} else {
						fmt.Printf("no data\n")
					}
				}
			}
		}
	}

	return nil
}

// Helper functions
func stringPtr(s string) *string { return &s }
func boolPtr(b bool) *bool       { return &b }
func int64Ptr(i int64) *int64    { return &i }

func stringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

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
	str := strconv.FormatInt(n, 10)
	result := ""
	for i, char := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result += ","
		}
		result += string(char)
	}
	return result
}
