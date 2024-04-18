// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// TokenEWallet represents class TokenEWallet
type TokenEWallet struct {
	Alias    *string           `json:"alias,omitempty"`
	Customer *CustomerToken    `json:"customer,omitempty"`
	Data     *TokenEWalletData `json:"data,omitempty"`
}

// NewTokenEWallet constructs a new TokenEWallet instance
func NewTokenEWallet() *TokenEWallet {
	return &TokenEWallet{}
}
