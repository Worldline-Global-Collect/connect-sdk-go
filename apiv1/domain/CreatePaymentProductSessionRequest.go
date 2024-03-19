// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CreatePaymentProductSessionRequest represents class CreatePaymentProductSessionRequest
type CreatePaymentProductSessionRequest struct {
	PaymentProductSession302SpecificInput *MobilePaymentProductSession302SpecificInput `json:"paymentProductSession302SpecificInput,omitempty"`
}

// NewCreatePaymentProductSessionRequest constructs a new CreatePaymentProductSessionRequest
func NewCreatePaymentProductSessionRequest() *CreatePaymentProductSessionRequest {
	return &CreatePaymentProductSessionRequest{}
}
