// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// NonSepaDirectDebitPaymentMethodSpecificOutput represents class NonSepaDirectDebitPaymentMethodSpecificOutput
type NonSepaDirectDebitPaymentMethodSpecificOutput struct {
	FraudResults     *FraudResults `json:"fraudResults,omitempty"`
	PaymentProductID *int32        `json:"paymentProductId,omitempty"`
}

// NewNonSepaDirectDebitPaymentMethodSpecificOutput constructs a new NonSepaDirectDebitPaymentMethodSpecificOutput
func NewNonSepaDirectDebitPaymentMethodSpecificOutput() *NonSepaDirectDebitPaymentMethodSpecificOutput {
	return &NonSepaDirectDebitPaymentMethodSpecificOutput{}
}
