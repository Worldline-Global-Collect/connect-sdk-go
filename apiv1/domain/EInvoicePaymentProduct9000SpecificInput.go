// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// EInvoicePaymentProduct9000SpecificInput represents class EInvoicePaymentProduct9000SpecificInput
type EInvoicePaymentProduct9000SpecificInput struct {
	BankAccountIban *BankAccountIban `json:"bankAccountIban,omitempty"`
	InstallmentID   *string          `json:"installmentId,omitempty"`
}

// NewEInvoicePaymentProduct9000SpecificInput constructs a new EInvoicePaymentProduct9000SpecificInput instance
func NewEInvoicePaymentProduct9000SpecificInput() *EInvoicePaymentProduct9000SpecificInput {
	return &EInvoicePaymentProduct9000SpecificInput{}
}
