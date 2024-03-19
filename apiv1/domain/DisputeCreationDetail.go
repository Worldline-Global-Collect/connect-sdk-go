// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// DisputeCreationDetail represents class DisputeCreationDetail
type DisputeCreationDetail struct {
	DisputeCreationDate *string `json:"disputeCreationDate,omitempty"`
	DisputeOriginator   *string `json:"disputeOriginator,omitempty"`
	UserName            *string `json:"userName,omitempty"`
}

// NewDisputeCreationDetail constructs a new DisputeCreationDetail
func NewDisputeCreationDetail() *DisputeCreationDetail {
	return &DisputeCreationDetail{}
}
