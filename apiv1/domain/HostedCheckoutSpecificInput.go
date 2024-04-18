// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// HostedCheckoutSpecificInput represents class HostedCheckoutSpecificInput
type HostedCheckoutSpecificInput struct {
	IsRecurring           *bool                                `json:"isRecurring,omitempty"`
	Locale                *string                              `json:"locale,omitempty"`
	PaymentProductFilters *PaymentProductFiltersHostedCheckout `json:"paymentProductFilters,omitempty"`
	RecurringPaymentsData *RecurringPaymentsData               `json:"recurringPaymentsData,omitempty"`
	ReturnCancelState     *bool                                `json:"returnCancelState,omitempty"`
	ReturnURL             *string                              `json:"returnUrl,omitempty"`
	ShowResultPage        *bool                                `json:"showResultPage,omitempty"`
	Tokens                *string                              `json:"tokens,omitempty"`
	ValidateShoppingCart  *bool                                `json:"validateShoppingCart,omitempty"`
	Variant               *string                              `json:"variant,omitempty"`
}

// NewHostedCheckoutSpecificInput constructs a new HostedCheckoutSpecificInput instance
func NewHostedCheckoutSpecificInput() *HostedCheckoutSpecificInput {
	return &HostedCheckoutSpecificInput{}
}
