// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// WebhooksEvent represents an event that is sent by webhooks
type WebhooksEvent struct {
	APIVersion *string          `json:"apiVersion,omitempty"`
	Created    *string          `json:"created,omitempty"`
	Dispute    *DisputeResponse `json:"dispute,omitempty"`
	ID         *string          `json:"id,omitempty"`
	MerchantID *string          `json:"merchantId,omitempty"`
	Payment    *PaymentResponse `json:"payment,omitempty"`
	Payout     *PayoutResponse  `json:"payout,omitempty"`
	Refund     *RefundResponse  `json:"refund,omitempty"`
	Token      *TokenResponse   `json:"token,omitempty"`
	Type       *string          `json:"type,omitempty"`
}

// NewWebhooksEvent constructs a new WebhooksEvent instance
func NewWebhooksEvent() *WebhooksEvent {
	return &WebhooksEvent{}
}
