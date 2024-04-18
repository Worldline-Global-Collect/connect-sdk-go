// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ValidationBankAccountCheck represents class ValidationBankAccountCheck
type ValidationBankAccountCheck struct {
	Code        *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
	Result      *string `json:"result,omitempty"`
}

// NewValidationBankAccountCheck constructs a new ValidationBankAccountCheck instance
func NewValidationBankAccountCheck() *ValidationBankAccountCheck {
	return &ValidationBankAccountCheck{}
}
