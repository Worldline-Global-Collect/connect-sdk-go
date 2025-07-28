// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CapturePaymentRequest represents class CapturePaymentRequest
type CapturePaymentRequest struct {
	Amount  *int64               `json:"amount,omitempty"`
	IsFinal *bool                `json:"isFinal,omitempty"`
	Order   *CapturePaymentOrder `json:"order,omitempty"`
}

// NewCapturePaymentRequest constructs a new CapturePaymentRequest instance
func NewCapturePaymentRequest() *CapturePaymentRequest {
	return &CapturePaymentRequest{}
}
