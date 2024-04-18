// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// InstallmentDisplayHints represents class InstallmentDisplayHints
type InstallmentDisplayHints struct {
	DisplayOrder *int32  `json:"displayOrder,omitempty"`
	Label        *string `json:"label,omitempty"`
	Logo         *string `json:"logo,omitempty"`
}

// NewInstallmentDisplayHints constructs a new InstallmentDisplayHints instance
func NewInstallmentDisplayHints() *InstallmentDisplayHints {
	return &InstallmentDisplayHints{}
}
