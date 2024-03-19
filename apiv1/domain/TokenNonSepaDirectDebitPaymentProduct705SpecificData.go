// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// TokenNonSepaDirectDebitPaymentProduct705SpecificData represents class TokenNonSepaDirectDebitPaymentProduct705SpecificData
type TokenNonSepaDirectDebitPaymentProduct705SpecificData struct {
	AuthorisationID *string          `json:"authorisationId,omitempty"`
	BankAccountBban *BankAccountBban `json:"bankAccountBban,omitempty"`
}

// NewTokenNonSepaDirectDebitPaymentProduct705SpecificData constructs a new TokenNonSepaDirectDebitPaymentProduct705SpecificData
func NewTokenNonSepaDirectDebitPaymentProduct705SpecificData() *TokenNonSepaDirectDebitPaymentProduct705SpecificData {
	return &TokenNonSepaDirectDebitPaymentProduct705SpecificData{}
}
