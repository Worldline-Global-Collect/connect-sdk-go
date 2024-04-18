// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// FindPaymentsResponse represents class FindPaymentsResponse
type FindPaymentsResponse struct {
	Limit      *int32     `json:"limit,omitempty"`
	Offset     *int32     `json:"offset,omitempty"`
	Payments   *[]Payment `json:"payments,omitempty"`
	TotalCount *int32     `json:"totalCount,omitempty"`
}

// NewFindPaymentsResponse constructs a new FindPaymentsResponse instance
func NewFindPaymentsResponse() *FindPaymentsResponse {
	return &FindPaymentsResponse{}
}
