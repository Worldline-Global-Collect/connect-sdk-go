// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RedirectPaymentProduct863SpecificInput represents class RedirectPaymentProduct863SpecificInput
type RedirectPaymentProduct863SpecificInput struct {
	IntegrationType *string `json:"integrationType,omitempty"`
	OpenID          *string `json:"openId,omitempty"`
}

// NewRedirectPaymentProduct863SpecificInput constructs a new RedirectPaymentProduct863SpecificInput instance
func NewRedirectPaymentProduct863SpecificInput() *RedirectPaymentProduct863SpecificInput {
	return &RedirectPaymentProduct863SpecificInput{}
}
