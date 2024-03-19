// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PayoutDetails represents class PayoutDetails
type PayoutDetails struct {
	AmountOfMoney *AmountOfMoney    `json:"amountOfMoney,omitempty"`
	Customer      *PayoutCustomer   `json:"customer,omitempty"`
	References    *PayoutReferences `json:"references,omitempty"`
}

// NewPayoutDetails constructs a new PayoutDetails
func NewPayoutDetails() *PayoutDetails {
	return &PayoutDetails{}
}
