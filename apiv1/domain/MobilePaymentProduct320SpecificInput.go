// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MobilePaymentProduct320SpecificInput represents class MobilePaymentProduct320SpecificInput
type MobilePaymentProduct320SpecificInput struct {
	CardholderName *string           `json:"cardholderName,omitempty"`
	ThreeDSecure   *GPayThreeDSecure `json:"threeDSecure,omitempty"`
}

// NewMobilePaymentProduct320SpecificInput constructs a new MobilePaymentProduct320SpecificInput instance
func NewMobilePaymentProduct320SpecificInput() *MobilePaymentProduct320SpecificInput {
	return &MobilePaymentProduct320SpecificInput{}
}
