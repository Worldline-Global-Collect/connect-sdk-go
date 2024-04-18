// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CustomerDeviceRiskAssessment represents class CustomerDeviceRiskAssessment
type CustomerDeviceRiskAssessment struct {
	DefaultFormFill                *string `json:"defaultFormFill,omitempty"`
	DeviceFingerprintTransactionID *string `json:"deviceFingerprintTransactionId,omitempty"`
}

// NewCustomerDeviceRiskAssessment constructs a new CustomerDeviceRiskAssessment instance
func NewCustomerDeviceRiskAssessment() *CustomerDeviceRiskAssessment {
	return &CustomerDeviceRiskAssessment{}
}
