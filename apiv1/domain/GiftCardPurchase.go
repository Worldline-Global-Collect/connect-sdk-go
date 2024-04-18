// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// GiftCardPurchase represents class GiftCardPurchase
type GiftCardPurchase struct {
	AmountOfMoney     *AmountOfMoney `json:"amountOfMoney,omitempty"`
	NumberOfGiftCards *int32         `json:"numberOfGiftCards,omitempty"`
}

// NewGiftCardPurchase constructs a new GiftCardPurchase instance
func NewGiftCardPurchase() *GiftCardPurchase {
	return &GiftCardPurchase{}
}
