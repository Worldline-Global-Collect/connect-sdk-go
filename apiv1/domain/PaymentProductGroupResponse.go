// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProductGroupResponse represents class PaymentProductGroupResponse
type PaymentProductGroupResponse struct {
	AccountsOnFile           *[]AccountOnFile            `json:"accountsOnFile,omitempty"`
	AllowsInstallments       *bool                       `json:"allowsInstallments,omitempty"`
	DeviceFingerprintEnabled *bool                       `json:"deviceFingerprintEnabled,omitempty"`
	DisplayHints             *PaymentProductDisplayHints `json:"displayHints,omitempty"`
	Fields                   *[]PaymentProductField      `json:"fields,omitempty"`
	ID                       *string                     `json:"id,omitempty"`
}

// NewPaymentProductGroupResponse constructs a new PaymentProductGroupResponse instance
func NewPaymentProductGroupResponse() *PaymentProductGroupResponse {
	return &PaymentProductGroupResponse{}
}
