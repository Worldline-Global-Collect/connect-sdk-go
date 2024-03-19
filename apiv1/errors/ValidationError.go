// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package errors

import (
	"strconv"

	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

// ValidationError represents an error response from the Worldline Global Collect platform when validation of requests failed.
type ValidationError struct {
	errorMessage string
	statusCode   int
	responseBody string
	errorID      string
	errors       []domain.APIError
}

// Message returns the error message
func (ve ValidationError) Message() string {
	return ve.errorMessage
}

// StatusCode returns the status code
func (ve ValidationError) StatusCode() int {
	return ve.statusCode
}

// ResponseBody returns the response body
func (ve ValidationError) ResponseBody() string {
	return ve.responseBody
}

// ErrorID returns the error id
func (ve ValidationError) ErrorID() string {
	return ve.errorID
}

// Errors returns a slice of underlying errors
func (ve ValidationError) Errors() []domain.APIError {
	// Return a clone instead of the original slice - immutability insurance
	return append([]domain.APIError{}, ve.errors...)
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (ve ValidationError) String() string {
	list := ve.errorMessage

	if ve.statusCode > 0 {
		list = list + "; statusCode=" + strconv.Itoa(ve.statusCode)
	}
	if len(ve.responseBody) != 0 {
		list = list + "; responseBody='" + ve.responseBody + "'"
	}

	return list
}

// Error implements the error interface
func (ve ValidationError) Error() string {
	return ve.String()
}

// NewValidationError creates a ValidationError with the given statusCode, responseBody, errorID and errors
func NewValidationError(statusCode int, responseBody, errorID string, errors []domain.APIError) (*ValidationError, error) {
	return &ValidationError{"the Worldline Global Collect platform returned an incorrect request error response", statusCode, responseBody, errorID, errors}, nil
}

// NewValidationErrorVerbose creates a ValidationError with the given message, statusCode, responseBody, errorID and errors
func NewValidationErrorVerbose(message string, statusCode int, responseBody, errorID string, errors []domain.APIError) (*ValidationError, error) {
	return &ValidationError{message, statusCode, responseBody, errorID, errors}, nil
}
