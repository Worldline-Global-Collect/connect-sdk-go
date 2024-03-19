// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentCreationOutput represents class PaymentCreationOutput
type PaymentCreationOutput struct {
	AdditionalReference   *string `json:"additionalReference,omitempty"`
	ExternalReference     *string `json:"externalReference,omitempty"`
	IsCheckedRememberMe   *bool   `json:"isCheckedRememberMe,omitempty"`
	IsNewToken            *bool   `json:"isNewToken,omitempty"`
	Token                 *string `json:"token,omitempty"`
	TokenizationSucceeded *bool   `json:"tokenizationSucceeded,omitempty"`
}

// NewPaymentCreationOutput constructs a new PaymentCreationOutput
func NewPaymentCreationOutput() *PaymentCreationOutput {
	return &PaymentCreationOutput{}
}
