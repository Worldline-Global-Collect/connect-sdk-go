// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// TokenSepaDirectDebit represents class TokenSepaDirectDebit
type TokenSepaDirectDebit struct {
	Alias    *string                          `json:"alias,omitempty"`
	Customer *CustomerTokenWithContactDetails `json:"customer,omitempty"`
	Mandate  *MandateSepaDirectDebit          `json:"mandate,omitempty"`
}

// NewTokenSepaDirectDebit constructs a new TokenSepaDirectDebit instance
func NewTokenSepaDirectDebit() *TokenSepaDirectDebit {
	return &TokenSepaDirectDebit{}
}
