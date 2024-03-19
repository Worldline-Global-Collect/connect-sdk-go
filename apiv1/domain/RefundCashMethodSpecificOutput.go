// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RefundCashMethodSpecificOutput represents class RefundCashMethodSpecificOutput
type RefundCashMethodSpecificOutput struct {
	RefundProductID     *int32 `json:"refundProductId,omitempty"`
	TotalAmountPaid     *int64 `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded *int64 `json:"totalAmountRefunded,omitempty"`
}

// NewRefundCashMethodSpecificOutput constructs a new RefundCashMethodSpecificOutput
func NewRefundCashMethodSpecificOutput() *RefundCashMethodSpecificOutput {
	return &RefundCashMethodSpecificOutput{}
}
