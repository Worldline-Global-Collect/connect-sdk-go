// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// AccountOnFileDisplayHints represents class AccountOnFileDisplayHints
type AccountOnFileDisplayHints struct {
	LabelTemplate *[]LabelTemplateElement `json:"labelTemplate,omitempty"`
	Logo          *string                 `json:"logo,omitempty"`
}

// NewAccountOnFileDisplayHints constructs a new AccountOnFileDisplayHints
func NewAccountOnFileDisplayHints() *AccountOnFileDisplayHints {
	return &AccountOnFileDisplayHints{}
}
