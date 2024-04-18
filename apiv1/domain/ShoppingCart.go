// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ShoppingCart represents class ShoppingCart
type ShoppingCart struct {
	AmountBreakdown              *[]AmountBreakdown `json:"amountBreakdown,omitempty"`
	GiftCardPurchase             *GiftCardPurchase  `json:"giftCardPurchase,omitempty"`
	IsPreOrder                   *bool              `json:"isPreOrder,omitempty"`
	Items                        *[]LineItem        `json:"items,omitempty"`
	PreOrderItemAvailabilityDate *string            `json:"preOrderItemAvailabilityDate,omitempty"`
	ReOrderIndicator             *bool              `json:"reOrderIndicator,omitempty"`
}

// NewShoppingCart constructs a new ShoppingCart instance
func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}
