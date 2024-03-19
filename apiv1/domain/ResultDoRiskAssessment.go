// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ResultDoRiskAssessment represents class ResultDoRiskAssessment
type ResultDoRiskAssessment struct {
	Category                          *string                            `json:"category,omitempty"`
	Result                            *string                            `json:"result,omitempty"`
	RetaildecisionsCCFraudCheckOutput *RetailDecisionsCCFraudCheckOutput `json:"retaildecisionsCCFraudCheckOutput,omitempty"`
	ValidationBankAccountOutput       *ValidationBankAccountOutput       `json:"validationBankAccountOutput,omitempty"`
}

// NewResultDoRiskAssessment constructs a new ResultDoRiskAssessment
func NewResultDoRiskAssessment() *ResultDoRiskAssessment {
	return &ResultDoRiskAssessment{}
}
