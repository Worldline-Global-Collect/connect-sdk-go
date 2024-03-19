// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CreateHostedMandateManagementRequest represents class CreateHostedMandateManagementRequest
type CreateHostedMandateManagementRequest struct {
	CreateMandateInfo                    *HostedMandateInfo                    `json:"createMandateInfo,omitempty"`
	HostedMandateManagementSpecificInput *HostedMandateManagementSpecificInput `json:"hostedMandateManagementSpecificInput,omitempty"`
}

// NewCreateHostedMandateManagementRequest constructs a new CreateHostedMandateManagementRequest
func NewCreateHostedMandateManagementRequest() *CreateHostedMandateManagementRequest {
	return &CreateHostedMandateManagementRequest{}
}
