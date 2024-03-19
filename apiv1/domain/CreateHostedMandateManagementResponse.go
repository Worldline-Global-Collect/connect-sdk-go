// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CreateHostedMandateManagementResponse represents class CreateHostedMandateManagementResponse
type CreateHostedMandateManagementResponse struct {
	RETURNMAC                 *string `json:"RETURNMAC,omitempty"`
	HostedMandateManagementID *string `json:"hostedMandateManagementId,omitempty"`
	PartialRedirectURL        *string `json:"partialRedirectUrl,omitempty"`
}

// NewCreateHostedMandateManagementResponse constructs a new CreateHostedMandateManagementResponse
func NewCreateHostedMandateManagementResponse() *CreateHostedMandateManagementResponse {
	return &CreateHostedMandateManagementResponse{}
}
