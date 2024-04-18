// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// AccountOnFileAttribute represents class AccountOnFileAttribute
type AccountOnFileAttribute struct {
	Key             *string `json:"key,omitempty"`
	MustWriteReason *string `json:"mustWriteReason,omitempty"`
	Status          *string `json:"status,omitempty"`
	Value           *string `json:"value,omitempty"`
}

// NewAccountOnFileAttribute constructs a new AccountOnFileAttribute instance
func NewAccountOnFileAttribute() *AccountOnFileAttribute {
	return &AccountOnFileAttribute{}
}
