// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// BankAccountIban represents class BankAccountIban
type BankAccountIban struct {
	AccountHolderName *string `json:"accountHolderName,omitempty"`
	Iban              *string `json:"iban,omitempty"`
}

// NewBankAccountIban constructs a new BankAccountIban
func NewBankAccountIban() *BankAccountIban {
	return &BankAccountIban{}
}
