// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// GetHostedCheckoutResponse represents class GetHostedCheckoutResponse
type GetHostedCheckoutResponse struct {
	CreatedPaymentOutput *CreatedPaymentOutput `json:"createdPaymentOutput,omitempty"`
	Status               *string               `json:"status,omitempty"`
}

// NewGetHostedCheckoutResponse constructs a new GetHostedCheckoutResponse instance
func NewGetHostedCheckoutResponse() *GetHostedCheckoutResponse {
	return &GetHostedCheckoutResponse{}
}
