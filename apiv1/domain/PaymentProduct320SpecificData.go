// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProduct320SpecificData represents class PaymentProduct320SpecificData
type PaymentProduct320SpecificData struct {
	Gateway  *string   `json:"gateway,omitempty"`
	Networks *[]string `json:"networks,omitempty"`
}

// NewPaymentProduct320SpecificData constructs a new PaymentProduct320SpecificData instance
func NewPaymentProduct320SpecificData() *PaymentProduct320SpecificData {
	return &PaymentProduct320SpecificData{}
}
