// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// OrderApprovePayment represents class OrderApprovePayment
type OrderApprovePayment struct {
	AdditionalInput *AdditionalOrderInputAirlineData `json:"additionalInput,omitempty"`
	Customer        *CustomerApprovePayment          `json:"customer,omitempty"`
	References      *OrderReferencesApprovePayment   `json:"references,omitempty"`
}

// NewOrderApprovePayment constructs a new OrderApprovePayment
func NewOrderApprovePayment() *OrderApprovePayment {
	return &OrderApprovePayment{}
}
