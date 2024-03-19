// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// DisputeResponse represents class DisputeResponse
type DisputeResponse struct {
	DisputeOutput *DisputeOutput       `json:"disputeOutput,omitempty"`
	ID            *string              `json:"id,omitempty"`
	PaymentID     *string              `json:"paymentId,omitempty"`
	Status        *string              `json:"status,omitempty"`
	StatusOutput  *DisputeStatusOutput `json:"statusOutput,omitempty"`
}

// NewDisputeResponse constructs a new DisputeResponse
func NewDisputeResponse() *DisputeResponse {
	return &DisputeResponse{}
}
