// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// SepaDirectDebitPaymentProduct771SpecificInputBase represents class SepaDirectDebitPaymentProduct771SpecificInputBase
type SepaDirectDebitPaymentProduct771SpecificInputBase struct {
	ExistingUniqueMandateReference *string            `json:"existingUniqueMandateReference,omitempty"`
	Mandate                        *CreateMandateBase `json:"mandate,omitempty"`
	// Deprecated: Use existingUniqueMandateReference or mandate.uniqueMandateReference instead
	MandateReference               *string            `json:"mandateReference,omitempty"`
}

// NewSepaDirectDebitPaymentProduct771SpecificInputBase constructs a new SepaDirectDebitPaymentProduct771SpecificInputBase instance
func NewSepaDirectDebitPaymentProduct771SpecificInputBase() *SepaDirectDebitPaymentProduct771SpecificInputBase {
	return &SepaDirectDebitPaymentProduct771SpecificInputBase{}
}
