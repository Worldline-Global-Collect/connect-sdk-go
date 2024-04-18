// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package errors

import (
	"strconv"

	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

// ReferenceError represents an error response from the Worldline Global Collect platform when a non-existing or removed object is trying to be accessed.
type ReferenceError struct {
	errorMessage string
	statusCode   int
	responseBody string
	errorID      string
	errors       []domain.APIError
}

// Message returns the error message
func (e ReferenceError) Message() string {
	return e.errorMessage
}

// StatusCode returns the status code
func (e ReferenceError) StatusCode() int {
	return e.statusCode
}

// ResponseBody returns the response body
func (e ReferenceError) ResponseBody() string {
	return e.responseBody
}

// ErrorID implements the APIError interface
func (e ReferenceError) ErrorID() string {
	return e.errorID
}

// Errors implements the APIError interface
func (e ReferenceError) Errors() []domain.APIError {
	// Return a clone instead of the original slice - immutability insurance
	return append([]domain.APIError{}, e.errors...)
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (e ReferenceError) String() string {
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
func (e ReferenceError) Error() string {
	return e.String()
}

// NewReferenceError creates a new ReferenceError with the given statusCode, responseBody and response fields
func NewReferenceError(statusCode int, responseBody, errorID string, errors []domain.APIError) (*ReferenceError, error) {
	return &ReferenceError{"the Worldline Global Collect platform returned a reference error response", statusCode, responseBody, errorID, errors}, nil
}

// NewReferenceErrorVerbose creates a new ReferenceError with the given message, statusCode and response fields
func NewReferenceErrorVerbose(message string, statusCode int, responseBody, errorID string, errors []domain.APIError) (*ReferenceError, error) {
	return &ReferenceError{message, statusCode, responseBody, errorID, errors}, nil
}
