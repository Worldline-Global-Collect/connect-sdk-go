// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// Creditor represents class Creditor
type Creditor struct {
	AdditionalAddressInfo *string `json:"additionalAddressInfo,omitempty"`
	City                  *string `json:"city,omitempty"`
	CountryCode           *string `json:"countryCode,omitempty"`
	HouseNumber           *string `json:"houseNumber,omitempty"`
	Iban                  *string `json:"iban,omitempty"`
	ID                    *string `json:"id,omitempty"`
	Name                  *string `json:"name,omitempty"`
	ReferenceParty        *string `json:"referenceParty,omitempty"`
	ReferencePartyID      *string `json:"referencePartyId,omitempty"`
	Street                *string `json:"street,omitempty"`
	Zip                   *string `json:"zip,omitempty"`
}

// NewCreditor constructs a new Creditor instance
func NewCreditor() *Creditor {
	return &Creditor{}
}
