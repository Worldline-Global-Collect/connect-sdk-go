// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CashPaymentMethodSpecificOutput represents class CashPaymentMethodSpecificOutput
type CashPaymentMethodSpecificOutput struct {
	FraudResults     *FraudResults `json:"fraudResults,omitempty"`
	PaymentProductID *int32        `json:"paymentProductId,omitempty"`
}

// NewCashPaymentMethodSpecificOutput constructs a new CashPaymentMethodSpecificOutput instance
func NewCashPaymentMethodSpecificOutput() *CashPaymentMethodSpecificOutput {
	return &CashPaymentMethodSpecificOutput{}
}
