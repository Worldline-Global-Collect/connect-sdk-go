// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RedirectPaymentMethodSpecificInputBase represents class RedirectPaymentMethodSpecificInputBase
type RedirectPaymentMethodSpecificInputBase struct {
	ExpirationPeriod                  *int32                                       `json:"expirationPeriod,omitempty"`
	PaymentProduct4101SpecificInput   *RedirectPaymentProduct4101SpecificInputBase `json:"paymentProduct4101SpecificInput,omitempty"`
	PaymentProduct840SpecificInput    *RedirectPaymentProduct840SpecificInputBase  `json:"paymentProduct840SpecificInput,omitempty"`
	PaymentProductID                  *int32                                       `json:"paymentProductId,omitempty"`
	RecurringPaymentSequenceIndicator *string                                      `json:"recurringPaymentSequenceIndicator,omitempty"`
	RequiresApproval                  *bool                                        `json:"requiresApproval,omitempty"`
	Token                             *string                                      `json:"token,omitempty"`
	Tokenize                          *bool                                        `json:"tokenize,omitempty"`
}

// NewRedirectPaymentMethodSpecificInputBase constructs a new RedirectPaymentMethodSpecificInputBase instance
func NewRedirectPaymentMethodSpecificInputBase() *RedirectPaymentMethodSpecificInputBase {
	return &RedirectPaymentMethodSpecificInputBase{}
}
