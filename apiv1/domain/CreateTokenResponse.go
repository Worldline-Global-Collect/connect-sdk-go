// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CreateTokenResponse represents class CreateTokenResponse
type CreateTokenResponse struct {
	IsNewToken        *bool   `json:"isNewToken,omitempty"`
	OriginalPaymentID *string `json:"originalPaymentId,omitempty"`
	Token             *string `json:"token,omitempty"`
}

// NewCreateTokenResponse constructs a new CreateTokenResponse instance
func NewCreateTokenResponse() *CreateTokenResponse {
	return &CreateTokenResponse{}
}
