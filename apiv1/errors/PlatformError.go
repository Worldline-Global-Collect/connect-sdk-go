// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package errors

import (
	"strconv"

	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

// PlatformError represents an error response from the Worldline Global Collect platform when something went wrong at the Worldline Global Collect platform or further downstream.
type PlatformError struct {
	errorMessage string
	statusCode   int
	responseBody string
	errorID      string
	errors       []domain.APIError
}

// Message returns the error message
func (e PlatformError) Message() string {
	return e.errorMessage
}

// StatusCode returns the status code
func (e PlatformError) StatusCode() int {
	return e.statusCode
}

// ResponseBody returns the response body
func (e PlatformError) ResponseBody() string {
	return e.responseBody
}

// ErrorID implements the APIError interface
func (e PlatformError) ErrorID() string {
	return e.errorID
}

// Errors implements the APIError interface
func (e PlatformError) Errors() []domain.APIError {
	// Return a clone instead of the original slice - immutability insurance
	return append([]domain.APIError{}, e.errors...)
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (e PlatformError) String() string {
	list := e.errorMessage

	if e.statusCode > 0 {
		list = list + "; statusCode=" + strconv.Itoa(e.statusCode)
	}
	if len(e.responseBody) != 0 {
		list = list + "; responseBody='" + e.responseBody + "'"
	}

	return list
}

// Error implements the error interface
func (e PlatformError) Error() string {
	return e.String()
}

// NewPlatformError creates a new PlatformError with the given statusCode, responseBody and response fields
func NewPlatformError(statusCode int, responseBody, errorID string, errors []domain.APIError) (*PlatformError, error) {
	return &PlatformError{"the Worldline Global Collect platform returned an error response", statusCode, responseBody, errorID, errors}, nil
}

// NewPlatformErrorVerbose creates a new PlatformError with the given message, statusCode and response fields
func NewPlatformErrorVerbose(message string, statusCode int, responseBody, errorID string, errors []domain.APIError) (*PlatformError, error) {
	return &PlatformError{message, statusCode, responseBody, errorID, errors}, nil
}
