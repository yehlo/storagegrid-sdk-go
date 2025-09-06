# Grid Health Check Example

This example demonstrates how to monitor StorageGRID health status using the Grid Management API.

## Purpose

Monitors the overall health of your StorageGRID deployment, including:
- Node connectivity status
- Active alerts and alarms
- Overall operational readiness

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
cd examples/grid/health-check
go mod init health-check-example
go mod tidy
go run main.go
```

## Expected Output

```
🔍 Checking StorageGRID health...

📊 Grid Health Summary:
  Overall Status: ✅ Healthy
  All Systems Green: true
  Operationally Ready: true

🖥️  Node Status:
  Connected: 6
  Administratively Down: 0
  Unknown Status: 0

🚨 Alerts:
  Critical: 0
  Major: 0
  Minor: 0

⚠️  Alarms (Legacy):
  Critical: 0
  Major: 0
  Minor: 0
  Notice: 0

💡 Recommendations:
  ✅ Grid is healthy - no action required
```

## Health Status Interpretation

### All Green (✅)
- All nodes are connected
- No active alerts or alarms
- System is fully operational

### Operational with Issues (⚠️)
- Most nodes are connected (at most 1 disconnected)
- No major alerts
- System is functional but should be monitored

### Critical Issues (🚨)
- Multiple nodes disconnected, OR
- Major alerts present
- Immediate attention required

## Use Cases

- **Monitoring Scripts**: Integrate into monitoring systems
- **Health Dashboards**: Build operational dashboards  
- **Automated Alerts**: Trigger notifications on status changes
- **Pre-deployment Checks**: Verify grid health before operations

## Integration Example

```bash
# Use in shell scripts for monitoring
#!/bin/bash
if go run main.go | grep -q "✅ Healthy"; then
    echo "Grid is healthy"
    exit 0
else
    echo "Grid has issues - check logs"
    exit 1
fi
```
