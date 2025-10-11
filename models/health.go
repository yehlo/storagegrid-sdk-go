package models

type Health struct {
	Alarms *Alarms `json:"alarms,omitempty"` // Details about alarms.
	Alerts *Alerts `json:"alerts,omitempty"` // Details about alerts.
	Nodes  *Nodes  `json:"nodes,omitempty"`  // Details about nodes.
}

// Alarms represents the counts of alarms by severity.
type Alarms struct {
	Critical *int `json:"critical,omitempty"` // Number of critical alarms.
	Major    *int `json:"major,omitempty"`    // Number of major alarms.
	Minor    *int `json:"minor,omitempty"`    // Number of minor alarms.
	Notice   *int `json:"notice,omitempty"`   // Number of notice alarms.
}

// Alerts represents the counts of alerts by severity.
type Alerts struct {
	Critical *int `json:"critical,omitempty"` // Number of critical alerts.
	Major    *int `json:"major,omitempty"`    // Number of major alerts.
	Minor    *int `json:"minor,omitempty"`    // Number of minor alerts.
}

// Nodes represents the counts of nodes by status.
type Nodes struct {
	Connected            *int `json:"connected,omitempty"`             // Number of connected nodes.
	AdministrativelyDown *int `json:"administratively-down,omitempty"` // Number of nodes that are administratively down.
	Unknown              *int `json:"unknown,omitempty"`               // Number of nodes with unknown status.
}

// NoAlarms checks if there are no alarms of any severity.
func (h *Health) NoAlarms() bool {
	if h.Alarms == nil {
		return true // No alarms present.
	}
	return isZero(h.Alarms.Critical) &&
		isZero(h.Alarms.Major) &&
		isZero(h.Alarms.Minor) &&
		isZero(h.Alarms.Notice)
}

// NoAlerts checks if there are no alerts of any severity.
func (h *Health) NoAlerts() bool {
	if h.Alerts == nil {
		return true // No alerts present.
	}
	return isZero(h.Alerts.Critical) &&
		isZero(h.Alerts.Major) &&
		isZero(h.Alerts.Minor)
}

// AllConnected checks if all nodes are connected and there are no administratively-down or unknown nodes.
func (h *Health) AllConnected() bool {
	if h.Nodes == nil {
		return true // No nodes present.
	}
	return isZero(h.Nodes.AdministrativelyDown) &&
		isZero(h.Nodes.Unknown)
}

// AllGreen returns true if all nodes are connected and no alerts and alarms are present
func (h *Health) AllGreen() bool {
	return h.AllConnected() && h.NoAlarms() && h.NoAlerts()
}

// Operative checks if the system is operational:
// 1. No Major alarms or alerts.
// 2. Node connectivity is within acceptable limits (maxUnavailable).
func (h *Health) Operative(maxUnavailable int) bool {
	// Check for no major alerts.
	if h.Alerts != nil && !isZero(h.Alerts.Major) {
		return false
	}

	// Check for node connectivity.
	if h.Nodes != nil {
		notConnected := getValue(h.Nodes.AdministrativelyDown) + getValue(h.Nodes.Unknown)
		if notConnected > maxUnavailable {
			return false
		}
	}

	return true
}

// isZero is a helper function to check if a pointer to an integer is nil or zero.
func isZero(val *int) bool {
	return val == nil || *val == 0
}

// Helper to get the value of a pointer to int, defaulting to zero if nil.
func getValue(val *int) int {
	if val == nil {
		return 0
	}
	return *val
}
