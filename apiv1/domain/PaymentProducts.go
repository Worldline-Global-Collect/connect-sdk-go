// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProducts represents class PaymentProducts
type PaymentProducts struct {
	PaymentProducts *[]PaymentProduct `json:"paymentProducts,omitempty"`
}

// NewPaymentProducts constructs a new PaymentProducts
func NewPaymentProducts() *PaymentProducts {
	return &PaymentProducts{}
}
