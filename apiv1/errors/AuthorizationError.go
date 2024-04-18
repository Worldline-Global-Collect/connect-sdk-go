// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package errors

import (
	"strconv"

	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

// AuthorizationError represents an error response from the Worldline Global Collect platform when authorization failed.
type AuthorizationError struct {
	errorMessage string
	statusCode   int
	responseBody string
	errorID      string
	errors       []domain.APIError
}

// Message returns the error message
func (e AuthorizationError) Message() string {
	return e.errorMessage
}

// StatusCode returns the status code
func (e AuthorizationError) StatusCode() int {
	return e.statusCode
}

// ResponseBody returns the response body
func (e AuthorizationError) ResponseBody() string {
	return e.responseBody
}

// ErrorID implements the APIError interface
func (e AuthorizationError) ErrorID() string {
	return e.errorID
}

// Errors implements the APIError interface
func (e AuthorizationError) Errors() []domain.APIError {
	// Return a clone instead of the original slice - immutability insurance
	return append([]domain.APIError{}, e.errors...)
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (e AuthorizationError) String() string {
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
func (e AuthorizationError) Error() string {
	return e.String()
}

// NewAuthorizationError creates a new AuthorizationError with the given statusCode, responseBody and response fields
func NewAuthorizationError(statusCode int, responseBody, errorID string, errors []domain.APIError) (*AuthorizationError, error) {
	return &AuthorizationError{"the Worldline Global Collect platform returned an authorization error response", statusCode, responseBody, errorID, errors}, nil
}

// NewAuthorizationErrorVerbose creates a new AuthorizationError with the given message, statusCode and response fields
func NewAuthorizationErrorVerbose(message string, statusCode int, responseBody, errorID string, errors []domain.APIError) (*AuthorizationError, error) {
	return &AuthorizationError{message, statusCode, responseBody, errorID, errors}, nil
}
