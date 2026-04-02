// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ClickToPayInput represents class ClickToPayInput.
type ClickToPayInput struct {
	CheckoutResponseSignature *string `json:"checkoutResponseSignature,omitempty"`
}

// NewClickToPayInput constructs a new ClickToPayInput instance.
func NewClickToPayInput() *ClickToPayInput {
	return &ClickToPayInput{}
}
