// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ApproveRefundRequest represents class ApproveRefundRequest
type ApproveRefundRequest struct {
	Amount *int64 `json:"amount,omitempty"`
}

// NewApproveRefundRequest constructs a new ApproveRefundRequest
func NewApproveRefundRequest() *ApproveRefundRequest {
	return &ApproveRefundRequest{}
}
