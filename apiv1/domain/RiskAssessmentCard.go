// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RiskAssessmentCard represents class RiskAssessmentCard
type RiskAssessmentCard struct {
	Card             *Card                   `json:"card,omitempty"`
	FraudFields      *FraudFields            `json:"fraudFields,omitempty"`
	Merchant         *MerchantRiskAssessment `json:"merchant,omitempty"`
	Order            *OrderRiskAssessment    `json:"order,omitempty"`
	PaymentProductID *int32                  `json:"paymentProductId,omitempty"`
}

// NewRiskAssessmentCard constructs a new RiskAssessmentCard
func NewRiskAssessmentCard() *RiskAssessmentCard {
	return &RiskAssessmentCard{}
}
