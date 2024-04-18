// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MobilePaymentMethodSpecificOutput represents class MobilePaymentMethodSpecificOutput
type MobilePaymentMethodSpecificOutput struct {
	AuthorisationCode   *string              `json:"authorisationCode,omitempty"`
	FraudResults        *CardFraudResults    `json:"fraudResults,omitempty"`
	Network             *string              `json:"network,omitempty"`
	PaymentData         *MobilePaymentData   `json:"paymentData,omitempty"`
	PaymentProductID    *int32               `json:"paymentProductId,omitempty"`
	ThreeDSecureResults *ThreeDSecureResults `json:"threeDSecureResults,omitempty"`
}

// NewMobilePaymentMethodSpecificOutput constructs a new MobilePaymentMethodSpecificOutput instance
func NewMobilePaymentMethodSpecificOutput() *MobilePaymentMethodSpecificOutput {
	return &MobilePaymentMethodSpecificOutput{}
}
