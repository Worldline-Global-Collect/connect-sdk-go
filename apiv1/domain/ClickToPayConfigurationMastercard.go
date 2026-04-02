// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ClickToPayConfigurationMastercard represents class ClickToPayConfigurationMastercard.
type ClickToPayConfigurationMastercard struct {
	SrcInitiatorID *string `json:"srcInitiatorId,omitempty"`
	SrciDpaID      *string `json:"srciDpaId,omitempty"`
}

// NewClickToPayConfigurationMastercard constructs a new ClickToPayConfigurationMastercard instance.
func NewClickToPayConfigurationMastercard() *ClickToPayConfigurationMastercard {
	return &ClickToPayConfigurationMastercard{}
}
