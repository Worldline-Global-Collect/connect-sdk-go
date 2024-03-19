// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// TokenNonSepaDirectDebit represents class TokenNonSepaDirectDebit
type TokenNonSepaDirectDebit struct {
	Alias    *string                    `json:"alias,omitempty"`
	Customer *CustomerToken             `json:"customer,omitempty"`
	Mandate  *MandateNonSepaDirectDebit `json:"mandate,omitempty"`
}

// NewTokenNonSepaDirectDebit constructs a new TokenNonSepaDirectDebit
func NewTokenNonSepaDirectDebit() *TokenNonSepaDirectDebit {
	return &TokenNonSepaDirectDebit{}
}
