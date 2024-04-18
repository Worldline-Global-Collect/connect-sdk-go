// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// GetMandateResponse represents class GetMandateResponse
type GetMandateResponse struct {
	Mandate *MandateResponse `json:"mandate,omitempty"`
}

// NewGetMandateResponse constructs a new GetMandateResponse instance
func NewGetMandateResponse() *GetMandateResponse {
	return &GetMandateResponse{}
}
