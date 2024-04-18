// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// AmountOfMoney represents class AmountOfMoney
type AmountOfMoney struct {
	Amount       *int64  `json:"amount,omitempty"`
	CurrencyCode *string `json:"currencyCode,omitempty"`
}

// NewAmountOfMoney constructs a new AmountOfMoney instance
func NewAmountOfMoney() *AmountOfMoney {
	return &AmountOfMoney{}
}
