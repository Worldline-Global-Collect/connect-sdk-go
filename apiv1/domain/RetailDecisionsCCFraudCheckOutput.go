// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RetailDecisionsCCFraudCheckOutput represents class RetailDecisionsCCFraudCheckOutput
type RetailDecisionsCCFraudCheckOutput struct {
	FraudCode   *string `json:"fraudCode,omitempty"`
	FraudNeural *string `json:"fraudNeural,omitempty"`
	FraudRCF    *string `json:"fraudRCF,omitempty"`
}

// NewRetailDecisionsCCFraudCheckOutput constructs a new RetailDecisionsCCFraudCheckOutput instance
func NewRetailDecisionsCCFraudCheckOutput() *RetailDecisionsCCFraudCheckOutput {
	return &RetailDecisionsCCFraudCheckOutput{}
}
