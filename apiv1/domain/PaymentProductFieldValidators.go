// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProductFieldValidators represents class PaymentProductFieldValidators
type PaymentProductFieldValidators struct {
	BoletoBancarioRequiredness *BoletoBancarioRequirednessValidator `json:"boletoBancarioRequiredness,omitempty"`
	EmailAddress               *EmptyValidator                      `json:"emailAddress,omitempty"`
	ExpirationDate             *EmptyValidator                      `json:"expirationDate,omitempty"`
	FixedList                  *FixedListValidator                  `json:"fixedList,omitempty"`
	Iban                       *EmptyValidator                      `json:"iban,omitempty"`
	Length                     *LengthValidator                     `json:"length,omitempty"`
	Luhn                       *EmptyValidator                      `json:"luhn,omitempty"`
	Range                      *RangeValidator                      `json:"range,omitempty"`
	RegularExpression          *RegularExpressionValidator          `json:"regularExpression,omitempty"`
	ResidentIDNumber           *EmptyValidator                      `json:"residentIdNumber,omitempty"`
	TermsAndConditions         *EmptyValidator                      `json:"termsAndConditions,omitempty"`
}

// NewPaymentProductFieldValidators constructs a new PaymentProductFieldValidators instance
func NewPaymentProductFieldValidators() *PaymentProductFieldValidators {
	return &PaymentProductFieldValidators{}
}
