// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RefundCustomer represents class RefundCustomer
type RefundCustomer struct {
	Address            *AddressPersonal    `json:"address,omitempty"`
	CompanyInformation *CompanyInformation `json:"companyInformation,omitempty"`
	ContactDetails     *ContactDetailsBase `json:"contactDetails,omitempty"`
	FiscalNumber       *string             `json:"fiscalNumber,omitempty"`
}

// NewRefundCustomer constructs a new RefundCustomer
func NewRefundCustomer() *RefundCustomer {
	return &RefundCustomer{}
}
