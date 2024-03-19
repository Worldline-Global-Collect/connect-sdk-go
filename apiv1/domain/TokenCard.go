// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// TokenCard represents class TokenCard
type TokenCard struct {
	Alias    *string        `json:"alias,omitempty"`
	Customer *CustomerToken `json:"customer,omitempty"`
	Data     *TokenCardData `json:"data,omitempty"`
}

// NewTokenCard constructs a new TokenCard
func NewTokenCard() *TokenCard {
	return &TokenCard{}
}
