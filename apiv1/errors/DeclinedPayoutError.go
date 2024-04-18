// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package errors

import (
	goerr "errors"
	"fmt"
	"strconv"

	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

// DeclinedPayoutError represents an error response from a payout call.
type DeclinedPayoutError struct {
	errorMessage  string
	statusCode    int
	responseBody  string
	errorResponse *domain.PayoutErrorResponse
}

// Message returns the error message
func (e DeclinedPayoutError) Message() string {
	return e.errorMessage
}

// StatusCode returns the status code
func (e DeclinedPayoutError) StatusCode() int {
	return e.statusCode
}

// ResponseBody returns the response body
func (e DeclinedPayoutError) ResponseBody() string {
	return e.responseBody
}

// ErrorID implements the APIError interface
func (e DeclinedPayoutError) ErrorID() string {
	if e.errorResponse.ErrorID == nil {
		return ""
	}
	return *e.errorResponse.ErrorID
}

// Errors implements the APIError interface
func (e DeclinedPayoutError) Errors() []domain.APIError {
	// Return a clone instead of the original slice - immutability insurance
	if e.errorResponse.Errors == nil {
		return []domain.APIError{}
	}
	return append([]domain.APIError{}, *e.errorResponse.Errors...)
}

// PayoutResult returns the result of creating a payout
func (e DeclinedPayoutError) PayoutResult() *domain.PayoutResult {
	return e.errorResponse.PayoutResult
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (e DeclinedPayoutError) String() string {
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
func (e DeclinedPayoutError) Error() string {
	return e.String()
}

// NewDeclinedPayoutError creates a new DeclinedPayoutError with the given statusCode, responseBody and errorResponse
func NewDeclinedPayoutError(statusCode int, responseBody string, errorResponse *domain.PayoutErrorResponse) (*DeclinedPayoutError, error) {
	errorMessage := "the Worldline Global Collect platform returned a declined payout response"
	if errorResponse != nil && errorResponse.PayoutResult != nil {
		payoutResult := errorResponse.PayoutResult
		if payoutResult.ID == nil || payoutResult.Status == nil {
			return nil, goerr.New("cannot get properties from payoutResult")
		}
		errorMessage = fmt.Sprintf("declined payout '%s' with status '%s'", *payoutResult.ID, *payoutResult.Status)
	}
	return &DeclinedPayoutError{errorMessage, statusCode, responseBody, errorResponse}, nil
}
