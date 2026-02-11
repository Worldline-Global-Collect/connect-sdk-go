// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ClickToPayConfiguration represents class ClickToPayConfiguration
type ClickToPayConfiguration struct {
	Mastercard *ClickToPayConfigurationMastercard `json:"mastercard,omitempty"`
	Visa       *ClickToPayConfigurationVisa       `json:"visa,omitempty"`
}

// NewClickToPayConfiguration constructs a new ClickToPayConfiguration instance
func NewClickToPayConfiguration() *ClickToPayConfiguration {
	return &ClickToPayConfiguration{}
}
