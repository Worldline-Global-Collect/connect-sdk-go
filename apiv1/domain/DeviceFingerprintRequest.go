// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// DeviceFingerprintRequest represents class DeviceFingerprintRequest
type DeviceFingerprintRequest struct {
	CollectorCallback *string `json:"collectorCallback,omitempty"`
}

// NewDeviceFingerprintRequest constructs a new DeviceFingerprintRequest
func NewDeviceFingerprintRequest() *DeviceFingerprintRequest {
	return &DeviceFingerprintRequest{}
}
