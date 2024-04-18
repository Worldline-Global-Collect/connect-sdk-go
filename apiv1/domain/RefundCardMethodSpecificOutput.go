// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RefundCardMethodSpecificOutput represents class RefundCardMethodSpecificOutput
type RefundCardMethodSpecificOutput struct {
	AuthorisationCode   *string         `json:"authorisationCode,omitempty"`
	Card                *CardEssentials `json:"card,omitempty"`
	RefundProductID     *int32          `json:"refundProductId,omitempty"`
	TotalAmountPaid     *int64          `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded *int64          `json:"totalAmountRefunded,omitempty"`
}

// NewRefundCardMethodSpecificOutput constructs a new RefundCardMethodSpecificOutput instance
func NewRefundCardMethodSpecificOutput() *RefundCardMethodSpecificOutput {
	return &RefundCardMethodSpecificOutput{}
}
