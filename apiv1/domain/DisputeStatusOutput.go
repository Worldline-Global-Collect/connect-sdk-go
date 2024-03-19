// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// DisputeStatusOutput represents class DisputeStatusOutput
type DisputeStatusOutput struct {
	IsCancellable            *bool   `json:"isCancellable,omitempty"`
	StatusCategory           *string `json:"statusCategory,omitempty"`
	StatusCode               *int32  `json:"statusCode,omitempty"`
	StatusCodeChangeDateTime *string `json:"statusCodeChangeDateTime,omitempty"`
}

// NewDisputeStatusOutput constructs a new DisputeStatusOutput
func NewDisputeStatusOutput() *DisputeStatusOutput {
	return &DisputeStatusOutput{}
}
