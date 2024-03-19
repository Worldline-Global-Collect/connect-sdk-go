// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProduct840SpecificOutput represents class PaymentProduct840SpecificOutput
type PaymentProduct840SpecificOutput struct {
	BillingAddress        *Address                          `json:"billingAddress,omitempty"`
	CustomerAccount       *PaymentProduct840CustomerAccount `json:"customerAccount,omitempty"`
	CustomerAddress       *Address                          `json:"customerAddress,omitempty"`
	ProtectionEligibility *ProtectionEligibility            `json:"protectionEligibility,omitempty"`
}

// NewPaymentProduct840SpecificOutput constructs a new PaymentProduct840SpecificOutput
func NewPaymentProduct840SpecificOutput() *PaymentProduct840SpecificOutput {
	return &PaymentProduct840SpecificOutput{}
}
