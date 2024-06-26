// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PersonalIdentification represents class PersonalIdentification
type PersonalIdentification struct {
	IDIssuingCountryCode *string `json:"idIssuingCountryCode,omitempty"`
	IDType               *string `json:"idType,omitempty"`
	IDValue              *string `json:"idValue,omitempty"`
}

// NewPersonalIdentification constructs a new PersonalIdentification instance
func NewPersonalIdentification() *PersonalIdentification {
	return &PersonalIdentification{}
}
