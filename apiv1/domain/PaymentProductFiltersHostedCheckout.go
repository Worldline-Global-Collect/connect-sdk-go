// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProductFiltersHostedCheckout represents class PaymentProductFiltersHostedCheckout
type PaymentProductFiltersHostedCheckout struct {
	Exclude    *PaymentProductFilter `json:"exclude,omitempty"`
	RestrictTo *PaymentProductFilter `json:"restrictTo,omitempty"`
	TokensOnly *bool                 `json:"tokensOnly,omitempty"`
}

// NewPaymentProductFiltersHostedCheckout constructs a new PaymentProductFiltersHostedCheckout
func NewPaymentProductFiltersHostedCheckout() *PaymentProductFiltersHostedCheckout {
	return &PaymentProductFiltersHostedCheckout{}
}
