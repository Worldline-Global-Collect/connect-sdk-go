// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ClickToPayConfigurationVisa represents class ClickToPayConfigurationVisa
type ClickToPayConfigurationVisa struct {
	EncryptionKey  *string `json:"encryptionKey,omitempty"`
	NModulus       *string `json:"nModulus,omitempty"`
	SrcDpaID       *string `json:"srcDpaId,omitempty"`
	SrcInitiatorID *string `json:"srcInitiatorId,omitempty"`
}

// NewClickToPayConfigurationVisa constructs a new ClickToPayConfigurationVisa instance
func NewClickToPayConfigurationVisa() *ClickToPayConfigurationVisa {
	return &ClickToPayConfigurationVisa{}
}
