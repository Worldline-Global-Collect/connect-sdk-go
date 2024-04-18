// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CreateHostedCheckoutRequest represents class CreateHostedCheckoutRequest
type CreateHostedCheckoutRequest struct {
	BankTransferPaymentMethodSpecificInput    *BankTransferPaymentMethodSpecificInputBase     `json:"bankTransferPaymentMethodSpecificInput,omitempty"`
	CardPaymentMethodSpecificInput            *CardPaymentMethodSpecificInputBase             `json:"cardPaymentMethodSpecificInput,omitempty"`
	CashPaymentMethodSpecificInput            *CashPaymentMethodSpecificInputBase             `json:"cashPaymentMethodSpecificInput,omitempty"`
	EInvoicePaymentMethodSpecificInput        *EInvoicePaymentMethodSpecificInputBase         `json:"eInvoicePaymentMethodSpecificInput,omitempty"`
	FraudFields                               *FraudFields                                    `json:"fraudFields,omitempty"`
	HostedCheckoutSpecificInput               *HostedCheckoutSpecificInput                    `json:"hostedCheckoutSpecificInput,omitempty"`
	Merchant                                  *Merchant                                       `json:"merchant,omitempty"`
	MobilePaymentMethodSpecificInput          *MobilePaymentMethodSpecificInputHostedCheckout `json:"mobilePaymentMethodSpecificInput,omitempty"`
	Order                                     *Order                                          `json:"order,omitempty"`
	RedirectPaymentMethodSpecificInput        *RedirectPaymentMethodSpecificInputBase         `json:"redirectPaymentMethodSpecificInput,omitempty"`
	SepaDirectDebitPaymentMethodSpecificInput *SepaDirectDebitPaymentMethodSpecificInputBase  `json:"sepaDirectDebitPaymentMethodSpecificInput,omitempty"`
}

// NewCreateHostedCheckoutRequest constructs a new CreateHostedCheckoutRequest instance
func NewCreateHostedCheckoutRequest() *CreateHostedCheckoutRequest {
	return &CreateHostedCheckoutRequest{}
}
