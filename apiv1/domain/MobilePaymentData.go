// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MobilePaymentData represents class MobilePaymentData
type MobilePaymentData struct {
	Dpan       *string `json:"dpan,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
}

// NewMobilePaymentData constructs a new MobilePaymentData instance
func NewMobilePaymentData() *MobilePaymentData {
	return &MobilePaymentData{}
}
