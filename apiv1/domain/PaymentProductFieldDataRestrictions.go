// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProductFieldDataRestrictions represents class PaymentProductFieldDataRestrictions
type PaymentProductFieldDataRestrictions struct {
	IsRequired *bool                          `json:"isRequired,omitempty"`
	Validators *PaymentProductFieldValidators `json:"validators,omitempty"`
}

// NewPaymentProductFieldDataRestrictions constructs a new PaymentProductFieldDataRestrictions
func NewPaymentProductFieldDataRestrictions() *PaymentProductFieldDataRestrictions {
	return &PaymentProductFieldDataRestrictions{}
}
