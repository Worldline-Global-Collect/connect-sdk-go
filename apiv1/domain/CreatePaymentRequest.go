// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CreatePaymentRequest represents class CreatePaymentRequest
type CreatePaymentRequest struct {
	BankTransferPaymentMethodSpecificInput    *BankTransferPaymentMethodSpecificInput       `json:"bankTransferPaymentMethodSpecificInput,omitempty"`
	CardPaymentMethodSpecificInput            *CardPaymentMethodSpecificInput               `json:"cardPaymentMethodSpecificInput,omitempty"`
	CashPaymentMethodSpecificInput            *CashPaymentMethodSpecificInput               `json:"cashPaymentMethodSpecificInput,omitempty"`
	DirectDebitPaymentMethodSpecificInput     *NonSepaDirectDebitPaymentMethodSpecificInput `json:"directDebitPaymentMethodSpecificInput,omitempty"`
	EInvoicePaymentMethodSpecificInput        *EInvoicePaymentMethodSpecificInput           `json:"eInvoicePaymentMethodSpecificInput,omitempty"`
	EncryptedCustomerInput                    *string                                       `json:"encryptedCustomerInput,omitempty"`
	FraudFields                               *FraudFields                                  `json:"fraudFields,omitempty"`
	InvoicePaymentMethodSpecificInput         *InvoicePaymentMethodSpecificInput            `json:"invoicePaymentMethodSpecificInput,omitempty"`
	Merchant                                  *Merchant                                     `json:"merchant,omitempty"`
	MobilePaymentMethodSpecificInput          *MobilePaymentMethodSpecificInput             `json:"mobilePaymentMethodSpecificInput,omitempty"`
	Order                                     *Order                                        `json:"order,omitempty"`
	RedirectPaymentMethodSpecificInput        *RedirectPaymentMethodSpecificInput           `json:"redirectPaymentMethodSpecificInput,omitempty"`
	SepaDirectDebitPaymentMethodSpecificInput *SepaDirectDebitPaymentMethodSpecificInput    `json:"sepaDirectDebitPaymentMethodSpecificInput,omitempty"`
}

// NewCreatePaymentRequest constructs a new CreatePaymentRequest
func NewCreatePaymentRequest() *CreatePaymentRequest {
	return &CreatePaymentRequest{}
}
