// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PayoutReferences represents class PayoutReferences
type PayoutReferences struct {
	InvoiceNumber     *string `json:"invoiceNumber,omitempty"`
	MerchantOrderID   *int64  `json:"merchantOrderId,omitempty"`
	MerchantReference *string `json:"merchantReference,omitempty"`
}

// NewPayoutReferences constructs a new PayoutReferences
func NewPayoutReferences() *PayoutReferences {
	return &PayoutReferences{}
}
