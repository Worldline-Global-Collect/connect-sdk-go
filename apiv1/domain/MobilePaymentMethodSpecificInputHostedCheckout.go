// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MobilePaymentMethodSpecificInputHostedCheckout represents class MobilePaymentMethodSpecificInputHostedCheckout
type MobilePaymentMethodSpecificInputHostedCheckout struct {
	AuthorizationMode                      *string                                             `json:"authorizationMode,omitempty"`
	CustomerReference                      *string                                             `json:"customerReference,omitempty"`
	InitialSchemeTransactionID             *string                                             `json:"initialSchemeTransactionId,omitempty"`
	PaymentProduct302SpecificInput         *MobilePaymentProduct302SpecificInputHostedCheckout `json:"paymentProduct302SpecificInput,omitempty"`
	PaymentProduct320SpecificInput         *MobilePaymentProduct320SpecificInputHostedCheckout `json:"paymentProduct320SpecificInput,omitempty"`
	PaymentProductID                       *int32                                              `json:"paymentProductId,omitempty"`
	Recurring                              *CardRecurrenceDetails                              `json:"recurring,omitempty"`
	RequiresApproval                       *bool                                               `json:"requiresApproval,omitempty"`
	SkipFraudService                       *bool                                               `json:"skipFraudService,omitempty"`
	Token                                  *string                                             `json:"token,omitempty"`
	Tokenize                               *bool                                               `json:"tokenize,omitempty"`
	UnscheduledCardOnFileRequestor         *string                                             `json:"unscheduledCardOnFileRequestor,omitempty"`
	UnscheduledCardOnFileSequenceIndicator *string                                             `json:"unscheduledCardOnFileSequenceIndicator,omitempty"`
}

// NewMobilePaymentMethodSpecificInputHostedCheckout constructs a new MobilePaymentMethodSpecificInputHostedCheckout instance
func NewMobilePaymentMethodSpecificInputHostedCheckout() *MobilePaymentMethodSpecificInputHostedCheckout {
	return &MobilePaymentMethodSpecificInputHostedCheckout{}
}
