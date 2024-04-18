// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PayoutCustomer represents class PayoutCustomer
type PayoutCustomer struct {
	Address            *Address            `json:"address,omitempty"`
	CompanyInformation *CompanyInformation `json:"companyInformation,omitempty"`
	ContactDetails     *ContactDetailsBase `json:"contactDetails,omitempty"`
	MerchantCustomerID *string             `json:"merchantCustomerId,omitempty"`
	Name               *PersonalName       `json:"name,omitempty"`
}

// NewPayoutCustomer constructs a new PayoutCustomer instance
func NewPayoutCustomer() *PayoutCustomer {
	return &PayoutCustomer{}
}
