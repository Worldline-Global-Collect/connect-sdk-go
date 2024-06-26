// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RedirectPaymentProduct869SpecificInput represents class RedirectPaymentProduct869SpecificInput
type RedirectPaymentProduct869SpecificInput struct {
	IssuerID         *string `json:"issuerId,omitempty"`
	ResidentIDName   *string `json:"residentIdName,omitempty"`
	ResidentIDNumber *string `json:"residentIdNumber,omitempty"`
}

// NewRedirectPaymentProduct869SpecificInput constructs a new RedirectPaymentProduct869SpecificInput instance
func NewRedirectPaymentProduct869SpecificInput() *RedirectPaymentProduct869SpecificInput {
	return &RedirectPaymentProduct869SpecificInput{}
}
