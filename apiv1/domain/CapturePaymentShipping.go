// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CapturePaymentShipping represents class CapturePaymentShipping.
type CapturePaymentShipping struct {
	Address        *AddressPersonal `json:"address,omitempty"`
	EmailAddress   *string          `json:"emailAddress,omitempty"`
	ShippedFromZip *string          `json:"shippedFromZip,omitempty"`
	TrackingNumber *string          `json:"trackingNumber,omitempty"`
}

// NewCapturePaymentShipping constructs a new CapturePaymentShipping instance.
func NewCapturePaymentShipping() *CapturePaymentShipping {
	return &CapturePaymentShipping{}
}
