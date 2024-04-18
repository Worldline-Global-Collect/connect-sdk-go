// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MandatePersonalInformation represents class MandatePersonalInformation
type MandatePersonalInformation struct {
	Name  *MandatePersonalName `json:"name,omitempty"`
	Title *string              `json:"title,omitempty"`
}

// NewMandatePersonalInformation constructs a new MandatePersonalInformation instance
func NewMandatePersonalInformation() *MandatePersonalInformation {
	return &MandatePersonalInformation{}
}
