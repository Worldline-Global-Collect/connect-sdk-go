// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CashPaymentProduct1503SpecificInput represents class CashPaymentProduct1503SpecificInput
//
// Deprecated: No replacement
type CashPaymentProduct1503SpecificInput struct {
	// Deprecated: No replacement, since Boleto Bancario no longer needs a return URL
	ReturnURL *string `json:"returnUrl,omitempty"`
}

// NewCashPaymentProduct1503SpecificInput constructs a new CashPaymentProduct1503SpecificInput instance
func NewCashPaymentProduct1503SpecificInput() *CashPaymentProduct1503SpecificInput {
	return &CashPaymentProduct1503SpecificInput{}
}
