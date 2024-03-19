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
func (gce PlatformError) Message() string {
	return gce.errorMessage
}

// StatusCode returns the status code
func (gce PlatformError) StatusCode() int {
	return gce.statusCode
}

// ResponseBody returns the response body
func (gce PlatformError) ResponseBody() string {
	return gce.responseBody
}

// ErrorID returns the error id
func (gce PlatformError) ErrorID() string {
	return gce.errorID
}

// Errors returns a slice of underlying errors
func (gce PlatformError) Errors() []domain.APIError {
	// Return a clone instead of the original slice - immutability insurance
	return append([]domain.APIError{}, gce.errors...)
}

// String implements the Stringer ineterface
// Format: 'errorMessage; statusCode=; responseBody='
func (gce PlatformError) String() string {
	list := gce.errorMessage

	if gce.statusCode > 0 {
		list = list + "; statusCode=" + strconv.Itoa(gce.statusCode)
	}
	if len(gce.responseBody) != 0 {
		list = list + "; responseBody='" + gce.responseBody + "'"
	}

	return list
}

// Error implements the error interface
func (gce PlatformError) Error() string {
	return gce.String()
}

// NewPlatformError creates a PlatformError with the given statusCode, responseBody, errorID and errors
func NewPlatformError(statusCode int, responseBody, errorID string, errors []domain.APIError) (*PlatformError, error) {
	return &PlatformError{"the Worldline Global Collect platform returned an error response", statusCode, responseBody, errorID, errors}, nil
}

// NewPlatformErrorVerbose creates a PlatformError with the given message, statusCode, responseBody, errorID and errors
func NewPlatformErrorVerbose(message string, statusCode int, responseBody, errorID string, errors []domain.APIError) (*PlatformError, error) {
	return &PlatformError{message, statusCode, responseBody, errorID, errors}, nil
}
