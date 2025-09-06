# StorageGRID SDK Examples

This directory contains practical examples demonstrating how to use the StorageGRID SDK for Go in various scenarios.

## Quick Start

To run any example:

```bash
cd examples/grid/health-check
go mod init health-check-example
go mod tidy
go run main.go
```

## Directory Structure

### 🔧 Grid Management (`grid/`)
Examples for grid administrators managing the entire StorageGRID system:

- **[`health-check/`](grid/health-check)** - Monitor grid health status
- **[`tenant-management/`](grid/tenant-management)** - Create and manage tenant accounts  
- **[`ha-groups/`](grid/ha-groups)** - Configure High Availability groups
- **[`gateway-config/`](grid/gateway-config)** - Manage load balancer endpoints

### 🏢 Tenant Management (`tenant/`)
Examples for tenant users managing their specific tenant resources:

- **[`bucket-operations/`](tenant/bucket-operations)** - Create and manage S3 buckets
- **[`user-management/`](tenant/user-management)** - Manage tenant users and permissions
- **[`s3-access-keys/`](tenant/s3-access-keys)** - Generate and manage S3 credentials
- **[`usage-monitoring/`](tenant/usage-monitoring)** - Monitor bucket and tenant usage

### 🧪 Testing (`testing/`)
Examples demonstrating testing patterns and mock usage:

- **[`unit-tests/`](testing/unit-tests)** - Unit testing with SDK mocks
- **[`integration-tests/`](testing/integration-tests)** - Integration testing patterns

## Environment Setup

Most examples require environment variables for authentication:

```bash
# For Grid Management examples
export STORAGEGRID_ENDPOINT="https://your-storagegrid.example.com"
export STORAGEGRID_USERNAME="grid-admin"
export STORAGEGRID_PASSWORD="your-password"

# For Tenant Management examples (additional)
export STORAGEGRID_ACCOUNT_ID="12345678901234567890"
export STORAGEGRID_TENANT_USERNAME="tenant-admin"
export STORAGEGRID_TENANT_PASSWORD="tenant-password"

# Optional: Skip SSL verification for development
export STORAGEGRID_SKIP_SSL="true"
```

## Example Categories

### Basic Operations
Start here if you're new to the SDK:
- [`grid/health-check`](grid/health-check) - Simple health monitoring
- [`tenant/bucket-operations`](tenant/bucket-operations) - Basic bucket management

### Advanced Operations  
For more complex scenarios:
- [`grid/tenant-management`](grid/tenant-management) - Complete tenant lifecycle
- [`tenant/user-management`](tenant/user-management) - User and group management

### Production Patterns
Real-world usage patterns:
- [`testing/integration-tests`](testing/integration-tests) - Production testing
- [`tenant/usage-monitoring`](tenant/usage-monitoring) - Monitoring and alerting

## Contributing Examples

We welcome new examples! When adding an example:

1. Create a new directory under the appropriate category
2. Include a `main.go` file with a complete, runnable example
3. Add a `README.md` explaining the example's purpose and usage
4. Use environment variables for configuration
5. Include proper error handling and logging
6. Update this main examples README

## Common Patterns

All examples follow these patterns:

- **Environment-based configuration** - No hardcoded credentials
- **Proper error handling** - All errors are checked and handled
- **Context usage** - All API calls use `context.Context`
- **Logging** - Clear output about what's happening
- **Documentation** - Each example explains its purpose and usage
