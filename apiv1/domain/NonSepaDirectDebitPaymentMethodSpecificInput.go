// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// NonSepaDirectDebitPaymentMethodSpecificInput represents class NonSepaDirectDebitPaymentMethodSpecificInput
type NonSepaDirectDebitPaymentMethodSpecificInput struct {
	DateCollect                       *string                                           `json:"dateCollect,omitempty"`
	DirectDebitText                   *string                                           `json:"directDebitText,omitempty"`
	IsRecurring                       *bool                                             `json:"isRecurring,omitempty"`
	PaymentProduct705SpecificInput    *NonSepaDirectDebitPaymentProduct705SpecificInput `json:"paymentProduct705SpecificInput,omitempty"`
	PaymentProduct730SpecificInput    *NonSepaDirectDebitPaymentProduct730SpecificInput `json:"paymentProduct730SpecificInput,omitempty"`
	PaymentProductID                  *int32                                            `json:"paymentProductId,omitempty"`
	RecurringPaymentSequenceIndicator *string                                           `json:"recurringPaymentSequenceIndicator,omitempty"`
	RequiresApproval                  *bool                                             `json:"requiresApproval,omitempty"`
	Token                             *string                                           `json:"token,omitempty"`
	Tokenize                          *bool                                             `json:"tokenize,omitempty"`
}

// NewNonSepaDirectDebitPaymentMethodSpecificInput constructs a new NonSepaDirectDebitPaymentMethodSpecificInput instance
func NewNonSepaDirectDebitPaymentMethodSpecificInput() *NonSepaDirectDebitPaymentMethodSpecificInput {
	return &NonSepaDirectDebitPaymentMethodSpecificInput{}
}
