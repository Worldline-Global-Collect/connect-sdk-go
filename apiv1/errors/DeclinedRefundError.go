// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package errors

import (
	goerr "errors"
	"fmt"
	"strconv"

	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

// DeclinedRefundError represents an error response from a refund call.
type DeclinedRefundError struct {
	errorMessage  string
	statusCode    int
	responseBody  string
	errorResponse *domain.RefundErrorResponse
}

// Message returns the error message
func (e DeclinedRefundError) Message() string {
	return e.errorMessage
}

// StatusCode returns the status code
func (e DeclinedRefundError) StatusCode() int {
	return e.statusCode
}

// ResponseBody returns the response body
func (e DeclinedRefundError) ResponseBody() string {
	return e.responseBody
}

// ErrorID implements the APIError interface
func (e DeclinedRefundError) ErrorID() string {
	if e.errorResponse.ErrorID == nil {
		return ""
	}
	return *e.errorResponse.ErrorID
}

// Errors implements the APIError interface
func (e DeclinedRefundError) Errors() []domain.APIError {
	// Return a clone instead of the original slice - immutability insurance
	if e.errorResponse.Errors == nil {
		return []domain.APIError{}
	}
	return append([]domain.APIError{}, *e.errorResponse.Errors...)
}

// RefundResult returns the result of creating a refund
func (e DeclinedRefundError) RefundResult() *domain.RefundResult {
	return e.errorResponse.RefundResult
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (e DeclinedRefundError) String() string {
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
func (e DeclinedRefundError) Error() string {
	return e.String()
}

// NewDeclinedRefundError creates a new DeclinedRefundError with the given statusCode, responseBody and errorResponse
func NewDeclinedRefundError(statusCode int, responseBody string, errorResponse *domain.RefundErrorResponse) (*DeclinedRefundError, error) {
	errorMessage := "the Worldline Global Collect platform returned a declined refund response"
	if errorResponse != nil && errorResponse.RefundResult != nil {
		refundResult := errorResponse.RefundResult
		if refundResult.ID == nil || refundResult.Status == nil {
			return nil, goerr.New("cannot get properties from refundResult")
		}
		errorMessage = fmt.Sprintf("declined refund '%s' with status '%s'", *refundResult.ID, *refundResult.Status)
	}
	return &DeclinedRefundError{errorMessage, statusCode, responseBody, errorResponse}, nil
}
