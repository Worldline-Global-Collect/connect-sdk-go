// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package disputes

import (
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator"
	commErrors "github.com/Worldline-Global-Collect/connect-sdk-go/communicator/errors"
)

// Client represents a disputes client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// Get represents the resource /{merchantId}/disputes/{disputeId} - Get dispute
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/disputes/get.html
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
func (c *Client) Get(disputeID string, context *communicator.CallContext) (domain.DisputeResponse, error) {
	var resultObject domain.DisputeResponse

	pathContext := map[string]string{
		"disputeId": disputeID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/disputes/{disputeId}", pathContext)
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

// Submit represents the resource /{merchantId}/disputes/{disputeId}/submit - Submit dispute
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/disputes/submit.html
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
func (c *Client) Submit(disputeID string, context *communicator.CallContext) (domain.DisputeResponse, error) {
	var resultObject domain.DisputeResponse

	pathContext := map[string]string{
		"disputeId": disputeID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/disputes/{disputeId}/submit", pathContext)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, nil, context, &resultObject)
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

// Cancel represents the resource /{merchantId}/disputes/{disputeId}/cancel - Cancel dispute
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/disputes/cancel.html
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
func (c *Client) Cancel(disputeID string, context *communicator.CallContext) (domain.DisputeResponse, error) {
	var resultObject domain.DisputeResponse

	pathContext := map[string]string{
		"disputeId": disputeID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/disputes/{disputeId}/cancel", pathContext)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, nil, context, &resultObject)
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

// UploadFile represents the resource /{merchantId}/disputes/{disputeId} - Upload File
// Documentation can be found at https://apireference.connect.worldline-solutions.com/fileserviceapi/v1/en_US/go/disputes/uploadFile.html
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
func (c *Client) UploadFile(disputeID string, body UploadFileRequest, context *communicator.CallContext) (domain.UploadDisputeFileResponse, error) {
	var resultObject domain.UploadDisputeFileResponse

	pathContext := map[string]string{
		"disputeId": disputeID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/files/v1/{merchantId}/disputes/{disputeId}", pathContext)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, &body, context, &resultObject)
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

// NewClient constructs a Disputes Client
//
// parent is the *communicator.APIResource on top of which we want to build the new Disputes Client
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
