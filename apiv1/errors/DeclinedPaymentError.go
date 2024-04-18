// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package errors

import (
	goerr "errors"
	"fmt"
	"strconv"

	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

// DeclinedPaymentError represents an error response from a create payment call.
type DeclinedPaymentError struct {
	errorMessage  string
	statusCode    int
	responseBody  string
	errorResponse *domain.PaymentErrorResponse
}

// Message returns the error message
func (e DeclinedPaymentError) Message() string {
	return e.errorMessage
}

// StatusCode returns the status code
func (e DeclinedPaymentError) StatusCode() int {
	return e.statusCode
}

// ResponseBody returns the response body
func (e DeclinedPaymentError) ResponseBody() string {
	return e.responseBody
}

// ErrorID implements the APIError interface
func (e DeclinedPaymentError) ErrorID() string {
	if e.errorResponse.ErrorID == nil {
		return ""
	}
	return *e.errorResponse.ErrorID
}

// Errors implements the APIError interface
func (e DeclinedPaymentError) Errors() []domain.APIError {
	// Return a clone instead of the original slice - immutability insurance
	if e.errorResponse.Errors == nil {
		return []domain.APIError{}
	}
	return append([]domain.APIError{}, *e.errorResponse.Errors...)
}

// PaymentResult returns the result of creating a payment
func (e DeclinedPaymentError) PaymentResult() *domain.CreatePaymentResult {
	return e.errorResponse.PaymentResult
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (e DeclinedPaymentError) String() string {
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
func (e DeclinedPaymentError) Error() string {
	return e.String()
}

// NewDeclinedPaymentError creates a new DeclinedPaymentError with the given statusCode, responseBody and errorResponse
func NewDeclinedPaymentError(statusCode int, responseBody string, errorResponse *domain.PaymentErrorResponse) (*DeclinedPaymentError, error) {
	errorMessage := "the Worldline Global Collect platform returned a declined payment response"
	if errorResponse != nil && errorResponse.PaymentResult != nil && errorResponse.PaymentResult.Payment != nil {
		payment := errorResponse.PaymentResult.Payment
		if payment.ID == nil || payment.Status == nil {
			return nil, goerr.New("cannot get properties from payment")
		}
		errorMessage = fmt.Sprintf("declined payment '%s' with status '%s'", *payment.ID, *payment.Status)
	}
	return &DeclinedPaymentError{errorMessage, statusCode, responseBody, errorResponse}, nil
}
