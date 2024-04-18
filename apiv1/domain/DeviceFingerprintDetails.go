// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// DeviceFingerprintDetails represents class DeviceFingerprintDetails
type DeviceFingerprintDetails struct {
	PaymentID                  *string `json:"paymentId,omitempty"`
	RawDeviceFingerprintOutput *string `json:"rawDeviceFingerprintOutput,omitempty"`
}

// NewDeviceFingerprintDetails constructs a new DeviceFingerprintDetails instance
func NewDeviceFingerprintDetails() *DeviceFingerprintDetails {
	return &DeviceFingerprintDetails{}
}
