// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProductFilter represents class PaymentProductFilter
type PaymentProductFilter struct {
	Groups   *[]string `json:"groups,omitempty"`
	Products *[]int32  `json:"products,omitempty"`
}

// NewPaymentProductFilter constructs a new PaymentProductFilter instance
func NewPaymentProductFilter() *PaymentProductFilter {
	return &PaymentProductFilter{}
}
