// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// TokenSepaDirectDebitWithoutCreditor represents class TokenSepaDirectDebitWithoutCreditor
type TokenSepaDirectDebitWithoutCreditor struct {
	Alias    *string                                `json:"alias,omitempty"`
	Customer *CustomerTokenWithContactDetails       `json:"customer,omitempty"`
	Mandate  *MandateSepaDirectDebitWithoutCreditor `json:"mandate,omitempty"`
}

// NewTokenSepaDirectDebitWithoutCreditor constructs a new TokenSepaDirectDebitWithoutCreditor instance
func NewTokenSepaDirectDebitWithoutCreditor() *TokenSepaDirectDebitWithoutCreditor {
	return &TokenSepaDirectDebitWithoutCreditor{}
}
