// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// Merchant represents class Merchant
type Merchant struct {
	ConfigurationID   *string `json:"configurationId,omitempty"`
	ContactWebsiteURL *string `json:"contactWebsiteUrl,omitempty"`
	Seller            *Seller `json:"seller,omitempty"`
	WebsiteURL        *string `json:"websiteUrl,omitempty"`
}

// NewMerchant constructs a new Merchant instance
func NewMerchant() *Merchant {
	return &Merchant{}
}
