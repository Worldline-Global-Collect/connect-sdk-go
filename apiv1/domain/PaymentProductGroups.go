// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProductGroups represents class PaymentProductGroups
type PaymentProductGroups struct {
	PaymentProductGroups *[]PaymentProductGroup `json:"paymentProductGroups,omitempty"`
}

// NewPaymentProductGroups constructs a new PaymentProductGroups
func NewPaymentProductGroups() *PaymentProductGroups {
	return &PaymentProductGroups{}
}
