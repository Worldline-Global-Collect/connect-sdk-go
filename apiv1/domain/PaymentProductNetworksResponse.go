// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProductNetworksResponse represents class PaymentProductNetworksResponse
type PaymentProductNetworksResponse struct {
	Networks *[]string `json:"networks,omitempty"`
}

// NewPaymentProductNetworksResponse constructs a new PaymentProductNetworksResponse
func NewPaymentProductNetworksResponse() *PaymentProductNetworksResponse {
	return &PaymentProductNetworksResponse{}
}
