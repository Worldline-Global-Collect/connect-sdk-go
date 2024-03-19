// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CreatePaymentResponse represents class CreatePaymentResponse
type CreatePaymentResponse struct {
	CreationOutput *PaymentCreationOutput `json:"creationOutput,omitempty"`
	MerchantAction *MerchantAction        `json:"merchantAction,omitempty"`
	Payment        *Payment               `json:"payment,omitempty"`
}

// NewCreatePaymentResponse constructs a new CreatePaymentResponse
func NewCreatePaymentResponse() *CreatePaymentResponse {
	return &CreatePaymentResponse{}
}
