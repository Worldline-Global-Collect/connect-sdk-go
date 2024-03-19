// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// Card represents class Card
type Card struct {
	CardNumber     *string `json:"cardNumber,omitempty"`
	CardholderName *string `json:"cardholderName,omitempty"`
	Cvv            *string `json:"cvv,omitempty"`
	ExpiryDate     *string `json:"expiryDate,omitempty"`
	IssueNumber    *string `json:"issueNumber,omitempty"`
	PartialPin     *string `json:"partialPin,omitempty"`
}

// NewCard constructs a new Card
func NewCard() *Card {
	return &Card{}
}
