// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RiskAssessmentResponse represents class RiskAssessmentResponse
type RiskAssessmentResponse struct {
	Results *[]ResultDoRiskAssessment `json:"results,omitempty"`
}

// NewRiskAssessmentResponse constructs a new RiskAssessmentResponse instance
func NewRiskAssessmentResponse() *RiskAssessmentResponse {
	return &RiskAssessmentResponse{}
}
