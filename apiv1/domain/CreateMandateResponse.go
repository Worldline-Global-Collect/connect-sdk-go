// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CreateMandateResponse represents class CreateMandateResponse
type CreateMandateResponse struct {
	Mandate        *MandateResponse       `json:"mandate,omitempty"`
	MerchantAction *MandateMerchantAction `json:"merchantAction,omitempty"`
}

// NewCreateMandateResponse constructs a new CreateMandateResponse
func NewCreateMandateResponse() *CreateMandateResponse {
	return &CreateMandateResponse{}
}
