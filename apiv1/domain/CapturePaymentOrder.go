// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CapturePaymentOrder represents class CapturePaymentOrder
type CapturePaymentOrder struct {
	AdditionalInput *CapturePaymentOrderAdditionalInput `json:"additionalInput,omitempty"`
	References      *CapturePaymentOrderReferences      `json:"references,omitempty"`
}

// NewCapturePaymentOrder constructs a new CapturePaymentOrder instance
func NewCapturePaymentOrder() *CapturePaymentOrder {
	return &CapturePaymentOrder{}
}
