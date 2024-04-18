// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CardPayoutMethodSpecificInput represents class CardPayoutMethodSpecificInput
type CardPayoutMethodSpecificInput struct {
	Card             *Card            `json:"card,omitempty"`
	PaymentProductID *int32           `json:"paymentProductId,omitempty"`
	Recipient        *PayoutRecipient `json:"recipient,omitempty"`
	Token            *string          `json:"token,omitempty"`
}

// NewCardPayoutMethodSpecificInput constructs a new CardPayoutMethodSpecificInput instance
func NewCardPayoutMethodSpecificInput() *CardPayoutMethodSpecificInput {
	return &CardPayoutMethodSpecificInput{}
}
