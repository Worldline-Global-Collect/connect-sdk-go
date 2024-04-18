// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// InstallmentOptions represents class InstallmentOptions
type InstallmentOptions struct {
	DisplayHints     *InstallmentDisplayHints `json:"displayHints,omitempty"`
	ID               *string                  `json:"id,omitempty"`
	InstallmentPlans *[]Installments          `json:"installmentPlans,omitempty"`
}

// NewInstallmentOptions constructs a new InstallmentOptions instance
func NewInstallmentOptions() *InstallmentOptions {
	return &InstallmentOptions{}
}
