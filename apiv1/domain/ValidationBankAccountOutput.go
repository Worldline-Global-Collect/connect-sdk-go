// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ValidationBankAccountOutput represents class ValidationBankAccountOutput
type ValidationBankAccountOutput struct {
	Checks                   *[]ValidationBankAccountCheck `json:"checks,omitempty"`
	NewBankName              *string                       `json:"newBankName,omitempty"`
	ReformattedAccountNumber *string                       `json:"reformattedAccountNumber,omitempty"`
	ReformattedBankCode      *string                       `json:"reformattedBankCode,omitempty"`
	ReformattedBranchCode    *string                       `json:"reformattedBranchCode,omitempty"`
}

// NewValidationBankAccountOutput constructs a new ValidationBankAccountOutput instance
func NewValidationBankAccountOutput() *ValidationBankAccountOutput {
	return &ValidationBankAccountOutput{}
}
