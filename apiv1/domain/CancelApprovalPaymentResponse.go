// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CancelApprovalPaymentResponse represents class CancelApprovalPaymentResponse
type CancelApprovalPaymentResponse struct {
	Payment *Payment `json:"payment,omitempty"`
}

// NewCancelApprovalPaymentResponse constructs a new CancelApprovalPaymentResponse
func NewCancelApprovalPaymentResponse() *CancelApprovalPaymentResponse {
	return &CancelApprovalPaymentResponse{}
}
