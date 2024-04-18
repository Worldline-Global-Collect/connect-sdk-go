// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// TrustlyBankAccount represents class TrustlyBankAccount
type TrustlyBankAccount struct {
	AccountLastDigits          *string `json:"accountLastDigits,omitempty"`
	BankName                   *string `json:"bankName,omitempty"`
	Clearinghouse              *string `json:"clearinghouse,omitempty"`
	PersonIdentificationNumber *string `json:"personIdentificationNumber,omitempty"`
}

// NewTrustlyBankAccount constructs a new TrustlyBankAccount instance
func NewTrustlyBankAccount() *TrustlyBankAccount {
	return &TrustlyBankAccount{}
}
