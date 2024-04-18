// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package hostedmandatemanagements

import (
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator"
	commErrors "github.com/Worldline-Global-Collect/connect-sdk-go/communicator/errors"
)

// Client represents a hostedmandatemanagements client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// Create represents the resource /{merchantId}/hostedmandatemanagements - Create hosted mandate management
//
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/hostedmandatemanagements/create.html
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
func (c *Client) Create(body domain.CreateHostedMandateManagementRequest, context *communicator.CallContext) (domain.CreateHostedMandateManagementResponse, error) {
	var resultObject domain.CreateHostedMandateManagementResponse

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/hostedmandatemanagements", nil)
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

// Get represents the resource /{merchantId}/hostedmandatemanagements/{hostedMandateManagementId} - Get hosted mandate management status
//
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/hostedmandatemanagements/get.html
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
func (c *Client) Get(hostedMandateManagementID string, context *communicator.CallContext) (domain.GetHostedMandateManagementResponse, error) {
	var resultObject domain.GetHostedMandateManagementResponse

	pathContext := map[string]string{
		"hostedMandateManagementId": hostedMandateManagementID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/hostedmandatemanagements/{hostedMandateManagementId}", pathContext)
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

// NewClient constructs a new Hostedmandatemanagements client
//
// parent is the communicator.APIResource on top of which we want to build the new Hostedmandatemanagements client
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
