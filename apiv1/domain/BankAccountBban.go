// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// BankAccountBban represents class BankAccountBban
type BankAccountBban struct {
	AccountHolderName *string `json:"accountHolderName,omitempty"`
	AccountNumber     *string `json:"accountNumber,omitempty"`
	BankCode          *string `json:"bankCode,omitempty"`
	BankName          *string `json:"bankName,omitempty"`
	BranchCode        *string `json:"branchCode,omitempty"`
	CheckDigit        *string `json:"checkDigit,omitempty"`
	CountryCode       *string `json:"countryCode,omitempty"`
}

// NewBankAccountBban constructs a new BankAccountBban instance
func NewBankAccountBban() *BankAccountBban {
	return &BankAccountBban{}
}
