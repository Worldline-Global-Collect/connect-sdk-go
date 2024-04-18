// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MicrosoftFraudResults represents class MicrosoftFraudResults
type MicrosoftFraudResults struct {
	ClauseName        *string `json:"clauseName,omitempty"`
	DeviceCountryCode *string `json:"deviceCountryCode,omitempty"`
	DeviceID          *string `json:"deviceId,omitempty"`
	FraudScore        *int32  `json:"fraudScore,omitempty"`
	PolicyApplied     *string `json:"policyApplied,omitempty"`
	TrueIPAddress     *string `json:"trueIpAddress,omitempty"`
	UserDeviceType    *string `json:"userDeviceType,omitempty"`
}

// NewMicrosoftFraudResults constructs a new MicrosoftFraudResults instance
func NewMicrosoftFraudResults() *MicrosoftFraudResults {
	return &MicrosoftFraudResults{}
}
