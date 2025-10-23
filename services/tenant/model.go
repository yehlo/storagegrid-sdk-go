package tenant

type Tenant struct {
	Name             *string                      `json:"name,omitempty"`         // the descriptive name specified for the account (This name is for display only and might not be unique.)
	Description      *string                      `json:"description,omitempty"`  // Additional identifying information for the tenant account, such as an email address. This information is not shown in the Tenant Manager.
	Capabilities     []string                     `json:"capabilities,omitempty"` // the high-level features enabled for this account, such as S3 or Swift protocols (Accounts must have the "management" capability if users will sign into the Tenant Manager.)
	SynchronizeRules *AccountNoIDSynchronizeRules `json:"synchronizeRules,omitempty"`
	Policy           *TenantPolicy                `json:"policy,omitempty"`
	ID               string                       `json:"id"`                       // a unique identifier for the account (automatically assigned when an account is created)
	AccountReplica   *bool                        `json:"accountReplica,omitempty"` // Automatically assigned when generating the response. Ignored in the PUT body. Present only if this tenant account has permission to use a grid federation connection. If true, this account on the local grid is a replica of an account created on another grid. If false, this account was created on the local grid and is not a copy.
	Password         *string                      `json:"password,omitempty"`       // the root password for the account. This field is not returned in the response.
}

// AccountNoIDSynchronizeRules Rules that specify which tenant data and operations will be cloned to the other grids in a grid federation connection.
type AccountNoIDSynchronizeRules struct {
	CreateUser  *bool `json:"createUser,omitempty"`  // If true, also create user on the other grid using the grid federation connection.
	CreateGroup *bool `json:"createGroup,omitempty"` // If true, also create group on the other grid using the grid federation connection.
	CreateKey   *bool `json:"createKey,omitempty"`   // If true, also create key on the other grid using the grid federation connection.
}

type TenantPolicy struct {
	UseAccountIdentitySource         bool     `json:"useAccountIdentitySource"`                   // whether the tenant account should configure its own identity source. If false, the tenant uses the grid-wide identity source.
	AllowPlatformServices            bool     `json:"allowPlatformServices"`                      // whether a tenant can use platform services features such as CloudMirror. These features send data to an external service that is specified using a StorageGRID endpoint.
	AllowSelectObjectContent         *bool    `json:"allowSelectObjectContent,omitempty"`         // whether a tenant can use the S3 SelectObjectContent API to filter and retrieve object data.
	AllowedGridFederationConnections []string `json:"allowedGridFederationConnections,omitempty"` // Connection IDs of tenants. When specified, cross-grid replication of this account and the buckets in this account will be allowed.
	AllowComplianceMode              bool     `json:"allowComplianceMode"`                        //Whether a tenant can use compliance mode for object lock and retention. If omitted, it defaults to true if global object lock is enabled. Otherwise it defaults to false unless explicitly set to true.
	QuotaObjectBytes                 *int64   `json:"quotaObjectBytes,omitempty"`                 // the maximum number of bytes available for this tenant's objects. Represents a logical amount (object size), not a physical amount (size on disk). If null, an unlimited number of bytes is available.
	MaxRetentionDays                 *int     `json:"maxRetentionDays"`                           // The maximum retention period in days allowed for new objects in compliance or governance mode. Does not affect existing objects. If both maxRetentionDays and maxRetentionYears are omitted, the maximum retention limit will default to 100 years. If both maxRetentionDays and maxRetentionYears are null, the maximum retention limit will default to 1 year.
	MaxRetentionYears                *int     `json:"maxRetentionYears"`                          // The maximum retention period in years allowed for new objects in compliance or governance mode. Does not affect existing objects. If both maxRetentionDays and maxRetentionYears are omitted, the maximum retention limit will default to 100 years. If both maxRetentionDays and maxRetentionYears are null, the maximum retention limit will default to 1 year.
}
