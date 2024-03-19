// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// BankRefundMethodSpecificInput represents class BankRefundMethodSpecificInput
type BankRefundMethodSpecificInput struct {
	BankAccountBban *BankAccountBbanRefund `json:"bankAccountBban,omitempty"`
	BankAccountIban *BankAccountIban       `json:"bankAccountIban,omitempty"`
	CountryCode     *string                `json:"countryCode,omitempty"`
}

// NewBankRefundMethodSpecificInput constructs a new BankRefundMethodSpecificInput
func NewBankRefundMethodSpecificInput() *BankRefundMethodSpecificInput {
	return &BankRefundMethodSpecificInput{}
}
