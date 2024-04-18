// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CreatePaymentResult represents class CreatePaymentResult
type CreatePaymentResult struct {
	CreationOutput *PaymentCreationOutput `json:"creationOutput,omitempty"`
	MerchantAction *MerchantAction        `json:"merchantAction,omitempty"`
	Payment        *Payment               `json:"payment,omitempty"`
}

// NewCreatePaymentResult constructs a new CreatePaymentResult instance
func NewCreatePaymentResult() *CreatePaymentResult {
	return &CreatePaymentResult{}
}
