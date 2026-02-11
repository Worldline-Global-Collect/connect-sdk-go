// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProductGroup represents class PaymentProductGroup
type PaymentProductGroup struct {
	AccountsOnFile           *[]AccountOnFile            `json:"accountsOnFile,omitempty"`
	AllowsClickToPay         *bool                       `json:"allowsClickToPay,omitempty"`
	AllowsInstallments       *bool                       `json:"allowsInstallments,omitempty"`
	ClickToPayConfiguration  *ClickToPayConfiguration    `json:"clickToPayConfiguration,omitempty"`
	DeviceFingerprintEnabled *bool                       `json:"deviceFingerprintEnabled,omitempty"`
	DisplayHints             *PaymentProductDisplayHints `json:"displayHints,omitempty"`
	Fields                   *[]PaymentProductField      `json:"fields,omitempty"`
	ID                       *string                     `json:"id,omitempty"`
}

// NewPaymentProductGroup constructs a new PaymentProductGroup instance
func NewPaymentProductGroup() *PaymentProductGroup {
	return &PaymentProductGroup{}
}
