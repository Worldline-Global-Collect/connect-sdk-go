// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// HostedMandateManagementSpecificInput represents class HostedMandateManagementSpecificInput
type HostedMandateManagementSpecificInput struct {
	Locale         *string `json:"locale,omitempty"`
	ReturnURL      *string `json:"returnUrl,omitempty"`
	ShowResultPage *bool   `json:"showResultPage,omitempty"`
	Variant        *string `json:"variant,omitempty"`
}

// NewHostedMandateManagementSpecificInput constructs a new HostedMandateManagementSpecificInput
func NewHostedMandateManagementSpecificInput() *HostedMandateManagementSpecificInput {
	return &HostedMandateManagementSpecificInput{}
}
