// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ApproveTokenRequest represents class ApproveTokenRequest
type ApproveTokenRequest struct {
	MandateSignatureDate  *string `json:"mandateSignatureDate,omitempty"`
	MandateSignaturePlace *string `json:"mandateSignaturePlace,omitempty"`
	MandateSigned         *bool   `json:"mandateSigned,omitempty"`
}

// NewApproveTokenRequest constructs a new ApproveTokenRequest instance
func NewApproveTokenRequest() *ApproveTokenRequest {
	return &ApproveTokenRequest{}
}
