// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// Order represents class Order
type Order struct {
	AdditionalInput *AdditionalOrderInput `json:"additionalInput,omitempty"`
	AmountOfMoney   *AmountOfMoney        `json:"amountOfMoney,omitempty"`
	Customer        *Customer             `json:"customer,omitempty"`
	// Deprecated: Use shoppingCart.items instead
	Items           *[]LineItem           `json:"items,omitempty"`
	References      *OrderReferences      `json:"references,omitempty"`
	// Deprecated: Use Merchant.seller instead
	Seller          *Seller               `json:"seller,omitempty"`
	Shipping        *Shipping             `json:"shipping,omitempty"`
	ShoppingCart    *ShoppingCart         `json:"shoppingCart,omitempty"`
}

// NewOrder constructs a new Order instance
func NewOrder() *Order {
	return &Order{}
}
