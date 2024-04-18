// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// UpdateTokenRequest represents class UpdateTokenRequest
type UpdateTokenRequest struct {
	Card               *TokenCard                           `json:"card,omitempty"`
	EWallet            *TokenEWallet                        `json:"eWallet,omitempty"`
	NonSepaDirectDebit *TokenNonSepaDirectDebit             `json:"nonSepaDirectDebit,omitempty"`
	PaymentProductID   *int32                               `json:"paymentProductId,omitempty"`
	SepaDirectDebit    *TokenSepaDirectDebitWithoutCreditor `json:"sepaDirectDebit,omitempty"`
}

// NewUpdateTokenRequest constructs a new UpdateTokenRequest instance
func NewUpdateTokenRequest() *UpdateTokenRequest {
	return &UpdateTokenRequest{}
}
