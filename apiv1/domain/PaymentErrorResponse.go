// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentErrorResponse represents class PaymentErrorResponse
type PaymentErrorResponse struct {
	ErrorID       *string              `json:"errorId,omitempty"`
	Errors        *[]APIError          `json:"errors,omitempty"`
	PaymentResult *CreatePaymentResult `json:"paymentResult,omitempty"`
}

// NewPaymentErrorResponse constructs a new PaymentErrorResponse instance
func NewPaymentErrorResponse() *PaymentErrorResponse {
	return &PaymentErrorResponse{}
}
