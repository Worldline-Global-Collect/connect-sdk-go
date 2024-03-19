// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProductFieldFormElement represents class PaymentProductFieldFormElement
type PaymentProductFieldFormElement struct {
	Type         *string                `json:"type,omitempty"`
	ValueMapping *[]ValueMappingElement `json:"valueMapping,omitempty"`
}

// NewPaymentProductFieldFormElement constructs a new PaymentProductFieldFormElement
func NewPaymentProductFieldFormElement() *PaymentProductFieldFormElement {
	return &PaymentProductFieldFormElement{}
}
