// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PersonalName represents class PersonalName
type PersonalName struct {
	FirstName     *string `json:"firstName,omitempty"`
	Surname       *string `json:"surname,omitempty"`
	SurnamePrefix *string `json:"surnamePrefix,omitempty"`
	Title         *string `json:"title,omitempty"`
}

// NewPersonalName constructs a new PersonalName instance
func NewPersonalName() *PersonalName {
	return &PersonalName{}
}
