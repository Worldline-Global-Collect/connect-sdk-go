// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package files

import (
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator"
	commErrors "github.com/Worldline-Global-Collect/connect-sdk-go/communicator/errors"
)

// Client represents a files client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// GetFile represents the resource /{merchantId}/files/{fileId} - Retrieve File
//
// Documentation can be found at https://apireference.connect.worldline-solutions.com/fileserviceapi/v1/en_US/go/files/getFile.html
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
func (c *Client) GetFile(fileID string, context *communicator.CallContext, bodyHandler communicator.BodyHandler) error {
	pathContext := map[string]string{
		"fileId": fileID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/files/v1/{merchantId}/files/{fileId}", pathContext)
	if err != nil {
		return err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	getErr := c.apiResource.Communicator().GetWithHandler(uri, clientHeaders, nil, context, bodyHandler)
	if getErr != nil {
		responseError, isResponseError := getErr.(*commErrors.ResponseError)
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

		return getErr
	}

	return nil
}

// NewClient constructs a new Files client
//
// parent is the communicator.APIResource on top of which we want to build the new Files client
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
