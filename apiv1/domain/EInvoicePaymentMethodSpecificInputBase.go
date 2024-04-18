// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// EInvoicePaymentMethodSpecificInputBase represents class EInvoicePaymentMethodSpecificInputBase
type EInvoicePaymentMethodSpecificInputBase struct {
	PaymentProductID *int32 `json:"paymentProductId,omitempty"`
	RequiresApproval *bool  `json:"requiresApproval,omitempty"`
}

// NewEInvoicePaymentMethodSpecificInputBase constructs a new EInvoicePaymentMethodSpecificInputBase instance
func NewEInvoicePaymentMethodSpecificInputBase() *EInvoicePaymentMethodSpecificInputBase {
	return &EInvoicePaymentMethodSpecificInputBase{}
}
