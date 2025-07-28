// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CapturePaymentOrderReferences represents class CapturePaymentOrderReferences
type CapturePaymentOrderReferences struct {
	MerchantCaptureReference *string `json:"merchantCaptureReference,omitempty"`
}

// NewCapturePaymentOrderReferences constructs a new CapturePaymentOrderReferences instance
func NewCapturePaymentOrderReferences() *CapturePaymentOrderReferences {
	return &CapturePaymentOrderReferences{}
}
