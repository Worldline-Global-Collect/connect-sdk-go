// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package hostedcheckouts

import (
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator"
	commErrors "github.com/Worldline-Global-Collect/connect-sdk-go/communicator/errors"
)

// Client represents a hostedcheckouts client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// Create represents the resource /{merchantId}/hostedcheckouts - Create hosted checkout
//
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/hostedcheckouts/create.html
//
// Can return any of the following errors:
//   * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
//   * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
//   * AuthorizationError if the request was not allowed (HTTP status code 403)
//   * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//     or there was a conflict (HTTP status code 404, 409 or 410)
//   * PlatformError if something went wrong at the Worldline Global Collect platform,
//     the Worldline Global Collect platform was unable to process a message from a downstream partner/acquirer,
//     or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
//   * APIError if the Worldline Global Collect platform returned any other error
func (c *Client) Create(body domain.CreateHostedCheckoutRequest, context *communicator.CallContext) (domain.CreateHostedCheckoutResponse, error) {
	var resultObject domain.CreateHostedCheckoutResponse

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/hostedcheckouts", nil)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, body, context, &resultObject)
	if postErr != nil {
		responseError, isResponseError := postErr.(*commErrors.ResponseError)
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

		return resultObject, postErr
	}

	return resultObject, nil
}

// Get represents the resource /{merchantId}/hostedcheckouts/{hostedCheckoutId} - Get hosted checkout status
//
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/hostedcheckouts/get.html
//
// Can return any of the following errors:
//   * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
//   * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
//   * AuthorizationError if the request was not allowed (HTTP status code 403)
//   * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//     or there was a conflict (HTTP status code 404, 409 or 410)
//   * PlatformError if something went wrong at the Worldline Global Collect platform,
//     the Worldline Global Collect platform was unable to process a message from a downstream partner/acquirer,
//     or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
//   * APIError if the Worldline Global Collect platform returned any other error
func (c *Client) Get(hostedCheckoutID string, context *communicator.CallContext) (domain.GetHostedCheckoutResponse, error) {
	var resultObject domain.GetHostedCheckoutResponse

	pathContext := map[string]string{
		"hostedCheckoutId": hostedCheckoutID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/hostedcheckouts/{hostedCheckoutId}", pathContext)
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

// Delete represents the resource /{merchantId}/hostedcheckouts/{hostedCheckoutId} - Delete hosted checkout
//
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/hostedcheckouts/delete.html
//
// Can return any of the following errors:
//   * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
//   * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
//   * AuthorizationError if the request was not allowed (HTTP status code 403)
//   * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//     or there was a conflict (HTTP status code 404, 409 or 410)
//   * PlatformError if something went wrong at the Worldline Global Collect platform,
//     the Worldline Global Collect platform was unable to process a message from a downstream partner/acquirer,
//     or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
//   * APIError if the Worldline Global Collect platform returned any other error
func (c *Client) Delete(hostedCheckoutID string, context *communicator.CallContext) error {
	pathContext := map[string]string{
		"hostedCheckoutId": hostedCheckoutID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/hostedcheckouts/{hostedCheckoutId}", pathContext)
	if err != nil {
		return err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	var resultObject map[string]interface{}
	deleteErr := c.apiResource.Communicator().Delete(uri, clientHeaders, nil, context, &resultObject)
	if deleteErr != nil {
		responseError, isResponseError := deleteErr.(*commErrors.ResponseError)
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

		return deleteErr
	}

	return nil
}

// NewClient constructs a new Hostedcheckouts client
//
// parent is the communicator.APIResource on top of which we want to build the new Hostedcheckouts client
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
