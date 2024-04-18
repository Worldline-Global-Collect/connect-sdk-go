// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MandateCustomer represents class MandateCustomer
type MandateCustomer struct {
	BankAccountIban     *BankAccountIban            `json:"bankAccountIban,omitempty"`
	CompanyName         *string                     `json:"companyName,omitempty"`
	ContactDetails      *MandateContactDetails      `json:"contactDetails,omitempty"`
	MandateAddress      *MandateAddress             `json:"mandateAddress,omitempty"`
	PersonalInformation *MandatePersonalInformation `json:"personalInformation,omitempty"`
}

// NewMandateCustomer constructs a new MandateCustomer instance
func NewMandateCustomer() *MandateCustomer {
	return &MandateCustomer{}
}
