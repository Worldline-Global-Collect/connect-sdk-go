// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// DeviceFingerprintResponse represents class DeviceFingerprintResponse
type DeviceFingerprintResponse struct {
	DeviceFingerprintTransactionID *string `json:"deviceFingerprintTransactionId,omitempty"`
	HTML                           *string `json:"html,omitempty"`
}

// NewDeviceFingerprintResponse constructs a new DeviceFingerprintResponse instance
func NewDeviceFingerprintResponse() *DeviceFingerprintResponse {
	return &DeviceFingerprintResponse{}
}
