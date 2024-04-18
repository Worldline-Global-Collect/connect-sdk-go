// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CapturesResponse represents class CapturesResponse
type CapturesResponse struct {
	Captures *[]Capture `json:"captures,omitempty"`
}

// NewCapturesResponse constructs a new CapturesResponse instance
func NewCapturesResponse() *CapturesResponse {
	return &CapturesResponse{}
}
