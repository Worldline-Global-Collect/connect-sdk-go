// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// Dispute represents class Dispute
type Dispute struct {
	DisputeOutput *DisputeOutput       `json:"disputeOutput,omitempty"`
	ID            *string              `json:"id,omitempty"`
	PaymentID     *string              `json:"paymentId,omitempty"`
	Status        *string              `json:"status,omitempty"`
	StatusOutput  *DisputeStatusOutput `json:"statusOutput,omitempty"`
}

// NewDispute constructs a new Dispute
func NewDispute() *Dispute {
	return &Dispute{}
}
