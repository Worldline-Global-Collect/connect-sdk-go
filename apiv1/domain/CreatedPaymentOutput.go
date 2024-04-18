// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CreatedPaymentOutput represents class CreatedPaymentOutput
type CreatedPaymentOutput struct {
	DisplayedData             *DisplayedData             `json:"displayedData,omitempty"`
	IsCheckedRememberMe       *bool                      `json:"isCheckedRememberMe,omitempty"`
	Payment                   *Payment                   `json:"payment,omitempty"`
	PaymentCreationReferences *PaymentCreationReferences `json:"paymentCreationReferences,omitempty"`
	// Deprecated: Use Payment.statusOutput.statusCategory instead
	PaymentStatusCategory     *string                    `json:"paymentStatusCategory,omitempty"`
	TokenizationSucceeded     *bool                      `json:"tokenizationSucceeded,omitempty"`
	Tokens                    *string                    `json:"tokens,omitempty"`
}

// NewCreatedPaymentOutput constructs a new CreatedPaymentOutput instance
func NewCreatedPaymentOutput() *CreatedPaymentOutput {
	return &CreatedPaymentOutput{}
}
