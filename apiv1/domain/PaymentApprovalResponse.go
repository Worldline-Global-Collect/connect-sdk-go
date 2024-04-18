// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentApprovalResponse represents class PaymentApprovalResponse
type PaymentApprovalResponse struct {
	CardPaymentMethodSpecificOutput   *ApprovePaymentCardPaymentMethodSpecificOutput   `json:"cardPaymentMethodSpecificOutput,omitempty"`
	MobilePaymentMethodSpecificOutput *ApprovePaymentMobilePaymentMethodSpecificOutput `json:"mobilePaymentMethodSpecificOutput,omitempty"`
	Payment                           *Payment                                         `json:"payment,omitempty"`
	// Deprecated: Use cardPaymentMethodSpecificOutput instead
	PaymentMethodSpecificOutput       *ApprovePaymentCardPaymentMethodSpecificOutput   `json:"paymentMethodSpecificOutput,omitempty"`
}

// NewPaymentApprovalResponse constructs a new PaymentApprovalResponse instance
func NewPaymentApprovalResponse() *PaymentApprovalResponse {
	return &PaymentApprovalResponse{}
}
