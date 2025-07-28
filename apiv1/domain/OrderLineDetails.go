// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// OrderLineDetails represents class OrderLineDetails
type OrderLineDetails struct {
	DiscountAmount          *int64  `json:"discountAmount,omitempty"`
	GoogleProductCategoryID *int64  `json:"googleProductCategoryId,omitempty"`
	LineAmountTotal         *int64  `json:"lineAmountTotal,omitempty"`
	NaicsCommodityCode      *string `json:"naicsCommodityCode,omitempty"`
	ProductCategory         *string `json:"productCategory,omitempty"`
	ProductCode             *string `json:"productCode,omitempty"`
	ProductImageURL         *string `json:"productImageUrl,omitempty"`
	ProductName             *string `json:"productName,omitempty"`
	ProductPrice            *int64  `json:"productPrice,omitempty"`
	ProductSku              *string `json:"productSku,omitempty"`
	ProductType             *string `json:"productType,omitempty"`
	ProductURL              *string `json:"productUrl,omitempty"`
	Quantity                *int64  `json:"quantity,omitempty"`
	TaxAmount               *int64  `json:"taxAmount,omitempty"`
	Unit                    *string `json:"unit,omitempty"`
}

// NewOrderLineDetails constructs a new OrderLineDetails instance
func NewOrderLineDetails() *OrderLineDetails {
	return &OrderLineDetails{}
}
