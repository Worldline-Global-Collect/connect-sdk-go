// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// EInvoicePaymentMethodSpecificInput represents class EInvoicePaymentMethodSpecificInput
type EInvoicePaymentMethodSpecificInput struct {
	AcceptedTermsAndConditions      *bool                                    `json:"acceptedTermsAndConditions,omitempty"`
	PaymentProduct9000SpecificInput *EInvoicePaymentProduct9000SpecificInput `json:"paymentProduct9000SpecificInput,omitempty"`
	PaymentProductID                *int32                                   `json:"paymentProductId,omitempty"`
	RequiresApproval                *bool                                    `json:"requiresApproval,omitempty"`
}

// NewEInvoicePaymentMethodSpecificInput constructs a new EInvoicePaymentMethodSpecificInput instance
func NewEInvoicePaymentMethodSpecificInput() *EInvoicePaymentMethodSpecificInput {
	return &EInvoicePaymentMethodSpecificInput{}
}
