// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RedirectData represents class RedirectData
type RedirectData struct {
	RETURNMAC   *string `json:"RETURNMAC,omitempty"`
	RedirectURL *string `json:"redirectURL,omitempty"`
}

// NewRedirectData constructs a new RedirectData
func NewRedirectData() *RedirectData {
	return &RedirectData{}
}
