package hagroup

type HAGroup struct {
	// ID is the unique identifier of the HA group.
	ID string `json:"id,omitempty"`
	// Name is the name of the HA group.
	Name *string `json:"name,omitempty"`
	// Description is the description of the HA group.
	Description *string `json:"description,omitempty"`
	// GatewayCidr is the gateway CIDR of the HA group.
	GatewayCidr *string `json:"gatewayCidr,omitempty"`
	// VirtualIps is the virtual IPs of the HA group.
	VirtualIps []string `json:"virtualIps,omitempty"`
	// Interfaces is the interfaces of the HA group.
	Interfaces []Interface `json:"interfaces,omitempty"`
}

type Interface struct {
	// Interface is the interface of the HA group.
	Interface *string `json:"interface,omitempty"`
	// NodeID is the node ID of the HA group.
	NodeID *string `json:"nodeId,omitempty"`
}
