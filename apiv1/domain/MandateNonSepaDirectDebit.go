// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MandateNonSepaDirectDebit represents class MandateNonSepaDirectDebit
type MandateNonSepaDirectDebit struct {
	PaymentProduct705SpecificData *TokenNonSepaDirectDebitPaymentProduct705SpecificData `json:"paymentProduct705SpecificData,omitempty"`
	PaymentProduct730SpecificData *TokenNonSepaDirectDebitPaymentProduct730SpecificData `json:"paymentProduct730SpecificData,omitempty"`
}

// NewMandateNonSepaDirectDebit constructs a new MandateNonSepaDirectDebit
func NewMandateNonSepaDirectDebit() *MandateNonSepaDirectDebit {
	return &MandateNonSepaDirectDebit{}
}
