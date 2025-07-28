// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CaptureStatusOutput represents class CaptureStatusOutput
type CaptureStatusOutput struct {
	IsRefundable             *bool           `json:"isRefundable,omitempty"`
	IsRetriable              *bool           `json:"isRetriable,omitempty"`
	ProviderRawOutput        *[]KeyValuePair `json:"providerRawOutput,omitempty"`
	StatusCode               *int32          `json:"statusCode,omitempty"`
	StatusCodeChangeDateTime *string         `json:"statusCodeChangeDateTime,omitempty"`
}

// NewCaptureStatusOutput constructs a new CaptureStatusOutput instance
func NewCaptureStatusOutput() *CaptureStatusOutput {
	return &CaptureStatusOutput{}
}
