// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// BankTransferPaymentMethodSpecificInputBase represents class BankTransferPaymentMethodSpecificInputBase
type BankTransferPaymentMethodSpecificInputBase struct {
	AdditionalReference *string `json:"additionalReference,omitempty"`
	PaymentProductID    *int32  `json:"paymentProductId,omitempty"`
}

// NewBankTransferPaymentMethodSpecificInputBase constructs a new BankTransferPaymentMethodSpecificInputBase instance
func NewBankTransferPaymentMethodSpecificInputBase() *BankTransferPaymentMethodSpecificInputBase {
	return &BankTransferPaymentMethodSpecificInputBase{}
}
