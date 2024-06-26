// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PersonalNameToken represents class PersonalNameToken
type PersonalNameToken struct {
	FirstName     *string `json:"firstName,omitempty"`
	Surname       *string `json:"surname,omitempty"`
	SurnamePrefix *string `json:"surnamePrefix,omitempty"`
}

// NewPersonalNameToken constructs a new PersonalNameToken instance
func NewPersonalNameToken() *PersonalNameToken {
	return &PersonalNameToken{}
}
