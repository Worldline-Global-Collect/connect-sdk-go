// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// AfrName represents class AfrName
type AfrName struct {
	FirstName *string `json:"firstName,omitempty"`
	Surname   *string `json:"surname,omitempty"`
}

// NewAfrName constructs a new AfrName instance
func NewAfrName() *AfrName {
	return &AfrName{}
}
