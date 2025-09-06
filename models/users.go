package models

import "strings"

type User struct {
	// the machine-readable name for the User (unique within an Account; must begin with user/ or federated-user/). The portion after the slash is the \"username\" that is used to sign in
	UniqueName string `json:"uniqueName"`
	// the human-readable name for the User (required for local Users and imported automatically for federated Users)
	FullName *string `json:"fullName,omitempty"`
	// Group memberships for this User (required for local Users and imported automatically for federated Users)
	MemberOf []string `json:"memberOf,omitempty"`
	// if true, the local User cannot sign in (does not apply to federated Users)
	Disable *bool `json:"disable,omitempty"`
	// Storage Tenant Account ID, or zero for Grid Administrators
	AccountId *string `json:"accountId,omitempty"`
	// UUID for the User (generated automatically)
	Id *string `json:"id,omitempty"`
	// true if the User is federated, for example, an LDAP User
	Federated *bool `json:"federated,omitempty"`
	// contains the User uniqueName and Account ID (generated automatically)
	UserURN *string `json:"userURN,omitempty"`
}

func (u *User) GetShortname() string {
	if u == nil {
		return ""
	}

	// split the unique name by the slash
	parts := strings.Split(u.UniqueName, "/")
	// return the last part
	return parts[len(parts)-1]
}
