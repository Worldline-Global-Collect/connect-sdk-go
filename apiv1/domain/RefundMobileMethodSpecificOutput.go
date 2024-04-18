// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RefundMobileMethodSpecificOutput represents class RefundMobileMethodSpecificOutput
type RefundMobileMethodSpecificOutput struct {
	Network             *string `json:"network,omitempty"`
	RefundProductID     *int32  `json:"refundProductId,omitempty"`
	TotalAmountPaid     *int64  `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded *int64  `json:"totalAmountRefunded,omitempty"`
}

// NewRefundMobileMethodSpecificOutput constructs a new RefundMobileMethodSpecificOutput instance
func NewRefundMobileMethodSpecificOutput() *RefundMobileMethodSpecificOutput {
	return &RefundMobileMethodSpecificOutput{}
}
