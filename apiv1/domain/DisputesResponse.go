// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// DisputesResponse represents class DisputesResponse
type DisputesResponse struct {
	Disputes *[]Dispute `json:"disputes,omitempty"`
}

// NewDisputesResponse constructs a new DisputesResponse
func NewDisputesResponse() *DisputesResponse {
	return &DisputesResponse{}
}
