// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentContext represents class PaymentContext
type PaymentContext struct {
	AmountOfMoney  *AmountOfMoney `json:"amountOfMoney,omitempty"`
	CountryCode    *string        `json:"countryCode,omitempty"`
	IsInstallments *bool          `json:"isInstallments,omitempty"`
	IsRecurring    *bool          `json:"isRecurring,omitempty"`
}

// NewPaymentContext constructs a new PaymentContext instance
func NewPaymentContext() *PaymentContext {
	return &PaymentContext{}
}
