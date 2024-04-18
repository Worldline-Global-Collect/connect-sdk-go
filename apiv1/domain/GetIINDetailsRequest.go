// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// GetIINDetailsRequest represents class GetIINDetailsRequest
type GetIINDetailsRequest struct {
	Bin            *string         `json:"bin,omitempty"`
	PaymentContext *PaymentContext `json:"paymentContext,omitempty"`
}

// NewGetIINDetailsRequest constructs a new GetIINDetailsRequest instance
func NewGetIINDetailsRequest() *GetIINDetailsRequest {
	return &GetIINDetailsRequest{}
}
