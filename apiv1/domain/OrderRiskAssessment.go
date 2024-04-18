// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// OrderRiskAssessment represents class OrderRiskAssessment
type OrderRiskAssessment struct {
	AdditionalInput *AdditionalOrderInputAirlineData `json:"additionalInput,omitempty"`
	AmountOfMoney   *AmountOfMoney                   `json:"amountOfMoney,omitempty"`
	Customer        *CustomerRiskAssessment          `json:"customer,omitempty"`
	Shipping        *ShippingRiskAssessment          `json:"shipping,omitempty"`
}

// NewOrderRiskAssessment constructs a new OrderRiskAssessment instance
func NewOrderRiskAssessment() *OrderRiskAssessment {
	return &OrderRiskAssessment{}
}
