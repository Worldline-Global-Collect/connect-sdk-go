// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// EInvoicePaymentMethodSpecificOutput represents class EInvoicePaymentMethodSpecificOutput
type EInvoicePaymentMethodSpecificOutput struct {
	FraudResults                     *FraudResults                             `json:"fraudResults,omitempty"`
	PaymentProduct9000SpecificOutput *EInvoicePaymentProduct9000SpecificOutput `json:"paymentProduct9000SpecificOutput,omitempty"`
	PaymentProductID                 *int32                                    `json:"paymentProductId,omitempty"`
}

// NewEInvoicePaymentMethodSpecificOutput constructs a new EInvoicePaymentMethodSpecificOutput
func NewEInvoicePaymentMethodSpecificOutput() *EInvoicePaymentMethodSpecificOutput {
	return &EInvoicePaymentMethodSpecificOutput{}
}
