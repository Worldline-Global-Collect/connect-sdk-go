// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RefundResult represents class RefundResult
type RefundResult struct {
	ID           *string            `json:"id,omitempty"`
	RefundOutput *RefundOutput      `json:"refundOutput,omitempty"`
	Status       *string            `json:"status,omitempty"`
	StatusOutput *OrderStatusOutput `json:"statusOutput,omitempty"`
}

// NewRefundResult constructs a new RefundResult
func NewRefundResult() *RefundResult {
	return &RefundResult{}
}
