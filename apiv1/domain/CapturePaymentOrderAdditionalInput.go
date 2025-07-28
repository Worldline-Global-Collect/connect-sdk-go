// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CapturePaymentOrderAdditionalInput represents class CapturePaymentOrderAdditionalInput
type CapturePaymentOrderAdditionalInput struct {
	AirlineData *AirlineData `json:"airlineData,omitempty"`
	LodgingData *LodgingData `json:"lodgingData,omitempty"`
}

// NewCapturePaymentOrderAdditionalInput constructs a new CapturePaymentOrderAdditionalInput instance
func NewCapturePaymentOrderAdditionalInput() *CapturePaymentOrderAdditionalInput {
	return &CapturePaymentOrderAdditionalInput{}
}
