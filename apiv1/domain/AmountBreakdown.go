// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// AmountBreakdown represents class AmountBreakdown
type AmountBreakdown struct {
	Amount *int64  `json:"amount,omitempty"`
	Type   *string `json:"type,omitempty"`
}

// NewAmountBreakdown constructs a new AmountBreakdown instance
func NewAmountBreakdown() *AmountBreakdown {
	return &AmountBreakdown{}
}
