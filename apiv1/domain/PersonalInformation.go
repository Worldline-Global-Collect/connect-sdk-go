// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PersonalInformation represents class PersonalInformation
type PersonalInformation struct {
	DateOfBirth    *string                 `json:"dateOfBirth,omitempty"`
	Gender         *string                 `json:"gender,omitempty"`
	Identification *PersonalIdentification `json:"identification,omitempty"`
	Name           *PersonalName           `json:"name,omitempty"`
}

// NewPersonalInformation constructs a new PersonalInformation instance
func NewPersonalInformation() *PersonalInformation {
	return &PersonalInformation{}
}
