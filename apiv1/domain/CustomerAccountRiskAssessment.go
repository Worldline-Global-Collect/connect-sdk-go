// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CustomerAccountRiskAssessment represents class CustomerAccountRiskAssessment
type CustomerAccountRiskAssessment struct {
	HasForgottenPassword *bool `json:"hasForgottenPassword,omitempty"`
	HasPassword          *bool `json:"hasPassword,omitempty"`
}

// NewCustomerAccountRiskAssessment constructs a new CustomerAccountRiskAssessment
func NewCustomerAccountRiskAssessment() *CustomerAccountRiskAssessment {
	return &CustomerAccountRiskAssessment{}
}
