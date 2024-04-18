// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CardEssentials represents class CardEssentials
type CardEssentials struct {
	CardNumber     *string `json:"cardNumber,omitempty"`
	CardholderName *string `json:"cardholderName,omitempty"`
	ExpiryDate     *string `json:"expiryDate,omitempty"`
}

// NewCardEssentials constructs a new CardEssentials instance
func NewCardEssentials() *CardEssentials {
	return &CardEssentials{}
}
