// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package sessions

import (
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator"
	commErrors "github.com/Worldline-Global-Collect/connect-sdk-go/communicator/errors"
)

// Client represents a sessions client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// Create represents the resource /{merchantId}/sessions - Create session
//
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/sessions/create.html
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
func (c *Client) Create(body domain.SessionRequest, context *communicator.CallContext) (domain.SessionResponse, error) {
	var resultObject domain.SessionResponse

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/sessions", nil)
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

// NewClient constructs a new Sessions client
//
// parent is the communicator.APIResource on top of which we want to build the new Sessions client
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
