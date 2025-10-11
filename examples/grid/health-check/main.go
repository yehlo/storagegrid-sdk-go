package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

	// Check grid health
	fmt.Println("🔍 Checking StorageGRID health...")

	health, err := gridClient.Health().Get(ctx)
	if err != nil {
		log.Fatalf("Failed to get health status: %v", err)
	}

	// Display overall status
	fmt.Printf("\n📊 Grid Health Summary:\n")
	fmt.Printf("  Overall Status: %s\n", getHealthStatus(health))
	fmt.Printf("  All Systems Green: %v\n", health.AllGreen())
	fmt.Printf("  Operationally Ready: %v\n", health.Operative(1))

	// Display node information
	if health.Nodes != nil {
		fmt.Printf("\n🖥️  Node Status:\n")
		if health.Nodes.Connected != nil {
			fmt.Printf("  Connected: %d\n", *health.Nodes.Connected)
		}
		if health.Nodes.AdministrativelyDown != nil {
			fmt.Printf("  Administratively Down: %d\n", *health.Nodes.AdministrativelyDown)
		}
		if health.Nodes.Unknown != nil {
			fmt.Printf("  Unknown Status: %d\n", *health.Nodes.Unknown)
		}
	}

	// Display alert information
	if health.Alerts != nil {
		fmt.Printf("\n🚨 Alerts:\n")
		if health.Alerts.Critical != nil {
			fmt.Printf("  Critical: %d\n", *health.Alerts.Critical)
		}
		if health.Alerts.Major != nil {
			fmt.Printf("  Major: %d\n", *health.Alerts.Major)
		}
		if health.Alerts.Minor != nil {
			fmt.Printf("  Minor: %d\n", *health.Alerts.Minor)
		}
	}

	// Display alarm information (legacy)
	if health.Alarms != nil {
		fmt.Printf("\n⚠️  Alarms (Legacy):\n")
		if health.Alarms.Critical != nil {
			fmt.Printf("  Critical: %d\n", *health.Alarms.Critical)
		}
		if health.Alarms.Major != nil {
			fmt.Printf("  Major: %d\n", *health.Alarms.Major)
		}
		if health.Alarms.Minor != nil {
			fmt.Printf("  Minor: %d\n", *health.Alarms.Minor)
		}
		if health.Alarms.Notice != nil {
			fmt.Printf("  Notice: %d\n", *health.Alarms.Notice)
		}
	}

	// Provide recommendations
	fmt.Printf("\n💡 Recommendations:\n")
	if health.AllGreen() {
		fmt.Printf("  ✅ Grid is healthy - no action required\n")
	} else if health.Operative(1) {
		fmt.Printf("  ⚠️  Grid is operational but has some issues - monitor closely\n")
	} else {
		fmt.Printf("  🚨 Grid has significant issues - immediate attention required\n")
	}
}

func getHealthStatus(health *models.Health) string {
	if health.AllGreen() {
		return "✅ Healthy"
	} else if health.Operative(1) {
		return "⚠️  Operational with Issues"
	} else {
		return "🚨 Critical Issues"
	}
}
