// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package errors

import (
	"net/http"

	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator"
)

// CreateAPIError is used internally in order to create an API error after an HTTP request is done
func CreateAPIError(statusCode int, responseBody string, errorObject interface{}, context *communicator.CallContext) (APIError, error) {
	var errorID string
	var errors []domain.APIError

	switch response := errorObject.(type) {
	case *domain.PaymentErrorResponse:
		{
			paymentResult := response.PaymentResult

			if paymentResult != nil {
				return NewDeclinedPaymentError(statusCode, responseBody, response)
			}

			if response.ErrorID != nil {
				errorID = *response.ErrorID
			}
			if response.Errors != nil {
				errors = *response.Errors
			}
			break
		}
	case *domain.PayoutErrorResponse:
		{
			payoutResult := response.PayoutResult

			if payoutResult != nil {
				return NewDeclinedPayoutError(statusCode, responseBody, response)
			}

			if response.ErrorID != nil {
				errorID = *response.ErrorID
			}
			if response.Errors != nil {
				errors = *response.Errors
			}

			break
		}
	case *domain.RefundErrorResponse:
		{
			refundResult := response.RefundResult

			if refundResult != nil {
				return NewDeclinedRefundError(statusCode, responseBody, response)
			}

			if response.ErrorID != nil {
				errorID = *response.ErrorID
			}
			if response.Errors != nil {
				errors = *response.Errors
			}

			break
		}
	case *domain.ErrorResponse:
		{
			if response.ErrorID != nil {
				errorID = *response.ErrorID
			}
			if response.Errors != nil {
				errors = *response.Errors
			}

			break
		}
	}

	switch statusCode {
	case http.StatusBadRequest:
		{
			return NewValidationError(statusCode, responseBody, errorID, errors)
		}
	case http.StatusForbidden:
		{
			return NewAuthorizationError(statusCode, responseBody, errorID, errors)
		}
	case http.StatusNotFound:
		{
			return NewReferenceError(statusCode, responseBody, errorID, errors)
		}
	case http.StatusConflict:
		{
			if isIdempotenceError(errors, context) {
				idempotenceKey := context.GetIdempotenceKey()
				idempotenceTimeStamp := context.GetIdempotenceRequestTimestamp()

				return NewIdempotenceError(idempotenceKey, idempotenceTimeStamp, statusCode, responseBody, errorID, errors)
			}

			return NewReferenceError(statusCode, responseBody, errorID, errors)
		}
	case http.StatusGone:
		{
			return NewReferenceError(statusCode, responseBody, errorID, errors)
		}
	case http.StatusInternalServerError:
		fallthrough
	case http.StatusBadGateway:
		fallthrough
	case http.StatusServiceUnavailable:
		fallthrough
	default:
		{
			return NewPlatformError(statusCode, responseBody, errorID, errors)
		}
	}
}

func isIdempotenceError(errors []domain.APIError, context *communicator.CallContext) (ok bool) {
	ok = context != nil && len(context.GetIdempotenceKey()) != 0 && len(errors) == 1 && errors[0].Code != nil && *errors[0].Code == "1409"

	return
}
