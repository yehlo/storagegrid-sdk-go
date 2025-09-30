package models

type TrafficClass struct {
	ID          string  `json:"id"`                    // a unique identifier for the traffic class (automatically assigned when a traffic clas is created)
	Name        string  `json:"name"`                  // the descriptive name specified for the traffic class (This name is for display only and might not be unique.)
	Description *string `json:"description,omitempty"` // A description of the policy
}
