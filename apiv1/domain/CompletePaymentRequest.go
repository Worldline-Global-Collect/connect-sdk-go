// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CompletePaymentRequest represents class CompletePaymentRequest
type CompletePaymentRequest struct {
	CardPaymentMethodSpecificInput *CompletePaymentCardPaymentMethodSpecificInput `json:"cardPaymentMethodSpecificInput,omitempty"`
	Merchant                       *Merchant                                      `json:"merchant,omitempty"`
	Order                          *Order                                         `json:"order,omitempty"`
}

// NewCompletePaymentRequest constructs a new CompletePaymentRequest instance
func NewCompletePaymentRequest() *CompletePaymentRequest {
	return &CompletePaymentRequest{}
}
