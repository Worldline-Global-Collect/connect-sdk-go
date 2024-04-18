// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// AccountOnFile represents class AccountOnFile
type AccountOnFile struct {
	Attributes       *[]AccountOnFileAttribute  `json:"attributes,omitempty"`
	DisplayHints     *AccountOnFileDisplayHints `json:"displayHints,omitempty"`
	ID               *int32                     `json:"id,omitempty"`
	PaymentProductID *int32                     `json:"paymentProductId,omitempty"`
}

// NewAccountOnFile constructs a new AccountOnFile instance
func NewAccountOnFile() *AccountOnFile {
	return &AccountOnFile{}
}
