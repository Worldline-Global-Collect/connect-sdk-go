// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CaptureOutput represents class CaptureOutput
type CaptureOutput struct {
	AmountOfMoney                              *AmountOfMoney                                 `json:"amountOfMoney,omitempty"`
	AmountPaid                                 *int64                                         `json:"amountPaid,omitempty"`
	AmountReversed                             *int64                                         `json:"amountReversed,omitempty"`
	BankTransferPaymentMethodSpecificOutput    *BankTransferPaymentMethodSpecificOutput       `json:"bankTransferPaymentMethodSpecificOutput,omitempty"`
	CardPaymentMethodSpecificOutput            *CardPaymentMethodSpecificOutput               `json:"cardPaymentMethodSpecificOutput,omitempty"`
	CashPaymentMethodSpecificOutput            *CashPaymentMethodSpecificOutput               `json:"cashPaymentMethodSpecificOutput,omitempty"`
	DirectDebitPaymentMethodSpecificOutput     *NonSepaDirectDebitPaymentMethodSpecificOutput `json:"directDebitPaymentMethodSpecificOutput,omitempty"`
	EInvoicePaymentMethodSpecificOutput        *EInvoicePaymentMethodSpecificOutput           `json:"eInvoicePaymentMethodSpecificOutput,omitempty"`
	InvoicePaymentMethodSpecificOutput         *InvoicePaymentMethodSpecificOutput            `json:"invoicePaymentMethodSpecificOutput,omitempty"`
	MobilePaymentMethodSpecificOutput          *MobilePaymentMethodSpecificOutput             `json:"mobilePaymentMethodSpecificOutput,omitempty"`
	PaymentMethod                              *string                                        `json:"paymentMethod,omitempty"`
	RedirectPaymentMethodSpecificOutput        *RedirectPaymentMethodSpecificOutput           `json:"redirectPaymentMethodSpecificOutput,omitempty"`
	References                                 *PaymentReferences                             `json:"references,omitempty"`
	ReversalReason                             *string                                        `json:"reversalReason,omitempty"`
	SepaDirectDebitPaymentMethodSpecificOutput *SepaDirectDebitPaymentMethodSpecificOutput    `json:"sepaDirectDebitPaymentMethodSpecificOutput,omitempty"`
}

// NewCaptureOutput constructs a new CaptureOutput
func NewCaptureOutput() *CaptureOutput {
	return &CaptureOutput{}
}
