package models

import "time"

// S3AccessKey S3 credential pair and associated user/account information
type S3AccessKey struct {
	// A unique identifier for the S3 credential pair (automatically assigned when an access key is created)
	Id *string `json:"id,omitempty"`
	// Storage Tenant Account ID
	AccountId *string `json:"accountId,omitempty"`
	// Obfuscated access key
	DisplayName *string `json:"displayName,omitempty"`
	// Contains the user name and account ID (generated automatically)
	UserURN *string `json:"userURN,omitempty"`
	// ID that uniquely identifies the user (generated automatically)
	UserUUID *string `json:"userUUID,omitempty"`
	// The time after which the key pair will no longer be valid. Null means the key pair never expires.
	Expires *time.Time `json:"expires,omitempty"`
	// generated automatically (returned only when generated and otherwise omitted)
	AccessKey *string `json:"accessKey,omitempty"`
	// generated automatically (returned only when generated and otherwise omitted)
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`
}
