// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentOperation represents class PaymentOperation
type PaymentOperation struct {
	AmountOfMoney *AmountOfMoney `json:"amountOfMoney,omitempty"`
	ID            *string        `json:"id,omitempty"`
	Status        *string        `json:"status,omitempty"`
	Timestamp     *string        `json:"timestamp,omitempty"`
	Type          *string        `json:"type,omitempty"`
}

// NewPaymentOperation constructs a new PaymentOperation instance
func NewPaymentOperation() *PaymentOperation {
	return &PaymentOperation{}
}
