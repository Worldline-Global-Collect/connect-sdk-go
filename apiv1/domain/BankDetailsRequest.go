// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// BankDetailsRequest represents class BankDetailsRequest
type BankDetailsRequest struct {
	BankAccountBban *BankAccountBban `json:"bankAccountBban,omitempty"`
	BankAccountIban *BankAccountIban `json:"bankAccountIban,omitempty"`
}

// NewBankDetailsRequest constructs a new BankDetailsRequest instance
func NewBankDetailsRequest() *BankDetailsRequest {
	return &BankDetailsRequest{}
}
