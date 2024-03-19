// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// OrderOutput represents class OrderOutput
type OrderOutput struct {
	AmountOfMoney *AmountOfMoney     `json:"amountOfMoney,omitempty"`
	References    *PaymentReferences `json:"references,omitempty"`
}

// NewOrderOutput constructs a new OrderOutput
func NewOrderOutput() *OrderOutput {
	return &OrderOutput{}
}
