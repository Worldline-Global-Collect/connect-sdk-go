// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// APIError represents class APIError
type APIError struct {
	Category       *string `json:"category,omitempty"`
	Code           *string `json:"code,omitempty"`
	HTTPStatusCode *int32  `json:"httpStatusCode,omitempty"`
	ID             *string `json:"id,omitempty"`
	Message        *string `json:"message,omitempty"`
	PropertyName   *string `json:"propertyName,omitempty"`
	RequestID      *string `json:"requestId,omitempty"`
}

// NewAPIError constructs a new APIError instance
func NewAPIError() *APIError {
	return &APIError{}
}
