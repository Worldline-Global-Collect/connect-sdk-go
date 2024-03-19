// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentStatusOutput represents class PaymentStatusOutput
type PaymentStatusOutput struct {
	Errors                   *[]APIError     `json:"errors,omitempty"`
	IsAuthorized             *bool           `json:"isAuthorized,omitempty"`
	IsCancellable            *bool           `json:"isCancellable,omitempty"`
	IsRefundable             *bool           `json:"isRefundable,omitempty"`
	IsRetriable              *bool           `json:"isRetriable,omitempty"`
	ProviderRawOutput        *[]KeyValuePair `json:"providerRawOutput,omitempty"`
	StatusCategory           *string         `json:"statusCategory,omitempty"`
	StatusCode               *int32          `json:"statusCode,omitempty"`
	StatusCodeChangeDateTime *string         `json:"statusCodeChangeDateTime,omitempty"`
	ThreeDSecureStatus       *string         `json:"threeDSecureStatus,omitempty"`
}

// NewPaymentStatusOutput constructs a new PaymentStatusOutput
func NewPaymentStatusOutput() *PaymentStatusOutput {
	return &PaymentStatusOutput{}
}
