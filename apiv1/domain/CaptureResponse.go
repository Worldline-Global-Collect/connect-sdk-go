// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CaptureResponse represents class CaptureResponse
type CaptureResponse struct {
	CaptureOutput *CaptureOutput       `json:"captureOutput,omitempty"`
	ID            *string              `json:"id,omitempty"`
	Status        *string              `json:"status,omitempty"`
	StatusOutput  *CaptureStatusOutput `json:"statusOutput,omitempty"`
}

// NewCaptureResponse constructs a new CaptureResponse
func NewCaptureResponse() *CaptureResponse {
	return &CaptureResponse{}
}
