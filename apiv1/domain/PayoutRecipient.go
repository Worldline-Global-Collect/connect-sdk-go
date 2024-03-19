// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PayoutRecipient represents class PayoutRecipient
type PayoutRecipient struct {
	FirstName     *string `json:"firstName,omitempty"`
	Surname       *string `json:"surname,omitempty"`
	SurnamePrefix *string `json:"surnamePrefix,omitempty"`
}

// NewPayoutRecipient constructs a new PayoutRecipient
func NewPayoutRecipient() *PayoutRecipient {
	return &PayoutRecipient{}
}
