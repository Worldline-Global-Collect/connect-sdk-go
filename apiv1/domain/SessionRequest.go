// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// SessionRequest represents class SessionRequest
type SessionRequest struct {
	PaymentProductFilters *PaymentProductFiltersClientSession `json:"paymentProductFilters,omitempty"`
	Tokens                *[]string                           `json:"tokens,omitempty"`
}

// NewSessionRequest constructs a new SessionRequest instance
func NewSessionRequest() *SessionRequest {
	return &SessionRequest{}
}
