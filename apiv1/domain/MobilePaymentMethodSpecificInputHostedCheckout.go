// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MobilePaymentMethodSpecificInputHostedCheckout represents class MobilePaymentMethodSpecificInputHostedCheckout
type MobilePaymentMethodSpecificInputHostedCheckout struct {
	AuthorizationMode              *string                                             `json:"authorizationMode,omitempty"`
	CustomerReference              *string                                             `json:"customerReference,omitempty"`
	PaymentProduct302SpecificInput *MobilePaymentProduct302SpecificInputHostedCheckout `json:"paymentProduct302SpecificInput,omitempty"`
	PaymentProduct320SpecificInput *MobilePaymentProduct320SpecificInputHostedCheckout `json:"paymentProduct320SpecificInput,omitempty"`
	PaymentProductID               *int32                                              `json:"paymentProductId,omitempty"`
	RequiresApproval               *bool                                               `json:"requiresApproval,omitempty"`
	SkipFraudService               *bool                                               `json:"skipFraudService,omitempty"`
}

// NewMobilePaymentMethodSpecificInputHostedCheckout constructs a new MobilePaymentMethodSpecificInputHostedCheckout instance
func NewMobilePaymentMethodSpecificInputHostedCheckout() *MobilePaymentMethodSpecificInputHostedCheckout {
	return &MobilePaymentMethodSpecificInputHostedCheckout{}
}
