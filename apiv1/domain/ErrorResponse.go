// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ErrorResponse represents class ErrorResponse
type ErrorResponse struct {
	ErrorID *string     `json:"errorId,omitempty"`
	Errors  *[]APIError `json:"errors,omitempty"`
}

// NewErrorResponse constructs a new ErrorResponse instance
func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{}
}
