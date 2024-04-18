// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RedirectPaymentProduct809SpecificInput represents class RedirectPaymentProduct809SpecificInput
type RedirectPaymentProduct809SpecificInput struct {
	// Deprecated: Use RedirectPaymentMethodSpecificInput.expirationPeriod instead
	ExpirationPeriod *string `json:"expirationPeriod,omitempty"`
	IssuerID         *string `json:"issuerId,omitempty"`
}

// NewRedirectPaymentProduct809SpecificInput constructs a new RedirectPaymentProduct809SpecificInput instance
func NewRedirectPaymentProduct809SpecificInput() *RedirectPaymentProduct809SpecificInput {
	return &RedirectPaymentProduct809SpecificInput{}
}
