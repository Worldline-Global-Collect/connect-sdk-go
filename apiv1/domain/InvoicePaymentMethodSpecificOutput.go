// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// InvoicePaymentMethodSpecificOutput represents class InvoicePaymentMethodSpecificOutput
type InvoicePaymentMethodSpecificOutput struct {
	FraudResults     *FraudResults `json:"fraudResults,omitempty"`
	PaymentProductID *int32        `json:"paymentProductId,omitempty"`
}

// NewInvoicePaymentMethodSpecificOutput constructs a new InvoicePaymentMethodSpecificOutput
func NewInvoicePaymentMethodSpecificOutput() *InvoicePaymentMethodSpecificOutput {
	return &InvoicePaymentMethodSpecificOutput{}
}
