// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package tokens

import (
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator"
	commErrors "github.com/Worldline-Global-Collect/connect-sdk-go/communicator/errors"
)

// Client represents a tokens client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// Create represents the resource /{merchantId}/tokens - Create token
//
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/tokens/create.html
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
func (c *Client) Create(body domain.CreateTokenRequest, context *communicator.CallContext) (domain.CreateTokenResponse, error) {
	var resultObject domain.CreateTokenResponse

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/tokens", nil)
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

// Get represents the resource /{merchantId}/tokens/{tokenId} - Get token
//
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/tokens/get.html
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
func (c *Client) Get(tokenID string, context *communicator.CallContext) (domain.TokenResponse, error) {
	var resultObject domain.TokenResponse

	pathContext := map[string]string{
		"tokenId": tokenID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/tokens/{tokenId}", pathContext)
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

// Update represents the resource /{merchantId}/tokens/{tokenId} - Update token
//
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/tokens/update.html
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
func (c *Client) Update(tokenID string, body domain.UpdateTokenRequest, context *communicator.CallContext) error {
	pathContext := map[string]string{
		"tokenId": tokenID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/tokens/{tokenId}", pathContext)
	if err != nil {
		return err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	var resultObject map[string]interface{}
	putErr := c.apiResource.Communicator().Put(uri, clientHeaders, nil, body, context, &resultObject)
	if putErr != nil {
		responseError, isResponseError := putErr.(*commErrors.ResponseError)
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

		return putErr
	}

	return nil
}

// Delete represents the resource /{merchantId}/tokens/{tokenId} - Delete token
//
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/tokens/delete.html
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
func (c *Client) Delete(tokenID string, query DeleteParams, context *communicator.CallContext) error {
	pathContext := map[string]string{
		"tokenId": tokenID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/tokens/{tokenId}", pathContext)
	if err != nil {
		return err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	var resultObject map[string]interface{}
	deleteErr := c.apiResource.Communicator().Delete(uri, clientHeaders, &query, context, &resultObject)
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

// Approvesepadirectdebit represents the resource /{merchantId}/tokens/{tokenId}/approvesepadirectdebit - Approve SEPA DD mandate
//
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/tokens/approvesepadirectdebit.html
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
func (c *Client) Approvesepadirectdebit(tokenID string, body domain.ApproveTokenRequest, context *communicator.CallContext) error {
	pathContext := map[string]string{
		"tokenId": tokenID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/tokens/{tokenId}/approvesepadirectdebit", pathContext)
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

// NewClient constructs a new Tokens client
//
// parent is the communicator.APIResource on top of which we want to build the new Tokens client
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
