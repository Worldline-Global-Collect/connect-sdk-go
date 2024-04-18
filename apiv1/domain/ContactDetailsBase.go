// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ContactDetailsBase represents class ContactDetailsBase
type ContactDetailsBase struct {
	EmailAddress     *string `json:"emailAddress,omitempty"`
	EmailMessageType *string `json:"emailMessageType,omitempty"`
}

// NewContactDetailsBase constructs a new ContactDetailsBase instance
func NewContactDetailsBase() *ContactDetailsBase {
	return &ContactDetailsBase{}
}
