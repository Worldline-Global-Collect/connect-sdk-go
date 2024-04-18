// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// Capture represents class Capture
type Capture struct {
	CaptureOutput *CaptureOutput       `json:"captureOutput,omitempty"`
	ID            *string              `json:"id,omitempty"`
	Status        *string              `json:"status,omitempty"`
	StatusOutput  *CaptureStatusOutput `json:"statusOutput,omitempty"`
}

// NewCapture constructs a new Capture instance
func NewCapture() *Capture {
	return &Capture{}
}
