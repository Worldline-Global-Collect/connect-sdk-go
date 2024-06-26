// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ExemptionOutput represents class ExemptionOutput
type ExemptionOutput struct {
	ExemptionRaised          *string `json:"exemptionRaised,omitempty"`
	ExemptionRejectionReason *string `json:"exemptionRejectionReason,omitempty"`
	ExemptionRequest         *string `json:"exemptionRequest,omitempty"`
}

// NewExemptionOutput constructs a new ExemptionOutput instance
func NewExemptionOutput() *ExemptionOutput {
	return &ExemptionOutput{}
}
