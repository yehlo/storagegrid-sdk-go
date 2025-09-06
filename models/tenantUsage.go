package models

import "time"

// TenantUsage represents the usage statistics of a tenant.
type TenantUsage struct {
	CalculationTime *time.Time     `json:"calculationTime,omitempty"` // The time the calculation was performed.
	ObjectCount     *int64         `json:"objectCount,omitempty"`     // Total number of objects.
	DataBytes       *int64         `json:"dataBytes,omitempty"`       // Total size of data in bytes.
	Buckets         []*BucketStats `json:"buckets,omitempty"`         // List of bucket-specific statistics.
}

// BucketStats represents the statistics of a specific bucket.
type BucketStats struct {
	Name                *string `json:"name,omitempty"`                // The name of the bucket.
	ObjectCount         *int    `json:"objectCount,omitempty"`         // Number of objects in the bucket.
	DataBytes           *int64  `json:"dataBytes,omitempty"`           // Total size of data in bytes in the bucket.
	Consistency         *string `json:"consistency,omitempty"`         // Consistency status of the bucket (e.g., "available").
	Encryption          *string `json:"encryption,omitempty"`          // Encryption type used by the bucket (e.g., "AES256").
	VersioningEnabled   *bool   `json:"versioningEnabled,omitempty"`   // Indicates if versioning is enabled.
	VersioningSuspended *bool   `json:"versioningSuspended,omitempty"` // Indicates if versioning is suspended.
	Region              *string `json:"region,omitempty"`              // The region where the bucket is located (e.g., "us-east-1").
}
