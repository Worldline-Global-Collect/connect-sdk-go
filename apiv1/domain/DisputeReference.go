// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// DisputeReference represents class DisputeReference
type DisputeReference struct {
	MerchantOrderID   *string `json:"merchantOrderId,omitempty"`
	MerchantReference *string `json:"merchantReference,omitempty"`
	PaymentReference  *string `json:"paymentReference,omitempty"`
	ProviderID        *string `json:"providerId,omitempty"`
	ProviderReference *string `json:"providerReference,omitempty"`
}

// NewDisputeReference constructs a new DisputeReference instance
func NewDisputeReference() *DisputeReference {
	return &DisputeReference{}
}
