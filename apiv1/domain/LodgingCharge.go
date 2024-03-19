// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// LodgingCharge represents class LodgingCharge
type LodgingCharge struct {
	ChargeAmount             *int64  `json:"chargeAmount,omitempty"`
	ChargeAmountCurrencyCode *string `json:"chargeAmountCurrencyCode,omitempty"`
	ChargeType               *string `json:"chargeType,omitempty"`
}

// NewLodgingCharge constructs a new LodgingCharge
func NewLodgingCharge() *LodgingCharge {
	return &LodgingCharge{}
}
