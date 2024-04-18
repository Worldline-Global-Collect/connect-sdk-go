// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package errors

import "github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"

// APIError represents an error response from the Worldline Global Collect platform which contains an ID and a list of errors.
type APIError interface {
	// Error implements the error interface
	Error() string

	// Message gets the raw response body that was returned by the Worldline Global Collect platform.
	Message() string

	// StatusCode gets the HTTP status code that was returned by the Worldline Global Collect platform.
	StatusCode() int

	// ResponseBody gets the raw response body that was returned by the Worldline Global Collect platform.
	ResponseBody() string

	// ErrorID gets the errorId received from the Worldline Global Collect platform if available.
	ErrorID() string

	// Errors gets the errors received from the Worldline Global Collect platform if available. Never nil.
	Errors() []domain.APIError
}
