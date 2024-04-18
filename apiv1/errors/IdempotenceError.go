// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package errors

import (
	"strconv"

	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

// IdempotenceError represents an error response from the Worldline Global Collect platform when an idempotent request failed because the first request has not finished yet.
type IdempotenceError struct {
	idempotenceKey              string
	idempotenceRequestTimestamp *int64

	errorMessage string
	statusCode   int
	responseBody string
	errorID      string
	errors       []domain.APIError
}

// Message returns the error message
func (e IdempotenceError) Message() string {
	return e.errorMessage
}

// StatusCode returns the status code
func (e IdempotenceError) StatusCode() int {
	return e.statusCode
}

// ResponseBody returns the response body
func (e IdempotenceError) ResponseBody() string {
	return e.responseBody
}

// ErrorID implements the APIError interface
func (e IdempotenceError) ErrorID() string {
	return e.errorID
}

// Errors implements the APIError interface
func (e IdempotenceError) Errors() []domain.APIError {
	// Return a clone instead of the original slice - immutability insurance
	return append([]domain.APIError{}, e.errors...)
}

// IdempotenceKey returns the idempotence key used
func (e IdempotenceError) IdempotenceKey() string {
	return e.idempotenceKey
}

// IdempotenceRequestTimestamp returns the timestamp of the request
func (e IdempotenceError) IdempotenceRequestTimestamp() *int64 {
	return e.idempotenceRequestTimestamp
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (e IdempotenceError) String() string {
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
func (e IdempotenceError) Error() string {
	return e.String()
}

// NewIdempotenceError creates a new IdempotenceError with the given idempotenceKey, idempotenceRequestTimestamp, statusCode, responseBody and response fields
func NewIdempotenceError(idempotenceKey string, idempotenceRequestTimestamp *int64, statusCode int, responseBody, errorID string, errors []domain.APIError) (*IdempotenceError, error) {
	return &IdempotenceError{idempotenceKey, idempotenceRequestTimestamp, "the Worldline Global Collect platform returned a duplicate request error response", statusCode, responseBody, errorID, errors}, nil
}

// NewIdempotenceErrorVerbose creates a new IdempotenceError with the given idempotenceKey, idempotenceRequestTimestamp, message, statusCode, responseBody and response fields
func NewIdempotenceErrorVerbose(idempotenceKey string, idempotenceRequestTimestamp *int64, message string, statusCode int, responseBody, errorID string, errors []domain.APIError) (*IdempotenceError, error) {
	return &IdempotenceError{idempotenceKey, idempotenceRequestTimestamp, message, statusCode, responseBody, errorID, errors}, nil
}
