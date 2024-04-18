// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// SepaDirectDebitPaymentMethodSpecificOutput represents class SepaDirectDebitPaymentMethodSpecificOutput
type SepaDirectDebitPaymentMethodSpecificOutput struct {
	FraudResults                    *FraudResults                    `json:"fraudResults,omitempty"`
	PaymentProduct771SpecificOutput *PaymentProduct771SpecificOutput `json:"paymentProduct771SpecificOutput,omitempty"`
	PaymentProductID                *int32                           `json:"paymentProductId,omitempty"`
	Token                           *string                          `json:"token,omitempty"`
}

// NewSepaDirectDebitPaymentMethodSpecificOutput constructs a new SepaDirectDebitPaymentMethodSpecificOutput instance
func NewSepaDirectDebitPaymentMethodSpecificOutput() *SepaDirectDebitPaymentMethodSpecificOutput {
	return &SepaDirectDebitPaymentMethodSpecificOutput{}
}
