// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CybersourceDecisionManager represents class CybersourceDecisionManager
type CybersourceDecisionManager struct {
	ClauseName    *string   `json:"clauseName,omitempty"`
	FraudScore    *int32    `json:"fraudScore,omitempty"`
	PolicyApplied *string   `json:"policyApplied,omitempty"`
	ReasonCodes   *[]string `json:"reasonCodes,omitempty"`
}

// NewCybersourceDecisionManager constructs a new CybersourceDecisionManager instance
func NewCybersourceDecisionManager() *CybersourceDecisionManager {
	return &CybersourceDecisionManager{}
}
