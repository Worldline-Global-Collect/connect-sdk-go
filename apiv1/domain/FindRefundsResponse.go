// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// FindRefundsResponse represents class FindRefundsResponse
type FindRefundsResponse struct {
	Limit      *int32          `json:"limit,omitempty"`
	Offset     *int32          `json:"offset,omitempty"`
	Refunds    *[]RefundResult `json:"refunds,omitempty"`
	TotalCount *int32          `json:"totalCount,omitempty"`
}

// NewFindRefundsResponse constructs a new FindRefundsResponse
func NewFindRefundsResponse() *FindRefundsResponse {
	return &FindRefundsResponse{}
}
