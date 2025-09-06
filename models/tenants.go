package models

type Tenant struct {
	// the descriptive name specified for the account (This name is for display only and might not be unique.)
	Name *string `json:"name,omitempty"`
	// Additional identifying information for the tenant account, such as an email address. This information is not shown in the Tenant Manager.
	Description *string `json:"description,omitempty"`
	// the high-level features enabled for this account, such as S3 or Swift protocols (Accounts must have the \"management\" capability if users will sign into the Tenant Manager.)
	Capabilities     []string                     `json:"capabilities,omitempty"`
	SynchronizeRules *AccountNoIdSynchronizeRules `json:"synchronizeRules,omitempty"`
	Policy           *TenantPolicy                `json:"policy,omitempty"`
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

type TenantPolicy struct {
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
