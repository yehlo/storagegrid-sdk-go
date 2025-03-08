package models

type GatewayConfig struct {
	// ID is the unique identifier of the load balancer endpoint.
	Id string `json:"id,omitempty"`
	// Name is the name of the load balancer endpoint.
	DisplayName *string `json:"displayName,omitempty"`
	// Port made availablee by the ndpoint
	Port *int `json:"port,omitempty"`

	AccountID                      *string               `json:"accountId,omitempty"`
	Secure                         *bool                 `json:"secure,omitempty"`
	EnableIPv4                     *bool                 `json:"enableIPv4,omitempty"`
	EnableIPv6                     *bool                 `json:"enableIPv6,omitempty"`
	PinTargets                     *PinTargets           `json:"pinTargets,omitempty"`
	ManagementInterfaces           *ManagementInterfaces `json:"managementInterfaces,omitempty"`
	ClosedOnUntrustedClientNetwork *bool                 `json:"closedOnUntrustedClientNetwork,omitempty"`
}

type PinTargets struct {
	HaGroups       *[]string         `json:"haGroups,omitempty"`
	NodeInterfaces *[]NodeInterfaces `json:"nodeInterfaces,omitempty"`
	NodeTypes      *[]string         `json:"nodeTypes,omitempty"`
}

type NodeInterfaces struct {
	NodeID    *string `json:"nodeId,omitempty"`
	Interface *string `json:"interface,omitempty"`
}

type ManagementInterfaces struct {
	EnableGridManager   *bool `json:"enableGridManager,omitempty"`
	EnableTenantManager *bool `json:"enableTenantManager,omitempty"`
}

type GWServerConfig struct {
	DefaultServiceType     *string            `json:"defaultServiceType,omitempty"`
	AccountRestrictionMode *string            `json:"accountRestrictionMode,omitempty"`
	AccountRestrictions    *[]string          `json:"accountRestrictions,omitempty"`
	CertSource             *string            `json:"certSource,omitempty"`
	PlaintextCertData      *PlaintextCertData `json:"plaintextCertData,omitempty"`
}

type PlaintextCertData struct {
	ServerCertificateEncoded *string   `json:"serverCertificateEncoded,omitempty"`
	CaBundleEncoded          *string   `json:"caBundleEncoded,omitempty"`
	Metadata                 *Metadata `json:"metadata,omitempty"`
}

type Metadata struct {
	ServerCertificateDetails *ServerCertificateDetails `json:"serverCertificateDetails,omitempty"`
	CaBundleDetails          *[]CaBundleDetails        `json:"caBundleDetails,omitempty"`
}

type ServerCertificateDetails struct {
	Subject      *string `json:"subject,omitempty"`
	Issuer       *string `json:"issuer,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty"`

	NotBefore       *string       `json:"notBefore,omitempty"`
	NotAfter        *string       `json:"notAfter,omitempty"`
	FingerPrints    *FingerPrints `json:"fingerPrints,omitempty"`
	SubjectAltNames *[]string     `json:"subjectAltNames,omitempty"`
	KeyUsage        *string       `json:"keyUsage,omitempty"`
}

type FingerPrints struct {
	SHA1   *string `json:"SHA-1,omitempty"`
	SHA256 *string `json:"SHA-256,omitempty"`
}

type CaBundleDetails struct {
	Subject      *string       `json:"subject,omitempty"`
	Issuer       *string       `json:"issuer,omitempty"`
	SerialNumber *string       `json:"serialNumber,omitempty"`
	NotBefore    *string       `json:"notBefore,omitempty"`
	NotAfter     *string       `json:"notAfter,omitempty"`
	FingerPrints *FingerPrints `json:"fingerPrints,omitempty"`
	KeyUsage     *string       `json:"keyUsage,omitempty"`
}
