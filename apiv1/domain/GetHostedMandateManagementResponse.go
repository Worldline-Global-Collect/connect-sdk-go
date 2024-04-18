// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// GetHostedMandateManagementResponse represents class GetHostedMandateManagementResponse
type GetHostedMandateManagementResponse struct {
	Mandate *MandateResponse `json:"mandate,omitempty"`
	Status  *string          `json:"status,omitempty"`
}

// NewGetHostedMandateManagementResponse constructs a new GetHostedMandateManagementResponse instance
func NewGetHostedMandateManagementResponse() *GetHostedMandateManagementResponse {
	return &GetHostedMandateManagementResponse{}
}
