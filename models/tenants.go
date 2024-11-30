package models

import (
	"time"
)

type Tenant struct {
	// the descriptive name specified for the account (This name is for display only and might not be unique.)
	Name *string `json:"name,omitempty"`
	// Additional identifying information for the tenant account, such as an email address. This information is not shown in the Tenant Manager.
	Description *string `json:"description,omitempty"`
	// the high-level features enabled for this account, such as S3 or Swift protocols (Accounts must have the \"management\" capability if users will sign into the Tenant Manager.)
	Capabilities     []string                     `json:"capabilities,omitempty"`
	SynchronizeRules *AccountNoIdSynchronizeRules `json:"synchronizeRules,omitempty"`
	Policy           *PolicyAccount               `json:"policy,omitempty"`
	// a unique identifier for the account (automatically assigned when an account is created)
	Id string `json:"id"`
	// Automatically assigned when generating the response. Ignored in the PUT body. Present only if this tenant account has permission to use a grid federation connection. If true, this account on the local grid is a replica of an account created on another grid. If false, this account was created on the local grid and is not a copy.
	AccountReplica *bool `json:"accountReplica,omitempty"`
	// the root password for the account. This field is not returned in the response.
	Password *string `json:"password,omitempty"`
}

// AccountNoIdSynchronizeRules Rules that specify which tenant data and operations will be cloned to the other grids in a grid federation connection.
type AccountNoIdSynchronizeRules struct {
	// If true, also create user on the other grid using the grid federation connection.
	CreateUser *bool `json:"createUser,omitempty"`
	// If true, also create group on the other grid using the grid federation connection.
	CreateGroup *bool `json:"createGroup,omitempty"`
	// If true, also create key on the other grid using the grid federation connection.
	CreateKey *bool `json:"createKey,omitempty"`
}

// PolicyAccount settings for the tenant account
type PolicyAccount struct {
	// whether the tenant account should configure its own identity source. If false, the tenant uses the grid-wide identity source.
	UseAccountIdentitySource bool `json:"useAccountIdentitySource"`
	// whether a tenant can use platform services features such as CloudMirror. These features send data to an external service that is specified using a StorageGRID endpoint.
	AllowPlatformServices bool `json:"allowPlatformServices"`
	// whether a tenant can use the S3 SelectObjectContent API to filter and retrieve object data.
	AllowSelectObjectContent *bool `json:"allowSelectObjectContent,omitempty"`
	// Connection IDs of tenants. When specified, cross-grid replication of this account and the buckets in this account will be allowed.
	AllowedGridFederationConnections []string `json:"allowedGridFederationConnections,omitempty"`
	// the maximum number of bytes available for this tenant's objects. Represents a logical amount (object size), not a physical amount (size on disk). If null, an unlimited number of bytes is available.
	QuotaObjectBytes *int64 `json:"quotaObjectBytes,omitempty"`
}

type ContainerCreate struct {
	// if true, object versioning will be enabled for the bucket.
	EnableVersioning *bool `json:"enableVersioning,omitempty"`
	// Bucket name. Must be unique across the grid and DNS compliant. See the instructions for using S3 for details.
	Name string `json:"name"`
	// the region for this bucket, which must already be defined (defaults to us-east-1 if not specified)
	Region       *string                        `json:"region,omitempty"`
	S3ObjectLock *ContainerS3ObjectLockSettings `json:"s3ObjectLock,omitempty"`
}

type S3AccessKeyWithSecrets struct {
	// A unique identifier for the S3 credential pair (automatically assigned when an access key is created)
	Id string `json:"id"`
	// Storage Tenant Account ID
	AccountId string `json:"accountId"`
	// Obfuscated access key
	DisplayName string `json:"displayName"`
	// Contains the user name and account ID (generated automatically)
	UserURN *string `json:"userURN,omitempty"`
	// ID that uniquely identifies the user (generated automatically)
	UserUUID string `json:"userUUID"`
	// The time after which the key pair will no longer be valid. Null means the key pair never expires.
	Expires *time.Time `json:"expires,omitempty"`
	// generated automatically (returned only when generated and otherwise omitted)
	AccessKey *string `json:"accessKey,omitempty"`
	// generated automatically (returned only when generated and otherwise omitted)
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`
}

type ContainerS3ObjectLockSettings struct {
	// whether the bucket has S3 Object Lock enabled
	Enabled                 bool                                           `json:"enabled"`
	DefaultRetentionSetting *ContainerS3ObjectLockDefaultRetentionSettings `json:"defaultRetentionSetting,omitempty"`
}

// ContainerS3ObjectLockDefaultRetentionSettings Default retention settings for S3 Object Lock.
type ContainerS3ObjectLockDefaultRetentionSettings struct {
	// The retention mode used for new objects added to this bucket. Must be compliance, which means that an object version cannot be overwritten or deleted by any user.
	Mode string `json:"mode"`
	// The length of the default retention period for new objects added to this bucket, in days. If provided, must be paired with retentionMode. Does not affect existing bucket objects or objects with their own retain-until-date settings.
	Days *int32 `json:"days,omitempty"`
	// The length of the default retention period for new objects added to this bucket, in years. If provided, must be paired with retentionMode. Does not affect existing bucket objects or objects with their own retain-until-date settings.
	Years *int32 `json:"years,omitempty"`
}
