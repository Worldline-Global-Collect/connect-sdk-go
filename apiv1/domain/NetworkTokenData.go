// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// NetworkTokenData represents class NetworkTokenData
type NetworkTokenData struct {
	NetworkToken     *string `json:"networkToken,omitempty"`
	TokenExpiryDate  *string `json:"tokenExpiryDate,omitempty"`
	TokenReferenceID *string `json:"tokenReferenceId,omitempty"`
}

// NewNetworkTokenData constructs a new NetworkTokenData instance
func NewNetworkTokenData() *NetworkTokenData {
	return &NetworkTokenData{}
}
