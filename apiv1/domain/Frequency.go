// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// Frequency represents class Frequency
type Frequency struct {
	Interval          *string `json:"interval,omitempty"`
	IntervalFrequency *int32  `json:"intervalFrequency,omitempty"`
}

// NewFrequency constructs a new Frequency
func NewFrequency() *Frequency {
	return &Frequency{}
}
