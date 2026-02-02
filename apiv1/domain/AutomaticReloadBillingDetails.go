// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// AutomaticReloadBillingDetails represents class AutomaticReloadBillingDetails
type AutomaticReloadBillingDetails struct {
	AutomaticReloadPaymentThresholdAmount *int64  `json:"automaticReloadPaymentThresholdAmount,omitempty"`
	Description                           *string `json:"description,omitempty"`
}

// NewAutomaticReloadBillingDetails constructs a new AutomaticReloadBillingDetails instance
func NewAutomaticReloadBillingDetails() *AutomaticReloadBillingDetails {
	return &AutomaticReloadBillingDetails{}
}
