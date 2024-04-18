// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PayoutResult represents class PayoutResult
type PayoutResult struct {
	ID           *string            `json:"id,omitempty"`
	PayoutOutput *OrderOutput       `json:"payoutOutput,omitempty"`
	Status       *string            `json:"status,omitempty"`
	StatusOutput *OrderStatusOutput `json:"statusOutput,omitempty"`
}

// NewPayoutResult constructs a new PayoutResult instance
func NewPayoutResult() *PayoutResult {
	return &PayoutResult{}
}
