// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// FraudResultsRetailDecisions represents class FraudResultsRetailDecisions
type FraudResultsRetailDecisions struct {
	FraudCode   *string `json:"fraudCode,omitempty"`
	FraudNeural *string `json:"fraudNeural,omitempty"`
	FraudRCF    *string `json:"fraudRCF,omitempty"`
}

// NewFraudResultsRetailDecisions constructs a new FraudResultsRetailDecisions
func NewFraudResultsRetailDecisions() *FraudResultsRetailDecisions {
	return &FraudResultsRetailDecisions{}
}
