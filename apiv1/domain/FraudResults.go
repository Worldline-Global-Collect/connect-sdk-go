// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// FraudResults represents class FraudResults
type FraudResults struct {
	FraudServiceResult       *string                `json:"fraudServiceResult,omitempty"`
	InAuth                   *InAuth                `json:"inAuth,omitempty"`
	MicrosoftFraudProtection *MicrosoftFraudResults `json:"microsoftFraudProtection,omitempty"`
}

// NewFraudResults constructs a new FraudResults
func NewFraudResults() *FraudResults {
	return &FraudResults{}
}
