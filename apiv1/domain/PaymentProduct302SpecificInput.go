// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProduct302SpecificInput represents class PaymentProduct302SpecificInput
type PaymentProduct302SpecificInput struct {
	AutomaticReloadBilling *AutomaticReloadBillingDetails `json:"automaticReloadBilling,omitempty"`
	BillingAgreement       *string                        `json:"billingAgreement,omitempty"`
	DeferredBilling        *DeferredBillingDetails        `json:"deferredBilling,omitempty"`
	ManagementURL          *string                        `json:"managementUrl,omitempty"`
	PaymentDescription     *string                        `json:"paymentDescription,omitempty"`
	RegularBilling         *RecurringBillingDetails       `json:"regularBilling,omitempty"`
	TokenNotificationURL   *string                        `json:"tokenNotificationUrl,omitempty"`
	TrialBilling           *RecurringBillingDetails       `json:"trialBilling,omitempty"`
}

// NewPaymentProduct302SpecificInput constructs a new PaymentProduct302SpecificInput instance
func NewPaymentProduct302SpecificInput() *PaymentProduct302SpecificInput {
	return &PaymentProduct302SpecificInput{}
}
