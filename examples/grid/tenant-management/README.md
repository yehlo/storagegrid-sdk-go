# Tenant Management Example

This example demonstrates comprehensive tenant management using the StorageGRID Grid Management API.

## Purpose

Shows how to:
- List existing tenant accounts
- Create new tenants with policies and quotas
- Monitor tenant usage and bucket statistics
- Handle tenant lifecycle management

## Prerequisites

- Grid administrator credentials
- Access to StorageGRID Grid Management API

## Environment Variables

```bash
export STORAGEGRID_ENDPOINT="https://your-storagegrid.example.com"
export STORAGEGRID_USERNAME="grid-admin"
export STORAGEGRID_PASSWORD="your-password"
export STORAGEGRID_SKIP_SSL="true"  # Optional: for development environments
```

## Running the Example

```bash
cd examples/grid/tenant-management
go mod init tenant-management-example
go mod tidy
go run main.go
```

## Expected Output

```
🏢 StorageGRID Tenant Management Example
======================================

📋 Listing existing tenants...
  Found 2 tenant(s):
    • production-tenant (ID: 12345678901234567890)
      Quota: 500.0 GB
      Capabilities: management, s3
    • development-tenant (ID: 98765432109876543210)
      Quota: 100.0 GB
      Capabilities: s3

🏗️  Creating example tenant...
  ✅ Created tenant: example-tenant
     Account ID: 11111111111111111111
     Capabilities: management, s3
     Quota: 10.0 GB

📊 Monitoring tenant usage...

  📈 Usage for production-tenant:
    Objects: 1,250,000
    Data: 2.3 TB
    Last Updated: 2025-01-15 14:30:45
    Buckets:
      • app-data: 800,000 objects, 1.8 TB
      • app-logs: 450,000 objects, 512.5 GB

  📈 Usage for development-tenant:
    Objects: 15,000
    Data: 45.2 GB
    Last Updated: 2025-01-15 14:30:45
    Buckets:
      • test-bucket: 15,000 objects, 45.2 GB
```

## Features Demonstrated

### Tenant Creation
- Setting tenant name and description
- Configuring capabilities (S3, management)
- Setting storage quotas
- Enabling platform services
- Identity source configuration

### Usage Monitoring
- Overall tenant statistics
- Per-bucket usage breakdown
- Object count and data volume
- Last calculation timestamps

### Error Handling
- Checking for existing tenants
- Graceful error handling for API failures
- Informative error messages

## Tenant Policy Options

```go
Policy: &models.TenantPolicy{
    UseAccountIdentitySource: false,           // Use grid identity source
    AllowPlatformServices:    true,            // Enable CloudMirror, etc.
    AllowSelectObjectContent: boolPtr(true),   // Enable S3 Select
    QuotaObjectBytes:         int64Ptr(quota), // Storage quota in bytes
}
```

## Common Use Cases

- **Multi-tenant Setup**: Provision tenants for different departments
- **Development Environments**: Create isolated test tenants
- **Capacity Planning**: Monitor usage trends across tenants
- **Compliance**: Set quotas and monitor usage for billing

## Integration Patterns

```bash
# Automated tenant provisioning
./tenant-management --create --name "dept-${DEPARTMENT}" --quota "${QUOTA_GB}GB"

# Daily usage reporting
./tenant-management --report --format json > daily-usage.json

# Quota monitoring
if ./tenant-management --check-quotas | grep -q "EXCEEDED"; then
    send-alert "Tenant quota exceeded"
fi
```
