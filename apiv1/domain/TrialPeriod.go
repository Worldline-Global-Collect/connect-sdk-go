// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// TrialPeriod represents class TrialPeriod
type TrialPeriod struct {
	Duration *int32  `json:"duration,omitempty"`
	Interval *string `json:"interval,omitempty"`
}

// NewTrialPeriod constructs a new TrialPeriod instance
func NewTrialPeriod() *TrialPeriod {
	return &TrialPeriod{}
}
