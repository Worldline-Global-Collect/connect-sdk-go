// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentOutput represents class PaymentOutput
type PaymentOutput struct {
	AmountCaptureRequested                     *int64                                         `json:"amountCaptureRequested,omitempty"`
	AmountCaptured                             *int64                                         `json:"amountCaptured,omitempty"`
	AmountOfMoney                              *AmountOfMoney                                 `json:"amountOfMoney,omitempty"`
	AmountPaid                                 *int64                                         `json:"amountPaid,omitempty"`
	AmountRefundRequested                      *int64                                         `json:"amountRefundRequested,omitempty"`
	AmountRefunded                             *int64                                         `json:"amountRefunded,omitempty"`
	AmountReversed                             *int64                                         `json:"amountReversed,omitempty"`
	BankTransferPaymentMethodSpecificOutput    *BankTransferPaymentMethodSpecificOutput       `json:"bankTransferPaymentMethodSpecificOutput,omitempty"`
	CardPaymentMethodSpecificOutput            *CardPaymentMethodSpecificOutput               `json:"cardPaymentMethodSpecificOutput,omitempty"`
	CashPaymentMethodSpecificOutput            *CashPaymentMethodSpecificOutput               `json:"cashPaymentMethodSpecificOutput,omitempty"`
	DirectDebitPaymentMethodSpecificOutput     *NonSepaDirectDebitPaymentMethodSpecificOutput `json:"directDebitPaymentMethodSpecificOutput,omitempty"`
	EInvoicePaymentMethodSpecificOutput        *EInvoicePaymentMethodSpecificOutput           `json:"eInvoicePaymentMethodSpecificOutput,omitempty"`
	InvoicePaymentMethodSpecificOutput         *InvoicePaymentMethodSpecificOutput            `json:"invoicePaymentMethodSpecificOutput,omitempty"`
	MobilePaymentMethodSpecificOutput          *MobilePaymentMethodSpecificOutput             `json:"mobilePaymentMethodSpecificOutput,omitempty"`
	Operations                                 *[]PaymentOperation                            `json:"operations,omitempty"`
	PaymentMethod                              *string                                        `json:"paymentMethod,omitempty"`
	RedirectPaymentMethodSpecificOutput        *RedirectPaymentMethodSpecificOutput           `json:"redirectPaymentMethodSpecificOutput,omitempty"`
	References                                 *PaymentReferences                             `json:"references,omitempty"`
	ReversalReason                             *string                                        `json:"reversalReason,omitempty"`
	SepaDirectDebitPaymentMethodSpecificOutput *SepaDirectDebitPaymentMethodSpecificOutput    `json:"sepaDirectDebitPaymentMethodSpecificOutput,omitempty"`
}

// NewPaymentOutput constructs a new PaymentOutput instance
func NewPaymentOutput() *PaymentOutput {
	return &PaymentOutput{}
}
