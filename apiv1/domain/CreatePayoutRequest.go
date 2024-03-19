// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CreatePayoutRequest represents class CreatePayoutRequest
type CreatePayoutRequest struct {
	// Deprecated: Moved to PayoutDetails
	AmountOfMoney                         *AmountOfMoney                         `json:"amountOfMoney,omitempty"`
	// Deprecated: Moved to BankTransferPayoutMethodSpecificInput
	BankAccountBban                       *BankAccountBban                       `json:"bankAccountBban,omitempty"`
	// Deprecated: Moved to BankTransferPayoutMethodSpecificInput
	BankAccountIban                       *BankAccountIban                       `json:"bankAccountIban,omitempty"`
	BankTransferPayoutMethodSpecificInput *BankTransferPayoutMethodSpecificInput `json:"bankTransferPayoutMethodSpecificInput,omitempty"`
	CardPayoutMethodSpecificInput         *CardPayoutMethodSpecificInput         `json:"cardPayoutMethodSpecificInput,omitempty"`
	// Deprecated: Moved to PayoutDetails
	Customer                              *PayoutCustomer                        `json:"customer,omitempty"`
	Merchant                              *PayoutMerchant                        `json:"merchant,omitempty"`
	// Deprecated: Moved to BankTransferPayoutMethodSpecificInput
	PayoutDate                            *string                                `json:"payoutDate,omitempty"`
	PayoutDetails                         *PayoutDetails                         `json:"payoutDetails,omitempty"`
	// Deprecated: Moved to BankTransferPayoutMethodSpecificInput
	PayoutText                            *string                                `json:"payoutText,omitempty"`
	// Deprecated: Moved to PayoutDetails
	References                            *PayoutReferences                      `json:"references,omitempty"`
	// Deprecated: Moved to BankTransferPayoutMethodSpecificInput
	SwiftCode                             *string                                `json:"swiftCode,omitempty"`
}

// NewCreatePayoutRequest constructs a new CreatePayoutRequest
func NewCreatePayoutRequest() *CreatePayoutRequest {
	return &CreatePayoutRequest{}
}
