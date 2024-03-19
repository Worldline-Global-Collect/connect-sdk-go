// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// TokenCardData represents class TokenCardData
type TokenCardData struct {
	CardWithoutCvv       *CardWithoutCvv `json:"cardWithoutCvv,omitempty"`
	FirstTransactionDate *string         `json:"firstTransactionDate,omitempty"`
	ProviderReference    *string         `json:"providerReference,omitempty"`
}

// NewTokenCardData constructs a new TokenCardData
func NewTokenCardData() *TokenCardData {
	return &TokenCardData{}
}
