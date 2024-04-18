// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PayoutErrorResponse represents class PayoutErrorResponse
type PayoutErrorResponse struct {
	ErrorID      *string       `json:"errorId,omitempty"`
	Errors       *[]APIError   `json:"errors,omitempty"`
	PayoutResult *PayoutResult `json:"payoutResult,omitempty"`
}

// NewPayoutErrorResponse constructs a new PayoutErrorResponse instance
func NewPayoutErrorResponse() *PayoutErrorResponse {
	return &PayoutErrorResponse{}
}
