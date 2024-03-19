// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RefundEWalletMethodSpecificOutput represents class RefundEWalletMethodSpecificOutput
type RefundEWalletMethodSpecificOutput struct {
	PaymentProduct840SpecificOutput *RefundPaymentProduct840SpecificOutput `json:"paymentProduct840SpecificOutput,omitempty"`
	RefundProductID                 *int32                                 `json:"refundProductId,omitempty"`
	TotalAmountPaid                 *int64                                 `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded             *int64                                 `json:"totalAmountRefunded,omitempty"`
}

// NewRefundEWalletMethodSpecificOutput constructs a new RefundEWalletMethodSpecificOutput
func NewRefundEWalletMethodSpecificOutput() *RefundEWalletMethodSpecificOutput {
	return &RefundEWalletMethodSpecificOutput{}
}
