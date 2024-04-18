// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// FixedListValidator represents class FixedListValidator
type FixedListValidator struct {
	AllowedValues *[]string `json:"allowedValues,omitempty"`
}

// NewFixedListValidator constructs a new FixedListValidator instance
func NewFixedListValidator() *FixedListValidator {
	return &FixedListValidator{}
}
