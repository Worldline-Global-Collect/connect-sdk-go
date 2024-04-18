// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RefundBankMethodSpecificOutput represents class RefundBankMethodSpecificOutput
type RefundBankMethodSpecificOutput struct {
	RefundProductID     *int32 `json:"refundProductId,omitempty"`
	TotalAmountPaid     *int64 `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded *int64 `json:"totalAmountRefunded,omitempty"`
}

// NewRefundBankMethodSpecificOutput constructs a new RefundBankMethodSpecificOutput instance
func NewRefundBankMethodSpecificOutput() *RefundBankMethodSpecificOutput {
	return &RefundBankMethodSpecificOutput{}
}
