// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// Address represents class Address
type Address struct {
	AdditionalInfo *string `json:"additionalInfo,omitempty"`
	City           *string `json:"city,omitempty"`
	CountryCode    *string `json:"countryCode,omitempty"`
	HouseNumber    *string `json:"houseNumber,omitempty"`
	State          *string `json:"state,omitempty"`
	StateCode      *string `json:"stateCode,omitempty"`
	Street         *string `json:"street,omitempty"`
	Zip            *string `json:"zip,omitempty"`
}

// NewAddress constructs a new Address instance
func NewAddress() *Address {
	return &Address{}
}
