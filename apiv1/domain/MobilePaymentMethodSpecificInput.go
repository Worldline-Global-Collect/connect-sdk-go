// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MobilePaymentMethodSpecificInput represents class MobilePaymentMethodSpecificInput
type MobilePaymentMethodSpecificInput struct {
	AuthorizationMode              *string                               `json:"authorizationMode,omitempty"`
	CustomerReference              *string                               `json:"customerReference,omitempty"`
	DecryptedPaymentData           *DecryptedPaymentData                 `json:"decryptedPaymentData,omitempty"`
	EncryptedPaymentData           *string                               `json:"encryptedPaymentData,omitempty"`
	PaymentProduct320SpecificInput *MobilePaymentProduct320SpecificInput `json:"paymentProduct320SpecificInput,omitempty"`
	PaymentProductID               *int32                                `json:"paymentProductId,omitempty"`
	RequiresApproval               *bool                                 `json:"requiresApproval,omitempty"`
	SkipFraudService               *bool                                 `json:"skipFraudService,omitempty"`
}

// NewMobilePaymentMethodSpecificInput constructs a new MobilePaymentMethodSpecificInput instance
func NewMobilePaymentMethodSpecificInput() *MobilePaymentMethodSpecificInput {
	return &MobilePaymentMethodSpecificInput{}
}
