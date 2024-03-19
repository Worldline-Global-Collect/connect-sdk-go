// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RefundsResponse represents class RefundsResponse
type RefundsResponse struct {
	Refunds *[]RefundResult `json:"refunds,omitempty"`
}

// NewRefundsResponse constructs a new RefundsResponse
func NewRefundsResponse() *RefundsResponse {
	return &RefundsResponse{}
}
