// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RefundResponse represents class RefundResponse
type RefundResponse struct {
	ID           *string            `json:"id,omitempty"`
	RefundOutput *RefundOutput      `json:"refundOutput,omitempty"`
	Status       *string            `json:"status,omitempty"`
	StatusOutput *OrderStatusOutput `json:"statusOutput,omitempty"`
}

// NewRefundResponse constructs a new RefundResponse
func NewRefundResponse() *RefundResponse {
	return &RefundResponse{}
}
