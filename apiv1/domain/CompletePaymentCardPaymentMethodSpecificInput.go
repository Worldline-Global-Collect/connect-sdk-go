// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CompletePaymentCardPaymentMethodSpecificInput represents class CompletePaymentCardPaymentMethodSpecificInput
type CompletePaymentCardPaymentMethodSpecificInput struct {
	Card *CardWithoutCvv `json:"card,omitempty"`
}

// NewCompletePaymentCardPaymentMethodSpecificInput constructs a new CompletePaymentCardPaymentMethodSpecificInput instance
func NewCompletePaymentCardPaymentMethodSpecificInput() *CompletePaymentCardPaymentMethodSpecificInput {
	return &CompletePaymentCardPaymentMethodSpecificInput{}
}
