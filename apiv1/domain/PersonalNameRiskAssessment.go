// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PersonalNameRiskAssessment represents class PersonalNameRiskAssessment
type PersonalNameRiskAssessment struct {
	FirstName     *string `json:"firstName,omitempty"`
	Surname       *string `json:"surname,omitempty"`
	SurnamePrefix *string `json:"surnamePrefix,omitempty"`
}

// NewPersonalNameRiskAssessment constructs a new PersonalNameRiskAssessment instance
func NewPersonalNameRiskAssessment() *PersonalNameRiskAssessment {
	return &PersonalNameRiskAssessment{}
}
