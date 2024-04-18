// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// AuthenticationIndicator represents class AuthenticationIndicator
type AuthenticationIndicator struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewAuthenticationIndicator constructs a new AuthenticationIndicator instance
func NewAuthenticationIndicator() *AuthenticationIndicator {
	return &AuthenticationIndicator{}
}
