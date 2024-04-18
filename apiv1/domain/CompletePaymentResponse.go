// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CompletePaymentResponse represents class CompletePaymentResponse
type CompletePaymentResponse struct {
	CreationOutput *PaymentCreationOutput `json:"creationOutput,omitempty"`
	MerchantAction *MerchantAction        `json:"merchantAction,omitempty"`
	Payment        *Payment               `json:"payment,omitempty"`
}

// NewCompletePaymentResponse constructs a new CompletePaymentResponse instance
func NewCompletePaymentResponse() *CompletePaymentResponse {
	return &CompletePaymentResponse{}
}
