// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput represents class ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput
type ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput struct {
	DateCollect *string `json:"dateCollect,omitempty"`
	Token       *string `json:"token,omitempty"`
}

// NewApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput constructs a new ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput
func NewApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput() *ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput {
	return &ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput{}
}
