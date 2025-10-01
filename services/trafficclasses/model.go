package trafficclasses

type TrafficClass struct {
	ID          string  `json:"id"`                    // a unique identifier for the traffic class (automatically assigned when a traffic class is created)
	Name        string  `json:"name"`                  // the descriptive name specified for the traffic class (This name is for display only and might not be unique.)
	Description *string `json:"description,omitempty"` // A description of the policy
}

type Policy struct {
	ID          string     `json:"id,omitempty"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Matchers    []Matchers `json:"matchers"`
	Limits      []Limits   `json:"limits"`
}

type Matchers struct {
	Type    string   `json:"type"`
	Inverse *bool    `json:"inverse,omitempty"`
	Members []string `json:"members"`
}

type Limits struct {
	Type  TrafficClassLimit `json:"type"`
	Value int               `json:"value"`
}

type TrafficClassLimit string

// Possible TrafficClassLimits
const (
	AggregateBandwidthIn    TrafficClassLimit = "aggregateBandwidthIn"
	AggregateBandwidthOut   TrafficClassLimit = "aggregateBandwidthOut"
	ConcurrentReadRequests  TrafficClassLimit = "concurrentReadRequests"
	ConcurrentWriteRequests TrafficClassLimit = "concurrentWriteRequests"
	ReadRequestRate         TrafficClassLimit = "readRequestRate"
	WriteRequestRate        TrafficClassLimit = "writeRequestRate"
	PerRequestBandwidthIn   TrafficClassLimit = "perRequestBandwidthIn"
	PerRequestBandwidthOut  TrafficClassLimit = "perRequestBandwidthOut"
)
