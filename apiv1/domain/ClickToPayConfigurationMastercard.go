// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ClickToPayConfigurationMastercard represents class ClickToPayConfigurationMastercard
type ClickToPayConfigurationMastercard struct {
	SrcDpaID       *string `json:"srcDpaId,omitempty"`
	SrcInitiatorID *string `json:"srcInitiatorId,omitempty"`
}

// NewClickToPayConfigurationMastercard constructs a new ClickToPayConfigurationMastercard instance
func NewClickToPayConfigurationMastercard() *ClickToPayConfigurationMastercard {
	return &ClickToPayConfigurationMastercard{}
}
