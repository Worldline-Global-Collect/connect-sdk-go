// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentCreationReferences represents class PaymentCreationReferences
type PaymentCreationReferences struct {
	AdditionalReference *string `json:"additionalReference,omitempty"`
	ExternalReference   *string `json:"externalReference,omitempty"`
}

// NewPaymentCreationReferences constructs a new PaymentCreationReferences
func NewPaymentCreationReferences() *PaymentCreationReferences {
	return &PaymentCreationReferences{}
}
