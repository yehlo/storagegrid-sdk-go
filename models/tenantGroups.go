package models

type TenantGroup struct {
	// the machine-readable name for the Group (unique within an Account; must begin with group/ or federated-group/)
	UniqueName string `json:"uniqueName,omitempty"`
	// the human-readable name for the Group (required for local Groups and imported automatically for federated Groups)
	DisplayName string `json:"displayName,omitempty"`
	// Whether the group is read-only. Users can view settings and features but cannot make changes or perform operations. Local users can change their passwords.
	ManagementReadOnly *bool                `json:"managementReadOnly,omitempty"`
	Policies           *TenantGroupPolicies `json:"policies,omitempty"`
	// Storage Tenant Account ID, or zero for Grid Administrators
	AccountId *string `json:"accountId,omitempty"`
	// UUID for the Group (generated automatically)
	Id *string `json:"id,omitempty"`
	// true if the Group is federated, for example, an LDAP Group
	Federated *bool `json:"federated,omitempty"`
	// contains the Group uniqueName and Account ID (generated automatically)
	GroupURN *string `json:"groupURN,omitempty"`
}

type TenantGroupPolicies struct {
	// Management-level permissions for the group.
	Management *TenantGroupManagementPolicy `json:"management,omitempty"`
	// S3-specific permissions and policies.
	S3 *S3Policy `json:"s3,omitempty"`
	// Swift-specific roles for the group.
	Swift *SwiftPolicy `json:"swift,omitempty"`
}

type TenantGroupManagementPolicy struct {
	// Permission to manage all containers.
	ManageAllContainers *bool `json:"manageAllContainers,omitempty"`
	// Permission to manage endpoints.
	ManageEndpoints *bool `json:"manageEndpoints,omitempty"`
	// Permission to manage their own S3 credentials.
	ManageOwnS3Credentials *bool `json:"manageOwnS3Credentials,omitempty"`
	// Permission to manage their own container objects.
	ManageOwnContainerObjects *bool `json:"manageOwnContainerObjects,omitempty"`
	// Permission to view all containers.
	ViewAllContainers *bool `json:"viewAllContainers,omitempty"`
	// Root-level access permission.
	RootAccess *bool `json:"rootAccess,omitempty"`
}

// S3Policy represents the S3-specific access control policies.
type S3Policy struct {
	// The ID of the S3 policy.
	ID *string `json:"Id,omitempty"`
	// The version of the S3 policy schema.
	Version *string `json:"Version,omitempty"`
	// A list of statements defining access controls.
	Statement []S3Statement `json:"Statement,omitempty"`
}

// S3Statement represents an individual access control statement in the S3 policy.
type S3Statement struct {
	// Statement ID for reference.
	Sid string `json:"Sid,omitempty"`
	// The effect of the statement (e.g., "Allow" or "Deny").
	Effect string `json:"Effect,omitempty"`
	// Actions allowed by this statement.
	Action *[]string `json:"Action,omitempty"`
	// Actions explicitly denied by this statement.
	NotAction *[]string `json:"NotAction,omitempty"`
	// Resources this statement applies to.
	Resource []string `json:"Resource,omitempty"`
	// Resources explicitly excluded.
	NotResource []string `json:"NotResource,omitempty"`
	// Conditions under which the statement applies.
	Condition *map[string]map[string]string `json:"Condition,omitempty"`
}

// SwiftPolicy represents the roles assigned to the group for Swift operations.
type SwiftPolicy struct {
	// A list of roles (e.g., "admin").
	Roles []string `json:"roles,omitempty"`
}
