// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// FindPayoutsResponse represents class FindPayoutsResponse
type FindPayoutsResponse struct {
	Limit      *int32          `json:"limit,omitempty"`
	Offset     *int32          `json:"offset,omitempty"`
	Payouts    *[]PayoutResult `json:"payouts,omitempty"`
	TotalCount *int32          `json:"totalCount,omitempty"`
}

// NewFindPayoutsResponse constructs a new FindPayoutsResponse instance
func NewFindPayoutsResponse() *FindPayoutsResponse {
	return &FindPayoutsResponse{}
}
