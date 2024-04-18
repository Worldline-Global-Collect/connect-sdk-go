// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MandatePersonalName represents class MandatePersonalName
type MandatePersonalName struct {
	FirstName *string `json:"firstName,omitempty"`
	Surname   *string `json:"surname,omitempty"`
}

// NewMandatePersonalName constructs a new MandatePersonalName instance
func NewMandatePersonalName() *MandatePersonalName {
	return &MandatePersonalName{}
}
