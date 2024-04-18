// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CreateHostedCheckoutResponse represents class CreateHostedCheckoutResponse
type CreateHostedCheckoutResponse struct {
	RETURNMAC          *string   `json:"RETURNMAC,omitempty"`
	HostedCheckoutID   *string   `json:"hostedCheckoutId,omitempty"`
	InvalidTokens      *[]string `json:"invalidTokens,omitempty"`
	MerchantReference  *string   `json:"merchantReference,omitempty"`
	PartialRedirectURL *string   `json:"partialRedirectUrl,omitempty"`
}

// NewCreateHostedCheckoutResponse constructs a new CreateHostedCheckoutResponse instance
func NewCreateHostedCheckoutResponse() *CreateHostedCheckoutResponse {
	return &CreateHostedCheckoutResponse{}
}
