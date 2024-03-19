// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// BankTransferPayoutMethodSpecificInput represents class BankTransferPayoutMethodSpecificInput
type BankTransferPayoutMethodSpecificInput struct {
	BankAccountBban *BankAccountBban `json:"bankAccountBban,omitempty"`
	BankAccountIban *BankAccountIban `json:"bankAccountIban,omitempty"`
	// Deprecated: Moved to PayoutDetails
	Customer        *PayoutCustomer  `json:"customer,omitempty"`
	PayoutDate      *string          `json:"payoutDate,omitempty"`
	PayoutText      *string          `json:"payoutText,omitempty"`
	SwiftCode       *string          `json:"swiftCode,omitempty"`
}

// NewBankTransferPayoutMethodSpecificInput constructs a new BankTransferPayoutMethodSpecificInput
func NewBankTransferPayoutMethodSpecificInput() *BankTransferPayoutMethodSpecificInput {
	return &BankTransferPayoutMethodSpecificInput{}
}
