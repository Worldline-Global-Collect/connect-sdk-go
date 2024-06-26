// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// SessionResponse represents class SessionResponse
type SessionResponse struct {
	AssetURL        *string   `json:"assetUrl,omitempty"`
	ClientAPIURL    *string   `json:"clientApiUrl,omitempty"`
	ClientSessionID *string   `json:"clientSessionId,omitempty"`
	CustomerID      *string   `json:"customerId,omitempty"`
	InvalidTokens   *[]string `json:"invalidTokens,omitempty"`
	Region          *string   `json:"region,omitempty"`
}

// NewSessionResponse constructs a new SessionResponse instance
func NewSessionResponse() *SessionResponse {
	return &SessionResponse{}
}
