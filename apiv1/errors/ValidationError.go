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
func (e ValidationError) Message() string {
	return e.errorMessage
}

// StatusCode returns the status code
func (e ValidationError) StatusCode() int {
	return e.statusCode
}

// ResponseBody returns the response body
func (e ValidationError) ResponseBody() string {
	return e.responseBody
}

// ErrorID implements the APIError interface
func (e ValidationError) ErrorID() string {
	return e.errorID
}

// Errors implements the APIError interface
func (e ValidationError) Errors() []domain.APIError {
	// Return a clone instead of the original slice - immutability insurance
	return append([]domain.APIError{}, e.errors...)
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (e ValidationError) String() string {
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
func (e ValidationError) Error() string {
	return e.String()
}

// NewValidationError creates a new ValidationError with the given statusCode, responseBody and response fields
func NewValidationError(statusCode int, responseBody, errorID string, errors []domain.APIError) (*ValidationError, error) {
	return &ValidationError{"the Worldline Global Collect platform returned an incorrect request error response", statusCode, responseBody, errorID, errors}, nil
}

// NewValidationErrorVerbose creates a new ValidationError with the given message, statusCode and response fields
func NewValidationErrorVerbose(message string, statusCode int, responseBody, errorID string, errors []domain.APIError) (*ValidationError, error) {
	return &ValidationError{message, statusCode, responseBody, errorID, errors}, nil
}
