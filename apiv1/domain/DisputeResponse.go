// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// DisputeResponse represents class DisputeResponse
type DisputeResponse struct {
	CaptureID     *string              `json:"captureId,omitempty"`
	DisputeOutput *DisputeOutput       `json:"disputeOutput,omitempty"`
	ID            *string              `json:"id,omitempty"`
	PaymentID     *string              `json:"paymentId,omitempty"`
	Status        *string              `json:"status,omitempty"`
	StatusOutput  *DisputeStatusOutput `json:"statusOutput,omitempty"`
}

// NewDisputeResponse constructs a new DisputeResponse instance
func NewDisputeResponse() *DisputeResponse {
	return &DisputeResponse{}
}
