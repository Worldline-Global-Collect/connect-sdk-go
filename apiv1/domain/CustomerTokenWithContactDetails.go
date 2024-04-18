// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CustomerTokenWithContactDetails represents class CustomerTokenWithContactDetails
type CustomerTokenWithContactDetails struct {
	BillingAddress      *Address                  `json:"billingAddress,omitempty"`
	CompanyInformation  *CompanyInformation       `json:"companyInformation,omitempty"`
	ContactDetails      *ContactDetailsToken      `json:"contactDetails,omitempty"`
	MerchantCustomerID  *string                   `json:"merchantCustomerId,omitempty"`
	PersonalInformation *PersonalInformationToken `json:"personalInformation,omitempty"`
	// Deprecated: Use companyInformation.vatNumber instead
	VatNumber           *string                   `json:"vatNumber,omitempty"`
}

// NewCustomerTokenWithContactDetails constructs a new CustomerTokenWithContactDetails instance
func NewCustomerTokenWithContactDetails() *CustomerTokenWithContactDetails {
	return &CustomerTokenWithContactDetails{}
}
