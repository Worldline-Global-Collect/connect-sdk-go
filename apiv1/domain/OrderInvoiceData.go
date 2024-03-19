// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// OrderInvoiceData represents class OrderInvoiceData
type OrderInvoiceData struct {
	AdditionalData *string   `json:"additionalData,omitempty"`
	InvoiceDate    *string   `json:"invoiceDate,omitempty"`
	InvoiceNumber  *string   `json:"invoiceNumber,omitempty"`
	TextQualifiers *[]string `json:"textQualifiers,omitempty"`
}

// NewOrderInvoiceData constructs a new OrderInvoiceData
func NewOrderInvoiceData() *OrderInvoiceData {
	return &OrderInvoiceData{}
}
