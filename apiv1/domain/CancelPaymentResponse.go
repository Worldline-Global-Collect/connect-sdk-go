// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CancelPaymentResponse represents class CancelPaymentResponse
type CancelPaymentResponse struct {
	CardPaymentMethodSpecificOutput   *CancelPaymentCardPaymentMethodSpecificOutput   `json:"cardPaymentMethodSpecificOutput,omitempty"`
	MobilePaymentMethodSpecificOutput *CancelPaymentMobilePaymentMethodSpecificOutput `json:"mobilePaymentMethodSpecificOutput,omitempty"`
	Payment                           *Payment                                        `json:"payment,omitempty"`
}

// NewCancelPaymentResponse constructs a new CancelPaymentResponse
func NewCancelPaymentResponse() *CancelPaymentResponse {
	return &CancelPaymentResponse{}
}
