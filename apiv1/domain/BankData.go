// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// BankData represents class BankData
type BankData struct {
	NewBankName              *string `json:"newBankName,omitempty"`
	ReformattedAccountNumber *string `json:"reformattedAccountNumber,omitempty"`
	ReformattedBankCode      *string `json:"reformattedBankCode,omitempty"`
	ReformattedBranchCode    *string `json:"reformattedBranchCode,omitempty"`
}

// NewBankData constructs a new BankData instance
func NewBankData() *BankData {
	return &BankData{}
}
