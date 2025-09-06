# Testing Examples for StorageGRID SDK

This directory contains comprehensive testing examples and patterns for the StorageGRID SDK, demonstrating both unit testing with mocks and integration testing with real StorageGRID environments.

## Directory Structure

```
testing/
├── README.md                    # This file - overview of testing approaches
├── unit-tests/
│   ├── README.md               # Unit testing guide and best practices
│   └── main_test.go            # Unit test examples with mocks
└── integration-tests/
    ├── README.md               # Integration testing guide and setup
    └── main_test.go            # Integration test examples with real API calls
```

## Testing Approaches

### Unit Tests (`unit-tests/`)

**Purpose**: Test your application logic in isolation using SDK mocks.

- **Fast execution** - No network calls or external dependencies
- **Perfect for CI/CD** - Reliable and deterministic
- **Complete coverage** - Test all code paths including error scenarios
- **Mock-based** - Uses the SDK's built-in mock services

**When to use**: 
- Testing business logic that uses the SDK
- Validating error handling
- Continuous integration pipelines
- Development and debugging

### Integration Tests (`integration-tests/`)

**Purpose**: Validate that the SDK works correctly with real StorageGRID environments.

- **Real API calls** - Tests actual SDK-to-StorageGRID communication
- **Environment validation** - Ensures SDK works with your StorageGRID version
- **End-to-end scenarios** - Complete workflow testing
- **Credential-based** - Requires actual StorageGRID access

**When to use**:
- Validating SDK against your StorageGRID deployment
- Pre-production testing
- Nightly regression testing
- SDK development and validation

## Quick Start

### Running Unit Tests

Unit tests run by default and require no setup:

```bash
cd unit-tests/
go test -v
```

### Running Integration Tests

Integration tests require environment setup:

```bash
# Set up environment
export STORAGEGRID_INTEGRATION_TESTS=true
export STORAGEGRID_HOST=https://your-storagegrid.com
export STORAGEGRID_USERNAME=your-username
export STORAGEGRID_PASSWORD=your-password

# Run tests
cd integration-tests/
go test -v
```

## Testing Patterns

### Mock-Based Unit Testing

```go
// Create mock service
mockTenantService := &sgTesting.MockTenantService{}

// Configure mock behavior
mockTenantService.CreateFunc = func(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error) {
    tenant.Id = "mock-tenant-id"
    return tenant, nil
}

// Use in your service
service := NewMyService(mockTenantService)
result, err := service.CreateTenant(ctx, "test-tenant")
```

### Real API Integration Testing

```go
// Create real client
gridClient, err := client.NewGridClient(
    client.WithEndpoint(host),
    client.WithCredentials(credentials),
)

// Test real operations
tenant, err := gridClient.Tenant().Create(ctx, &models.Tenant{
    Name: &tenantName,
})
```

## Available Mock Services

The SDK provides comprehensive mocks for all service interfaces:

| Service | Mock Class | Description |
|---------|------------|-------------|
| Tenant | `MockTenantService` | Tenant lifecycle management |
| Bucket | `MockBucketService` | Bucket operations |
| Health | `MockHealthService` | System health monitoring |
| HAGroup | `MockHAGroupService` | High availability groups |
| Gateway | `MockGatewayConfigService` | Gateway configuration |
| Region | `MockRegionService` | Region management |
| S3AccessKey | `MockS3AccessKeyService` | S3 access key management |
| TenantGroup | `MockTenantGroupService` | Tenant group management |
| TenantUser | `MockTenantUserService` | Tenant user management |

## Testing Best Practices

### Unit Testing
- Use table-driven tests for comprehensive coverage
- Test both success and failure scenarios
- Mock external dependencies consistently
- Keep tests fast and deterministic
- Use meaningful test names

### Integration Testing
- Run against dedicated test environments
- Include proper cleanup logic
- Use timeouts for all operations
- Handle network failures gracefully
- Never run against production systems

### General Guidelines
- Write tests before implementing features (TDD)
- Maintain high test coverage (>80%)
- Use descriptive assertions and error messages
- Document complex test scenarios
- Keep tests simple and focused

## CI/CD Integration

### Unit Tests in CI
```yaml
- name: Run Unit Tests
  run: |
    cd examples/testing/unit-tests
    go test -v -race -coverprofile=coverage.out
    go tool cover -html=coverage.out -o coverage.html
```

### Integration Tests in CI
```yaml
- name: Run Integration Tests
  if: github.event_name == 'schedule'  # Nightly only
  env:
    STORAGEGRID_INTEGRATION_TESTS: true
    STORAGEGRID_HOST: ${{ secrets.STORAGEGRID_HOST }}
    STORAGEGRID_USERNAME: ${{ secrets.STORAGEGRID_USERNAME }}
    STORAGEGRID_PASSWORD: ${{ secrets.STORAGEGRID_PASSWORD }}
  run: |
    cd examples/testing/integration-tests
    go test -v -timeout 10m
```

## Test Environment Setup

### For Unit Tests
No special setup required - tests use mocks and run entirely in memory.

### For Integration Tests
Requires access to a StorageGRID environment:

1. **Test Environment**: Set up a dedicated StorageGRID test environment
2. **Credentials**: Create service accounts with minimal required permissions
3. **Security**: Use environment variables, never hardcode credentials
4. **Isolation**: Ensure tests don't interfere with production data

## Troubleshooting

### Common Unit Test Issues
- **Import conflicts**: Use `sgTesting` alias for SDK testing package
- **Mock configuration**: Ensure all required mock functions are set
- **Test isolation**: Reset mocks between test cases

### Common Integration Test Issues
- **Authentication failures**: Verify credentials and permissions
- **Network timeouts**: Check connectivity and increase timeouts if needed
- **Resource cleanup**: Review cleanup warnings and manual verification

## Contributing

When adding new tests:

1. **Unit tests** should cover all new service methods and error conditions
2. **Integration tests** should validate real-world usage scenarios
3. **Documentation** should be updated to reflect new testing patterns
4. **CI/CD** configurations should include new test paths

## Related Documentation

- [Main SDK Documentation](../../../README.md)
- [Unit Testing Guide](unit-tests/README.md)
- [Integration Testing Guide](integration-tests/README.md)
- [SDK Service Interfaces](../../../services/)
- [Example Applications](../../)
