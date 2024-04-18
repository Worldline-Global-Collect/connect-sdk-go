// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MobilePaymentProduct320SpecificInputHostedCheckout represents class MobilePaymentProduct320SpecificInputHostedCheckout
type MobilePaymentProduct320SpecificInputHostedCheckout struct {
	MerchantName   *string           `json:"merchantName,omitempty"`
	MerchantOrigin *string           `json:"merchantOrigin,omitempty"`
	ThreeDSecure   *GPayThreeDSecure `json:"threeDSecure,omitempty"`
}

// NewMobilePaymentProduct320SpecificInputHostedCheckout constructs a new MobilePaymentProduct320SpecificInputHostedCheckout instance
func NewMobilePaymentProduct320SpecificInputHostedCheckout() *MobilePaymentProduct320SpecificInputHostedCheckout {
	return &MobilePaymentProduct320SpecificInputHostedCheckout{}
}
