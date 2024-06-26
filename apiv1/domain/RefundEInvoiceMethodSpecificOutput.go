// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RefundEInvoiceMethodSpecificOutput represents class RefundEInvoiceMethodSpecificOutput
type RefundEInvoiceMethodSpecificOutput struct {
	RefundProductID     *int32 `json:"refundProductId,omitempty"`
	TotalAmountPaid     *int64 `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded *int64 `json:"totalAmountRefunded,omitempty"`
}

// NewRefundEInvoiceMethodSpecificOutput constructs a new RefundEInvoiceMethodSpecificOutput instance
func NewRefundEInvoiceMethodSpecificOutput() *RefundEInvoiceMethodSpecificOutput {
	return &RefundEInvoiceMethodSpecificOutput{}
}
