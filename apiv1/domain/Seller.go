// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// Seller represents class Seller
type Seller struct {
	Address             *Address `json:"address,omitempty"`
	ChannelCode         *string  `json:"channelCode,omitempty"`
	Description         *string  `json:"description,omitempty"`
	ExternalReferenceID *string  `json:"externalReferenceId,omitempty"`
	Geocode             *string  `json:"geocode,omitempty"`
	ID                  *string  `json:"id,omitempty"`
	InvoiceNumber       *string  `json:"invoiceNumber,omitempty"`
	IsForeignRetailer   *bool    `json:"isForeignRetailer,omitempty"`
	Mcc                 *string  `json:"mcc,omitempty"`
	Name                *string  `json:"name,omitempty"`
	PhoneNumber         *string  `json:"phoneNumber,omitempty"`
	Type                *string  `json:"type,omitempty"`
}

// NewSeller constructs a new Seller
func NewSeller() *Seller {
	return &Seller{}
}
