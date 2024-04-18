// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// SepaDirectDebitPaymentMethodSpecificInputBase represents class SepaDirectDebitPaymentMethodSpecificInputBase
type SepaDirectDebitPaymentMethodSpecificInputBase struct {
	PaymentProduct771SpecificInput *SepaDirectDebitPaymentProduct771SpecificInputBase `json:"paymentProduct771SpecificInput,omitempty"`
	PaymentProductID               *int32                                             `json:"paymentProductId,omitempty"`
}

// NewSepaDirectDebitPaymentMethodSpecificInputBase constructs a new SepaDirectDebitPaymentMethodSpecificInputBase instance
func NewSepaDirectDebitPaymentMethodSpecificInputBase() *SepaDirectDebitPaymentMethodSpecificInputBase {
	return &SepaDirectDebitPaymentMethodSpecificInputBase{}
}
