// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// HostedCheckoutSpecificOutput represents class HostedCheckoutSpecificOutput
type HostedCheckoutSpecificOutput struct {
	HostedCheckoutID *string `json:"hostedCheckoutId,omitempty"`
	Variant          *string `json:"variant,omitempty"`
}

// NewHostedCheckoutSpecificOutput constructs a new HostedCheckoutSpecificOutput instance
func NewHostedCheckoutSpecificOutput() *HostedCheckoutSpecificOutput {
	return &HostedCheckoutSpecificOutput{}
}
