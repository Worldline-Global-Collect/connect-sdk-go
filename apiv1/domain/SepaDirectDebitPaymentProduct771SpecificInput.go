// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// SepaDirectDebitPaymentProduct771SpecificInput represents class SepaDirectDebitPaymentProduct771SpecificInput
type SepaDirectDebitPaymentProduct771SpecificInput struct {
	ExistingUniqueMandateReference *string                     `json:"existingUniqueMandateReference,omitempty"`
	Mandate                        *CreateMandateWithReturnURL `json:"mandate,omitempty"`
	// Deprecated: Use existingUniqueMandateReference or mandate.uniqueMandateReference instead
	MandateReference               *string                     `json:"mandateReference,omitempty"`
}

// NewSepaDirectDebitPaymentProduct771SpecificInput constructs a new SepaDirectDebitPaymentProduct771SpecificInput instance
func NewSepaDirectDebitPaymentProduct771SpecificInput() *SepaDirectDebitPaymentProduct771SpecificInput {
	return &SepaDirectDebitPaymentProduct771SpecificInput{}
}
