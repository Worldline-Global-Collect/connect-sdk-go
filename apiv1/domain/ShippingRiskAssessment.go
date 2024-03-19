// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ShippingRiskAssessment represents class ShippingRiskAssessment
type ShippingRiskAssessment struct {
	Address        *AddressPersonal `json:"address,omitempty"`
	Comments       *string          `json:"comments,omitempty"`
	TrackingNumber *string          `json:"trackingNumber,omitempty"`
}

// NewShippingRiskAssessment constructs a new ShippingRiskAssessment
func NewShippingRiskAssessment() *ShippingRiskAssessment {
	return &ShippingRiskAssessment{}
}
