// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PersonalInformationToken represents class PersonalInformationToken
type PersonalInformationToken struct {
	Name *PersonalNameToken `json:"name,omitempty"`
}

// NewPersonalInformationToken constructs a new PersonalInformationToken instance
func NewPersonalInformationToken() *PersonalInformationToken {
	return &PersonalInformationToken{}
}
