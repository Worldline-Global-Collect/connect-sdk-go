// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// LineItemInvoiceData represents class LineItemInvoiceData
type LineItemInvoiceData struct {
	Description        *string `json:"description,omitempty"`
	MerchantLinenumber *string `json:"merchantLinenumber,omitempty"`
	MerchantPagenumber *string `json:"merchantPagenumber,omitempty"`
	NrOfItems          *string `json:"nrOfItems,omitempty"`
	PricePerItem       *int64  `json:"pricePerItem,omitempty"`
}

// NewLineItemInvoiceData constructs a new LineItemInvoiceData instance
func NewLineItemInvoiceData() *LineItemInvoiceData {
	return &LineItemInvoiceData{}
}
