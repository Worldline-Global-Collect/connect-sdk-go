// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CompanyInformation represents class CompanyInformation
type CompanyInformation struct {
	DateOfIncorporation *string `json:"dateOfIncorporation,omitempty"`
	Name                *string `json:"name,omitempty"`
	VatNumber           *string `json:"vatNumber,omitempty"`
}

// NewCompanyInformation constructs a new CompanyInformation instance
func NewCompanyInformation() *CompanyInformation {
	return &CompanyInformation{}
}
