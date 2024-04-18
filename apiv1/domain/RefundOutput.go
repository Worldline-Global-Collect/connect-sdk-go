// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RefundOutput represents class RefundOutput
type RefundOutput struct {
	AmountOfMoney                      *AmountOfMoney                      `json:"amountOfMoney,omitempty"`
	AmountPaid                         *int64                              `json:"amountPaid,omitempty"`
	BankRefundMethodSpecificOutput     *RefundBankMethodSpecificOutput     `json:"bankRefundMethodSpecificOutput,omitempty"`
	CardRefundMethodSpecificOutput     *RefundCardMethodSpecificOutput     `json:"cardRefundMethodSpecificOutput,omitempty"`
	CashRefundMethodSpecificOutput     *RefundCashMethodSpecificOutput     `json:"cashRefundMethodSpecificOutput,omitempty"`
	EInvoiceRefundMethodSpecificOutput *RefundEInvoiceMethodSpecificOutput `json:"eInvoiceRefundMethodSpecificOutput,omitempty"`
	EWalletRefundMethodSpecificOutput  *RefundEWalletMethodSpecificOutput  `json:"eWalletRefundMethodSpecificOutput,omitempty"`
	MobileRefundMethodSpecificOutput   *RefundMobileMethodSpecificOutput   `json:"mobileRefundMethodSpecificOutput,omitempty"`
	PaymentMethod                      *string                             `json:"paymentMethod,omitempty"`
	References                         *PaymentReferences                  `json:"references,omitempty"`
}

// NewRefundOutput constructs a new RefundOutput instance
func NewRefundOutput() *RefundOutput {
	return &RefundOutput{}
}
