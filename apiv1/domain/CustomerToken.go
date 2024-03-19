// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CustomerToken represents class CustomerToken
type CustomerToken struct {
	BillingAddress      *Address                  `json:"billingAddress,omitempty"`
	CompanyInformation  *CompanyInformation       `json:"companyInformation,omitempty"`
	MerchantCustomerID  *string                   `json:"merchantCustomerId,omitempty"`
	PersonalInformation *PersonalInformationToken `json:"personalInformation,omitempty"`
	// Deprecated: Use companyInformation.vatNumber instead
	VatNumber           *string                   `json:"vatNumber,omitempty"`
}

// NewCustomerToken constructs a new CustomerToken
func NewCustomerToken() *CustomerToken {
	return &CustomerToken{}
}
