// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package refunds

import (
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator"
	commErrors "github.com/Worldline-Global-Collect/connect-sdk-go/communicator/errors"
)

// Client represents a refunds client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// Find represents the resource /{merchantId}/refunds - Find refunds
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/refunds/find.html
//
// Can return any of the following errors:
// * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// * AuthorizationError if the request was not allowed (HTTP status code 403)
// * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
// or there was a conflict (HTTP status code 404, 409 or 410)
// * PlatformError if something went wrong at the Worldline Global Collect platform,
// the Worldline Global Collect platform was unable to process a message from a downstream partner/acquirer,
// or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// * APIError if the Worldline Global Collect platform returned any other error
func (c *Client) Find(query FindParams, context *communicator.CallContext) (domain.FindRefundsResponse, error) {
	var resultObject domain.FindRefundsResponse

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/refunds", nil)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	getErr := c.apiResource.Communicator().Get(uri, clientHeaders, &query, context, &resultObject)
	if getErr != nil {
		responseError, isResponseError := getErr.(*commErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			errorObject = &domain.ErrorResponse{}
			err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
			if err != nil {
				return resultObject, err
			}

			err, createErr := v1Errors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return resultObject, createErr
			}

			return resultObject, err
		}

		return resultObject, getErr
	}

	return resultObject, nil
}

// Get represents the resource /{merchantId}/refunds/{refundId} - Get refund
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/refunds/get.html
//
// Can return any of the following errors:
// * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// * AuthorizationError if the request was not allowed (HTTP status code 403)
// * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
// or there was a conflict (HTTP status code 404, 409 or 410)
// * PlatformError if something went wrong at the Worldline Global Collect platform,
// the Worldline Global Collect platform was unable to process a message from a downstream partner/acquirer,
// or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// * APIError if the Worldline Global Collect platform returned any other error
func (c *Client) Get(refundID string, context *communicator.CallContext) (domain.RefundResponse, error) {
	var resultObject domain.RefundResponse

	pathContext := map[string]string{
		"refundId": refundID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/refunds/{refundId}", pathContext)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	getErr := c.apiResource.Communicator().Get(uri, clientHeaders, nil, context, &resultObject)
	if getErr != nil {
		responseError, isResponseError := getErr.(*commErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			errorObject = &domain.ErrorResponse{}
			err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
			if err != nil {
				return resultObject, err
			}

			err, createErr := v1Errors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return resultObject, createErr
			}

			return resultObject, err
		}

		return resultObject, getErr
	}

	return resultObject, nil
}

// Approve represents the resource /{merchantId}/refunds/{refundId}/approve - Approve refund
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/refunds/approve.html
//
// Can return any of the following errors:
// * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// * AuthorizationError if the request was not allowed (HTTP status code 403)
// * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
// or there was a conflict (HTTP status code 404, 409 or 410)
// * PlatformError if something went wrong at the Worldline Global Collect platform,
// the Worldline Global Collect platform was unable to process a message from a downstream partner/acquirer,
// or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// * APIError if the Worldline Global Collect platform returned any other error
func (c *Client) Approve(refundID string, body domain.ApproveRefundRequest, context *communicator.CallContext) error {
	pathContext := map[string]string{
		"refundId": refundID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/refunds/{refundId}/approve", pathContext)
	if err != nil {
		return err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	var resultObject map[string]interface{}
	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, body, context, &resultObject)
	if postErr != nil {
		responseError, isResponseError := postErr.(*commErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			errorObject = &domain.ErrorResponse{}
			err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
			if err != nil {
				return err
			}

			err, createErr := v1Errors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return createErr
			}

			return err
		}

		return postErr
	}

	return nil
}

// Cancel represents the resource /{merchantId}/refunds/{refundId}/cancel - Cancel refund
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/refunds/cancel.html
//
// Can return any of the following errors:
// * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// * AuthorizationError if the request was not allowed (HTTP status code 403)
// * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
// or there was a conflict (HTTP status code 404, 409 or 410)
// * PlatformError if something went wrong at the Worldline Global Collect platform,
// the Worldline Global Collect platform was unable to process a message from a downstream partner/acquirer,
// or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// * APIError if the Worldline Global Collect platform returned any other error
func (c *Client) Cancel(refundID string, context *communicator.CallContext) error {
	pathContext := map[string]string{
		"refundId": refundID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/refunds/{refundId}/cancel", pathContext)
	if err != nil {
		return err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	var resultObject map[string]interface{}
	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, nil, context, &resultObject)
	if postErr != nil {
		responseError, isResponseError := postErr.(*commErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			errorObject = &domain.ErrorResponse{}
			err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
			if err != nil {
				return err
			}

			err, createErr := v1Errors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return createErr
			}

			return err
		}

		return postErr
	}

	return nil
}

// Cancelapproval represents the resource /{merchantId}/refunds/{refundId}/cancelapproval - Undo approve refund
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/refunds/cancelapproval.html
//
// Can return any of the following errors:
// * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// * AuthorizationError if the request was not allowed (HTTP status code 403)
// * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
// or there was a conflict (HTTP status code 404, 409 or 410)
// * PlatformError if something went wrong at the Worldline Global Collect platform,
// the Worldline Global Collect platform was unable to process a message from a downstream partner/acquirer,
// or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// * APIError if the Worldline Global Collect platform returned any other error
func (c *Client) Cancelapproval(refundID string, context *communicator.CallContext) error {
	pathContext := map[string]string{
		"refundId": refundID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/refunds/{refundId}/cancelapproval", pathContext)
	if err != nil {
		return err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	var resultObject map[string]interface{}
	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, nil, context, &resultObject)
	if postErr != nil {
		responseError, isResponseError := postErr.(*commErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			errorObject = &domain.ErrorResponse{}
			err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
			if err != nil {
				return err
			}

			err, createErr := v1Errors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return createErr
			}

			return err
		}

		return postErr
	}

	return nil
}

// NewClient constructs a Refunds Client
//
// parent is the *communicator.APIResource on top of which we want to build the new Refunds Client
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
