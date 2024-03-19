// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ApprovePaymentRequest represents class ApprovePaymentRequest
type ApprovePaymentRequest struct {
	Amount                                    *int64                                                      `json:"amount,omitempty"`
	DirectDebitPaymentMethodSpecificInput     *ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput `json:"directDebitPaymentMethodSpecificInput,omitempty"`
	Order                                     *OrderApprovePayment                                        `json:"order,omitempty"`
	SepaDirectDebitPaymentMethodSpecificInput *ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput    `json:"sepaDirectDebitPaymentMethodSpecificInput,omitempty"`
}

// NewApprovePaymentRequest constructs a new ApprovePaymentRequest
func NewApprovePaymentRequest() *ApprovePaymentRequest {
	return &ApprovePaymentRequest{}
}
