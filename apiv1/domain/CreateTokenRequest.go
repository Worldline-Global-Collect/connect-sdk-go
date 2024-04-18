// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CreateTokenRequest represents class CreateTokenRequest
type CreateTokenRequest struct {
	Card                   *TokenCard                           `json:"card,omitempty"`
	EWallet                *TokenEWallet                        `json:"eWallet,omitempty"`
	EncryptedCustomerInput *string                              `json:"encryptedCustomerInput,omitempty"`
	NonSepaDirectDebit     *TokenNonSepaDirectDebit             `json:"nonSepaDirectDebit,omitempty"`
	PaymentProductID       *int32                               `json:"paymentProductId,omitempty"`
	SepaDirectDebit        *TokenSepaDirectDebitWithoutCreditor `json:"sepaDirectDebit,omitempty"`
}

// NewCreateTokenRequest constructs a new CreateTokenRequest instance
func NewCreateTokenRequest() *CreateTokenRequest {
	return &CreateTokenRequest{}
}
