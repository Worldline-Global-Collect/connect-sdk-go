// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// TokenResponse represents class TokenResponse
type TokenResponse struct {
	Card               *TokenCard               `json:"card,omitempty"`
	EWallet            *TokenEWallet            `json:"eWallet,omitempty"`
	ID                 *string                  `json:"id,omitempty"`
	NonSepaDirectDebit *TokenNonSepaDirectDebit `json:"nonSepaDirectDebit,omitempty"`
	OriginalPaymentID  *string                  `json:"originalPaymentId,omitempty"`
	PaymentProductID   *int32                   `json:"paymentProductId,omitempty"`
	SepaDirectDebit    *TokenSepaDirectDebit    `json:"sepaDirectDebit,omitempty"`
}

// NewTokenResponse constructs a new TokenResponse
func NewTokenResponse() *TokenResponse {
	return &TokenResponse{}
}
