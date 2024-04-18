// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// BankTransferPaymentMethodSpecificOutput represents class BankTransferPaymentMethodSpecificOutput
type BankTransferPaymentMethodSpecificOutput struct {
	FraudResults     *FraudResults `json:"fraudResults,omitempty"`
	PaymentProductID *int32        `json:"paymentProductId,omitempty"`
}

// NewBankTransferPaymentMethodSpecificOutput constructs a new BankTransferPaymentMethodSpecificOutput instance
func NewBankTransferPaymentMethodSpecificOutput() *BankTransferPaymentMethodSpecificOutput {
	return &BankTransferPaymentMethodSpecificOutput{}
}
