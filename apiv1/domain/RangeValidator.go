// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RangeValidator represents class RangeValidator
type RangeValidator struct {
	MaxValue *int32 `json:"maxValue,omitempty"`
	MinValue *int32 `json:"minValue,omitempty"`
}

// NewRangeValidator constructs a new RangeValidator instance
func NewRangeValidator() *RangeValidator {
	return &RangeValidator{}
}
