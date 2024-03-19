// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ContactDetailsToken represents class ContactDetailsToken
type ContactDetailsToken struct {
	EmailAddress     *string `json:"emailAddress,omitempty"`
	EmailMessageType *string `json:"emailMessageType,omitempty"`
}

// NewContactDetailsToken constructs a new ContactDetailsToken
func NewContactDetailsToken() *ContactDetailsToken {
	return &ContactDetailsToken{}
}
