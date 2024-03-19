// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RedirectionData represents class RedirectionData
type RedirectionData struct {
	ReturnURL *string `json:"returnUrl,omitempty"`
	Variant   *string `json:"variant,omitempty"`
}

// NewRedirectionData constructs a new RedirectionData
func NewRedirectionData() *RedirectionData {
	return &RedirectionData{}
}
