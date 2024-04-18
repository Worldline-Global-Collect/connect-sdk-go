// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MandateRedirectData represents class MandateRedirectData
type MandateRedirectData struct {
	RETURNMAC   *string `json:"RETURNMAC,omitempty"`
	RedirectURL *string `json:"redirectURL,omitempty"`
}

// NewMandateRedirectData constructs a new MandateRedirectData instance
func NewMandateRedirectData() *MandateRedirectData {
	return &MandateRedirectData{}
}
