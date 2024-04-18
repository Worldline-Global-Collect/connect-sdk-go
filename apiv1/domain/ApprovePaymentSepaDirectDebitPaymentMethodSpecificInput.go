// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput represents class ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput
type ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput struct {
	DateCollect *string `json:"dateCollect,omitempty"`
	Token       *string `json:"token,omitempty"`
}

// NewApprovePaymentSepaDirectDebitPaymentMethodSpecificInput constructs a new ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput instance
func NewApprovePaymentSepaDirectDebitPaymentMethodSpecificInput() *ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput {
	return &ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput{}
}
