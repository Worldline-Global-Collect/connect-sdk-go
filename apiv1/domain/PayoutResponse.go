// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PayoutResponse represents class PayoutResponse
type PayoutResponse struct {
	ID           *string            `json:"id,omitempty"`
	PayoutOutput *OrderOutput       `json:"payoutOutput,omitempty"`
	Status       *string            `json:"status,omitempty"`
	StatusOutput *OrderStatusOutput `json:"statusOutput,omitempty"`
}

// NewPayoutResponse constructs a new PayoutResponse instance
func NewPayoutResponse() *PayoutResponse {
	return &PayoutResponse{}
}
