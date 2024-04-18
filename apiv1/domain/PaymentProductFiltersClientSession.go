// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProductFiltersClientSession represents class PaymentProductFiltersClientSession
type PaymentProductFiltersClientSession struct {
	Exclude    *PaymentProductFilter `json:"exclude,omitempty"`
	RestrictTo *PaymentProductFilter `json:"restrictTo,omitempty"`
}

// NewPaymentProductFiltersClientSession constructs a new PaymentProductFiltersClientSession instance
func NewPaymentProductFiltersClientSession() *PaymentProductFiltersClientSession {
	return &PaymentProductFiltersClientSession{}
}
