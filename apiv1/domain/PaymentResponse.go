// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentResponse represents class PaymentResponse
type PaymentResponse struct {
	HostedCheckoutSpecificOutput *HostedCheckoutSpecificOutput `json:"hostedCheckoutSpecificOutput,omitempty"`
	ID                           *string                       `json:"id,omitempty"`
	PaymentOutput                *PaymentOutput                `json:"paymentOutput,omitempty"`
	Status                       *string                       `json:"status,omitempty"`
	StatusOutput                 *PaymentStatusOutput          `json:"statusOutput,omitempty"`
}

// NewPaymentResponse constructs a new PaymentResponse
func NewPaymentResponse() *PaymentResponse {
	return &PaymentResponse{}
}
