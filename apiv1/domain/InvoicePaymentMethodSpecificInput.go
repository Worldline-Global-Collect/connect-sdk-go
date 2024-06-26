// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// InvoicePaymentMethodSpecificInput represents class InvoicePaymentMethodSpecificInput
type InvoicePaymentMethodSpecificInput struct {
	AdditionalReference *string `json:"additionalReference,omitempty"`
	PaymentProductID    *int32  `json:"paymentProductId,omitempty"`
}

// NewInvoicePaymentMethodSpecificInput constructs a new InvoicePaymentMethodSpecificInput instance
func NewInvoicePaymentMethodSpecificInput() *InvoicePaymentMethodSpecificInput {
	return &InvoicePaymentMethodSpecificInput{}
}
