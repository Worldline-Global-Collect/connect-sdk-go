// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProduct806SpecificOutput represents class PaymentProduct806SpecificOutput
type PaymentProduct806SpecificOutput struct {
	BillingAddress  *Address            `json:"billingAddress,omitempty"`
	CustomerAccount *TrustlyBankAccount `json:"customerAccount,omitempty"`
}

// NewPaymentProduct806SpecificOutput constructs a new PaymentProduct806SpecificOutput
func NewPaymentProduct806SpecificOutput() *PaymentProduct806SpecificOutput {
	return &PaymentProduct806SpecificOutput{}
}
