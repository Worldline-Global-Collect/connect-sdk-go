// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RefundErrorResponse represents class RefundErrorResponse
type RefundErrorResponse struct {
	ErrorID      *string       `json:"errorId,omitempty"`
	Errors       *[]APIError   `json:"errors,omitempty"`
	RefundResult *RefundResult `json:"refundResult,omitempty"`
}

// NewRefundErrorResponse constructs a new RefundErrorResponse instance
func NewRefundErrorResponse() *RefundErrorResponse {
	return &RefundErrorResponse{}
}
