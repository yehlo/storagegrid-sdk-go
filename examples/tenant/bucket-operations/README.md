# Bucket Operations Example

This example demonstrates comprehensive S3 bucket management using the StorageGRID Tenant Management API.

## Purpose

Shows how to:
- List existing buckets in a tenant
- Create buckets with different configurations
- Monitor bucket usage and statistics
- Manage bucket lifecycle (create, monitor, delete)

## Prerequisites

- Tenant user credentials with appropriate permissions
- Existing tenant account in StorageGRID
- Access to StorageGRID Tenant Management API

## Environment Variables

```bash
export STORAGEGRID_ENDPOINT="https://your-storagegrid.example.com"
export STORAGEGRID_ACCOUNT_ID="12345678901234567890"
export STORAGEGRID_TENANT_USERNAME="tenant-admin"
export STORAGEGRID_TENANT_PASSWORD="tenant-password"
export STORAGEGRID_SKIP_SSL="true"  # Optional: for development environments
```

## Running the Example

```bash
cd examples/tenant/bucket-operations
go mod init bucket-operations-example
go mod tidy
go run main.go
```

## Expected Output

```
🪣 StorageGRID Bucket Operations Example
=======================================

📋 Listing existing buckets...
  Found 2 bucket(s):
    • production-data
      Region: us-east-1
      Created: 2025-01-10 09:15:30
      Versioning: false

    • archive-storage
      Region: us-west-2
      Created: 2025-01-12 14:22:15
      Versioning: true
      Object Lock: true

🏗️  Creating example buckets...
  ✅ Created app-data-20250115
     Application data storage
     Versioning: false, Object Lock: false
  ✅ Created app-backups-20250115
     Application backups with versioning
     Versioning: true, Object Lock: false
  ✅ Created compliance-data-20250115
     Compliance data with object lock
     Versioning: true, Object Lock: true

📊 Monitoring bucket usage...

  📈 Usage for production-data:
    Objects: 150,000
    Data: 2.3 TB
    Region: us-east-1
    Versioning: false
    Encryption: AES256

  📈 Usage for archive-storage:
    Objects: 50,000
    Data: 10.5 TB
    Region: us-west-2
    Versioning: true
    Encryption: AES256

🔄 Demonstrating bucket lifecycle management...
  Creating temporary bucket: temp-demo-20250115-143025
  ✅ Created: temp-demo-20250115-143025
  Checking bucket status...
    ✅ Bucket exists and is accessible
    Created: 2025-01-15 14:30:25

  💡 To delete the bucket, you would use:
    client.Bucket().Delete(ctx, "temp-demo-20250115-143025")
  
  ⚠️  Note: Bucket must be empty before deletion

  💡 To remove all objects from a bucket:
    status, err := client.Bucket().Drain(ctx, "temp-demo-20250115-143025")
    // Check drain status with:
    status, err := client.Bucket().DrainStatus(ctx, "temp-demo-20250115-143025")
```

## Features Demonstrated

### Bucket Creation Options
- **Standard buckets**: Basic S3 storage
- **Versioned buckets**: Object versioning enabled
- **Compliance buckets**: S3 Object Lock for retention

### Configuration Options
```go
bucket := &models.Bucket{
    Name:             "my-bucket",
    Region:           "us-east-1",
    EnableVersioning: &versioning,
    S3ObjectLock: &models.BucketS3ObjectLockSettings{
        Enabled: &objectLock,
        DefaultRetentionSetting: &models.BucketS3ObjectLockDefaultRetentionSettings{
            Mode:  "COMPLIANCE",
            Years: 1,
        },
    },
}
```

### Usage Monitoring
- Object count and data volume
- Regional distribution
- Versioning status
- Encryption settings

### Lifecycle Management
- **Creation**: Set up buckets with policies
- **Monitoring**: Track usage and status
- **Draining**: Remove all objects safely
- **Deletion**: Clean up empty buckets

## Common Use Cases

### Application Data Storage
```bash
# Create buckets for different data types
./bucket-operations --create --name "app-data" --region "us-east-1"
./bucket-operations --create --name "app-logs" --region "us-east-1" --versioning
```

### Backup and Archive
```bash
# Create versioned bucket for backups
./bucket-operations --create --name "backups" --versioning --retention-days 2555
```

### Compliance Storage
```bash
# Create object-lock enabled bucket
./bucket-operations --create --name "compliance" --object-lock --retention-years 7
```

## Best Practices

1. **Naming**: Use descriptive, consistent naming conventions
2. **Regions**: Choose regions close to your applications
3. **Versioning**: Enable for critical data and backups
4. **Object Lock**: Use for compliance and regulatory requirements
5. **Monitoring**: Regular usage monitoring for capacity planning

## Error Handling

The example demonstrates:
- Checking for existing buckets before creation
- Graceful handling of API failures
- Informative error messages
- Safe operations (like commented deletion)

## Integration Patterns

```bash
# Automated bucket provisioning
for app in web api worker; do
    ./bucket-operations --create --name "${app}-data-$(date +%Y%m%d)"
done

# Usage reporting
./bucket-operations --usage --format json > bucket-usage.json

# Capacity alerts
usage=$(./bucket-operations --usage --bucket "critical-data" --field bytes)
if [ "$usage" -gt $((80 * 1024**3)) ]; then  # 80GB threshold
    alert "Bucket approaching capacity limit"
fi
```
