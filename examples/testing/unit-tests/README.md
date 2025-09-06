# Unit Testing with StorageGRID SDK

This example demonstrates how to write unit tests using the mock services provided by the StorageGRID SDK.

## Overview

The StorageGRID SDK includes mock implementations of all service interfaces in the `testing` package, making it easy to write isolated unit tests for your application logic.

## Key Features

- **Complete Mock Services**: All service interfaces have corresponding mock implementations
- **Flexible Function Mocking**: Override any service method behavior for testing
- **Table-Driven Tests**: Examples of comprehensive test coverage using Go's testing patterns
- **Benchmark Tests**: Performance testing examples with mocks

## Running the Tests

```bash
# Run all tests
go test

# Run with verbose output
go test -v

# Run specific test
go test -run TestTenantManager_CreateTenantWithBucket

# Run benchmarks
go test -bench=.

# Run tests with coverage
go test -cover
```

## Example Structure

The example includes:

1. **TenantManager Service**: A service that depends on StorageGRID SDK services
2. **Comprehensive Tests**: Various test scenarios including success and failure cases
3. **Mock Setup Functions**: Reusable mock configuration for different test scenarios
4. **Benchmark Tests**: Performance testing examples

## Test Patterns

### Basic Mock Setup

```go
mockTenantService := &sgTesting.MockTenantService{}
mockTenantService.CreateFunc = func(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error) {
    tenant.Id = "tenant-123"
    return tenant, nil
}
```

### Table-Driven Tests

The example demonstrates table-driven tests with different scenarios:
- Successful operations
- Service failures
- Edge cases (empty responses, validation failures)

### Error Testing

Each test case can specify expected errors and validate error messages:

```go
{
    name: "tenant creation fails",
    setupMocks: func(mockTenant *sgTesting.MockTenantService, mockBucket *sgTesting.MockBucketService) {
        mockTenant.CreateFunc = func(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error) {
            return nil, errors.New("tenant creation failed")
        }
    },
    expectError: true,
    errorContains: "tenant creation failed",
}
```

## Best Practices

1. **Import Aliasing**: Use `sgTesting` alias to avoid conflicts with Go's standard `testing` package
2. **Mock Function Assignment**: Always assign mock functions in setup functions for clarity
3. **Test Isolation**: Each test should setup its own mocks independently
4. **Error Validation**: Test both success and failure scenarios
5. **Resource Cleanup**: Use `t.Cleanup()` for any resources that need cleanup

## Available Mock Services

The SDK provides mocks for all service interfaces:
- `MockTenantService`
- `MockBucketService`
- `MockHealthService`
- `MockHAGroupService`
- `MockGatewayConfigService`
- `MockRegionService`
- `MockS3AccessKeyService`
- `MockTenantGroupService`
- `MockTenantUserService`

## Integration with CI/CD

These unit tests are designed to run quickly and without external dependencies, making them perfect for CI/CD pipelines:

```yaml
- name: Run Unit Tests
  run: |
    go test ./examples/testing/unit-tests/... -v
    go test ./examples/testing/unit-tests/... -cover
```
