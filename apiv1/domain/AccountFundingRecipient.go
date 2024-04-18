// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// AccountFundingRecipient represents class AccountFundingRecipient
type AccountFundingRecipient struct {
	AccountNumber     *string  `json:"accountNumber,omitempty"`
	AccountNumberType *string  `json:"accountNumberType,omitempty"`
	Address           *Address `json:"address,omitempty"`
	DateOfBirth       *string  `json:"dateOfBirth,omitempty"`
	Name              *AfrName `json:"name,omitempty"`
	PartialPan        *string  `json:"partialPan,omitempty"`
}

// NewAccountFundingRecipient constructs a new AccountFundingRecipient instance
func NewAccountFundingRecipient() *AccountFundingRecipient {
	return &AccountFundingRecipient{}
}
