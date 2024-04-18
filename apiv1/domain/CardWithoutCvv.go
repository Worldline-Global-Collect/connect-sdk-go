// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CardWithoutCvv represents class CardWithoutCvv
type CardWithoutCvv struct {
	CardNumber     *string `json:"cardNumber,omitempty"`
	CardholderName *string `json:"cardholderName,omitempty"`
	ExpiryDate     *string `json:"expiryDate,omitempty"`
	IssueNumber    *string `json:"issueNumber,omitempty"`
}

// NewCardWithoutCvv constructs a new CardWithoutCvv instance
func NewCardWithoutCvv() *CardWithoutCvv {
	return &CardWithoutCvv{}
}
