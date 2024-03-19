// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// NonSepaDirectDebitPaymentProduct705SpecificInput represents class NonSepaDirectDebitPaymentProduct705SpecificInput
type NonSepaDirectDebitPaymentProduct705SpecificInput struct {
	AuthorisationID *string          `json:"authorisationId,omitempty"`
	BankAccountBban *BankAccountBban `json:"bankAccountBban,omitempty"`
	TransactionType *string          `json:"transactionType,omitempty"`
}

// NewNonSepaDirectDebitPaymentProduct705SpecificInput constructs a new NonSepaDirectDebitPaymentProduct705SpecificInput
func NewNonSepaDirectDebitPaymentProduct705SpecificInput() *NonSepaDirectDebitPaymentProduct705SpecificInput {
	return &NonSepaDirectDebitPaymentProduct705SpecificInput{}
}
