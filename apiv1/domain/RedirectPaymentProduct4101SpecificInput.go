// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RedirectPaymentProduct4101SpecificInput represents class RedirectPaymentProduct4101SpecificInput
type RedirectPaymentProduct4101SpecificInput struct {
	DisplayName           *string `json:"displayName,omitempty"`
	IntegrationType       *string `json:"integrationType,omitempty"`
	VirtualPaymentAddress *string `json:"virtualPaymentAddress,omitempty"`
}

// NewRedirectPaymentProduct4101SpecificInput constructs a new RedirectPaymentProduct4101SpecificInput instance
func NewRedirectPaymentProduct4101SpecificInput() *RedirectPaymentProduct4101SpecificInput {
	return &RedirectPaymentProduct4101SpecificInput{}
}
