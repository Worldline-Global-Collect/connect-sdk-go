// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// BankDetailsResponse represents class BankDetailsResponse
type BankDetailsResponse struct {
	BankAccountBban *BankAccountBban `json:"bankAccountBban,omitempty"`
	BankAccountIban *BankAccountIban `json:"bankAccountIban,omitempty"`
	BankData        *BankData        `json:"bankData,omitempty"`
	Swift           *Swift           `json:"swift,omitempty"`
}

// NewBankDetailsResponse constructs a new BankDetailsResponse
func NewBankDetailsResponse() *BankDetailsResponse {
	return &BankDetailsResponse{}
}
