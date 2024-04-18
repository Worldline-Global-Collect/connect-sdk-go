// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// OrderStatusOutput represents class OrderStatusOutput
type OrderStatusOutput struct {
	Errors                   *[]APIError     `json:"errors,omitempty"`
	IsCancellable            *bool           `json:"isCancellable,omitempty"`
	IsRetriable              *bool           `json:"isRetriable,omitempty"`
	ProviderRawOutput        *[]KeyValuePair `json:"providerRawOutput,omitempty"`
	StatusCategory           *string         `json:"statusCategory,omitempty"`
	StatusCode               *int32          `json:"statusCode,omitempty"`
	StatusCodeChangeDateTime *string         `json:"statusCodeChangeDateTime,omitempty"`
}

// NewOrderStatusOutput constructs a new OrderStatusOutput instance
func NewOrderStatusOutput() *OrderStatusOutput {
	return &OrderStatusOutput{}
}
