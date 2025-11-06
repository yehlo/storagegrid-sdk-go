# Integration Testing with StorageGRID SDK

This example demonstrates how to write integration tests that interact with a real StorageGRID environment using the StorageGRID SDK.

## Overview

Integration tests validate that the SDK works correctly with an actual StorageGRID deployment. These tests require access to a StorageGRID environment and appropriate credentials.

## Prerequisites

- Access to a StorageGRID deployment
- Grid administrator credentials for system-level operations
- Tenant administrator credentials for tenant-specific operations (optional)

## Environment Variables

The integration tests require the following environment variables to be set:

### Required for All Tests
```bash
export STORAGEGRID_INTEGRATION_TESTS=true
export STORAGEGRID_HOST=https://your-storagegrid-host.com
export STORAGEGRID_USERNAME=grid-admin-username
export STORAGEGRID_PASSWORD=grid-admin-password
```

### Optional for Tenant Tests
```bash
export STORAGEGRID_TENANT_USERNAME=tenant-admin-username
export STORAGEGRID_TENANT_PASSWORD=tenant-admin-password
export STORAGEGRID_TEST_TENANT=test-tenant-name
```

## Running Integration Tests

### Enable Integration Tests
By default, integration tests are disabled. Set the environment variable to enable them:
```bash
export STORAGEGRID_INTEGRATION_TESTS=true
```

### Run All Integration Tests
```bash
go test -v
```

### Run Specific Tests
```bash
# Run only health check tests
go test -run TestGridClient_HealthCheck -v

# Run only tenant management tests
go test -run TestGridClient_TenantManagement -v

# Run only bucket operation tests
go test -run TestTenantClient_BucketOperations -v
```

### Skip Long-Running Tests
```bash
# Run only quick tests
go test -short -v
```

### Run Benchmarks
```bash
# Run all benchmarks
go test -bench=. -v

# Run specific benchmark
go test -bench=BenchmarkGridClient_HealthCheck -v
```

## Test Categories

### Grid Client Tests

#### Health Check Tests
- **Purpose**: Validate grid health monitoring
- **Requirements**: Grid admin credentials
- **Operations**: Get grid health status

#### Tenant Management Tests
- **Purpose**: Test tenant lifecycle operations
- **Requirements**: Grid admin credentials
- **Operations**: Create, list, retrieve, delete tenants
- **Cleanup**: Automatically removes created tenants

#### HA Group Tests
- **Purpose**: Test high availability group operations
- **Requirements**: Grid admin credentials
- **Operations**: List and retrieve HA group details (read-only)

### Tenant Client Tests

#### Bucket Operations Tests
- **Purpose**: Test bucket lifecycle operations
- **Requirements**: Tenant admin credentials, existing tenant
- **Operations**: Create, list, retrieve, delete buckets
- **Cleanup**: Automatically removes created buckets

## Test Features

### Automatic Cleanup
All integration tests that create resources include automatic cleanup using `t.Cleanup()` to ensure test environments remain clean.

### Timeout Handling
All operations use context with timeouts to prevent tests from hanging indefinitely.

### Conditional Execution
Tests can be conditionally skipped based on:
- Environment variable settings
- Available credentials
- Test mode (short vs full)

### Error Validation
Tests validate both successful operations and proper error handling.

## Best Practices

### Security
- Use dedicated test credentials with minimal required permissions
- Use test environments, never production systems
- Store credentials securely (environment variables, not hardcoded)

### Reliability
- Always include cleanup logic for created resources
- Use reasonable timeouts for operations
- Handle network failures gracefully

### Maintainability
- Use descriptive test names that explain what is being tested
- Include detailed logging for debugging failed tests
- Group related tests logically

## CI/CD Integration

### GitHub Actions Example
```yaml
name: Integration Tests

on:
  schedule:
    - cron: '0 2 * * *'  # Run nightly
  workflow_dispatch:     # Allow manual trigger

jobs:
  integration:
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.25
    
    - name: Run Integration Tests
      env:
        STORAGEGRID_INTEGRATION_TESTS: true
        STORAGEGRID_HOST: ${{ secrets.STORAGEGRID_HOST }}
        STORAGEGRID_USERNAME: ${{ secrets.STORAGEGRID_USERNAME }}
        STORAGEGRID_PASSWORD: ${{ secrets.STORAGEGRID_PASSWORD }}
      run: |
        cd examples/testing/integration-tests
        go test -v -timeout 10m
```

### Jenkins Pipeline Example
```groovy
pipeline {
    agent any
    
    triggers {
        cron('H 2 * * *')  // Run nightly
    }
    
    environment {
        STORAGEGRID_INTEGRATION_TESTS = 'true'
        STORAGEGRID_HOST = credentials('storagegrid-host')
        STORAGEGRID_USERNAME = credentials('storagegrid-username')
        STORAGEGRID_PASSWORD = credentials('storagegrid-password')
    }
    
    stages {
        stage('Integration Tests') {
            steps {
                dir('examples/testing/integration-tests') {
                    sh 'go test -v -timeout 10m'
                }
            }
        }
    }
    
    post {
        always {
            // Archive test results
            publishTestResults testResultsPattern: '**/*test*.xml'
        }
    }
}
```

## Troubleshooting

### Common Issues

1. **Tests Skip Automatically**
   - Check that `STORAGEGRID_INTEGRATION_TESTS=true` is set
   - Verify all required environment variables are set

2. **Authentication Failures**
   - Verify credentials are correct
   - Check that the user has required permissions
   - Ensure the StorageGRID host is accessible

3. **Network Timeouts**
   - Check network connectivity to StorageGRID
   - Increase timeout values if needed
   - Verify firewall settings

4. **Resource Cleanup Failures**
   - Check cleanup warnings in test output
   - Manually verify test resources were removed
   - Review permissions for delete operations

### Debug Mode
Enable verbose logging by running tests with `-v` flag:
```bash
go test -v
```

This will show detailed test progress and any cleanup warnings.
