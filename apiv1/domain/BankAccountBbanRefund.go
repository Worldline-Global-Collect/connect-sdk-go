// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// BankAccountBbanRefund represents class BankAccountBbanRefund
type BankAccountBbanRefund struct {
	AccountHolderName *string `json:"accountHolderName,omitempty"`
	AccountNumber     *string `json:"accountNumber,omitempty"`
	BankCity          *string `json:"bankCity,omitempty"`
	BankCode          *string `json:"bankCode,omitempty"`
	BankName          *string `json:"bankName,omitempty"`
	BranchCode        *string `json:"branchCode,omitempty"`
	CheckDigit        *string `json:"checkDigit,omitempty"`
	CountryCode       *string `json:"countryCode,omitempty"`
	PatronymicName    *string `json:"patronymicName,omitempty"`
	SwiftCode         *string `json:"swiftCode,omitempty"`
}

// NewBankAccountBbanRefund constructs a new BankAccountBbanRefund instance
func NewBankAccountBbanRefund() *BankAccountBbanRefund {
	return &BankAccountBbanRefund{}
}
