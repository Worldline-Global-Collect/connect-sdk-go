// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// DeferredBillingDetails represents class DeferredBillingDetails
type DeferredBillingDetails struct {
	DeferredPaymentDate          *string `json:"deferredPaymentDate,omitempty"`
	Description                  *string `json:"description,omitempty"`
	FreeCancellationDate         *string `json:"freeCancellationDate,omitempty"`
	FreeCancellationDateTimeZone *string `json:"freeCancellationDateTimeZone,omitempty"`
}

// NewDeferredBillingDetails constructs a new DeferredBillingDetails instance
func NewDeferredBillingDetails() *DeferredBillingDetails {
	return &DeferredBillingDetails{}
}
