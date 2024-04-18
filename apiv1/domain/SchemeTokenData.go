// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// SchemeTokenData represents class SchemeTokenData
type SchemeTokenData struct {
	CardholderName  *string `json:"cardholderName,omitempty"`
	Cryptogram      *string `json:"cryptogram,omitempty"`
	Eci             *string `json:"eci,omitempty"`
	NetworkToken    *string `json:"networkToken,omitempty"`
	TokenExpiryDate *string `json:"tokenExpiryDate,omitempty"`
}

// NewSchemeTokenData constructs a new SchemeTokenData instance
func NewSchemeTokenData() *SchemeTokenData {
	return &SchemeTokenData{}
}
