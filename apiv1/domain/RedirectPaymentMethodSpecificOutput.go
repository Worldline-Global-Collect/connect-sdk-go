// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RedirectPaymentMethodSpecificOutput represents class RedirectPaymentMethodSpecificOutput
type RedirectPaymentMethodSpecificOutput struct {
	BankAccountBban                  *BankAccountBban                  `json:"bankAccountBban,omitempty"`
	BankAccountIban                  *BankAccountIban                  `json:"bankAccountIban,omitempty"`
	Bic                              *string                           `json:"bic,omitempty"`
	FraudResults                     *FraudResults                     `json:"fraudResults,omitempty"`
	PaymentProduct3201SpecificOutput *PaymentProduct3201SpecificOutput `json:"paymentProduct3201SpecificOutput,omitempty"`
	PaymentProduct806SpecificOutput  *PaymentProduct806SpecificOutput  `json:"paymentProduct806SpecificOutput,omitempty"`
	PaymentProduct836SpecificOutput  *PaymentProduct836SpecificOutput  `json:"paymentProduct836SpecificOutput,omitempty"`
	PaymentProduct840SpecificOutput  *PaymentProduct840SpecificOutput  `json:"paymentProduct840SpecificOutput,omitempty"`
	PaymentProduct866SpecificOutput  *PaymentProduct866SpecificOutput  `json:"paymentProduct866SpecificOutput,omitempty"`
	PaymentProductID                 *int32                            `json:"paymentProductId,omitempty"`
	Token                            *string                           `json:"token,omitempty"`
}

// NewRedirectPaymentMethodSpecificOutput constructs a new RedirectPaymentMethodSpecificOutput instance
func NewRedirectPaymentMethodSpecificOutput() *RedirectPaymentMethodSpecificOutput {
	return &RedirectPaymentMethodSpecificOutput{}
}
