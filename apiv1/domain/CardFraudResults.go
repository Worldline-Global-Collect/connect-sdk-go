// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CardFraudResults represents class CardFraudResults
type CardFraudResults struct {
	AvsResult                  *string                      `json:"avsResult,omitempty"`
	CvvResult                  *string                      `json:"cvvResult,omitempty"`
	CybersourceDecisionManager *CybersourceDecisionManager  `json:"cybersourceDecisionManager,omitempty"`
	FraudServiceResult         *string                      `json:"fraudServiceResult,omitempty"`
	Fraugster                  *FraugsterResults            `json:"fraugster,omitempty"`
	InAuth                     *InAuth                      `json:"inAuth,omitempty"`
	MicrosoftFraudProtection   *MicrosoftFraudResults       `json:"microsoftFraudProtection,omitempty"`
	RetailDecisions            *FraudResultsRetailDecisions `json:"retailDecisions,omitempty"`
}

// NewCardFraudResults constructs a new CardFraudResults instance
func NewCardFraudResults() *CardFraudResults {
	return &CardFraudResults{}
}
