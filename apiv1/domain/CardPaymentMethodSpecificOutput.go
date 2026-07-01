// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CardPaymentMethodSpecificOutput represents class CardPaymentMethodSpecificOutput.
type CardPaymentMethodSpecificOutput struct {
	AuthorisationCode          *string              `json:"authorisationCode,omitempty"`
	Card                       *CardEssentials      `json:"card,omitempty"`
	ClickToPayUsed             *bool                `json:"clickToPayUsed,omitempty"`
	FraudResults               *CardFraudResults    `json:"fraudResults,omitempty"`
	InitialSchemeTransactionID *string              `json:"initialSchemeTransactionId,omitempty"`
	NetworkTokenData           *NetworkTokenData    `json:"networkTokenData,omitempty"`
	NetworkTokenUsed           *bool                `json:"networkTokenUsed,omitempty"`
	OriginalTransactionLinkID  *string              `json:"originalTransactionLinkId,omitempty"`
	PaymentAccountReference    *string              `json:"paymentAccountReference,omitempty"`
	PaymentProductID           *int32               `json:"paymentProductId,omitempty"`
	SchemeTransactionID        *string              `json:"schemeTransactionId,omitempty"`
	ThreeDSecureResults        *ThreeDSecureResults `json:"threeDSecureResults,omitempty"`
	Token                      *string              `json:"token,omitempty"`
	TransactionLinkID          *string              `json:"transactionLinkId,omitempty"`
}

// NewCardPaymentMethodSpecificOutput constructs a new CardPaymentMethodSpecificOutput instance.
func NewCardPaymentMethodSpecificOutput() *CardPaymentMethodSpecificOutput {
	return &CardPaymentMethodSpecificOutput{}
}
