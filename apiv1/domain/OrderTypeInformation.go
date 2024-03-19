// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// OrderTypeInformation represents class OrderTypeInformation
type OrderTypeInformation struct {
	FundingType     *string `json:"fundingType,omitempty"`
	PurchaseType    *string `json:"purchaseType,omitempty"`
	TransactionType *string `json:"transactionType,omitempty"`
	UsageType       *string `json:"usageType,omitempty"`
}

// NewOrderTypeInformation constructs a new OrderTypeInformation
func NewOrderTypeInformation() *OrderTypeInformation {
	return &OrderTypeInformation{}
}
