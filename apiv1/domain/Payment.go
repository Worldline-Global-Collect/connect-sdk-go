// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// Payment represents class Payment
type Payment struct {
	HostedCheckoutSpecificOutput *HostedCheckoutSpecificOutput `json:"hostedCheckoutSpecificOutput,omitempty"`
	ID                           *string                       `json:"id,omitempty"`
	PaymentOutput                *PaymentOutput                `json:"paymentOutput,omitempty"`
	Status                       *string                       `json:"status,omitempty"`
	StatusOutput                 *PaymentStatusOutput          `json:"statusOutput,omitempty"`
}

// NewPayment constructs a new Payment
func NewPayment() *Payment {
	return &Payment{}
}
