// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// AddressPersonal represents class AddressPersonal
type AddressPersonal struct {
	AdditionalInfo *string       `json:"additionalInfo,omitempty"`
	City           *string       `json:"city,omitempty"`
	CountryCode    *string       `json:"countryCode,omitempty"`
	HouseNumber    *string       `json:"houseNumber,omitempty"`
	Name           *PersonalName `json:"name,omitempty"`
	State          *string       `json:"state,omitempty"`
	StateCode      *string       `json:"stateCode,omitempty"`
	Street         *string       `json:"street,omitempty"`
	Zip            *string       `json:"zip,omitempty"`
}

// NewAddressPersonal constructs a new AddressPersonal instance
func NewAddressPersonal() *AddressPersonal {
	return &AddressPersonal{}
}
