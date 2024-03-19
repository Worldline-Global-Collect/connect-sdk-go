// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package products

import (
	"strconv"

	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator"
	commErrors "github.com/Worldline-Global-Collect/connect-sdk-go/communicator/errors"
)

// Client represents a products client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// Find represents the resource /{merchantId}/products - Get payment products
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/products/find.html
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
func (c *Client) Find(query FindParams, context *communicator.CallContext) (domain.PaymentProducts, error) {
	var resultObject domain.PaymentProducts

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/products", nil)
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

// Get represents the resource /{merchantId}/products/{paymentProductId} - Get payment product
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/products/get.html
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
func (c *Client) Get(paymentProductID int32, query GetParams, context *communicator.CallContext) (domain.PaymentProductResponse, error) {
	var resultObject domain.PaymentProductResponse

	pathContext := map[string]string{
		"paymentProductId": strconv.FormatInt(int64(paymentProductID), 10),
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/products/{paymentProductId}", pathContext)
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

// Directory represents the resource /{merchantId}/products/{paymentProductId}/directory - Get payment product directory
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/products/directory.html
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
func (c *Client) Directory(paymentProductID int32, query DirectoryParams, context *communicator.CallContext) (domain.Directory, error) {
	var resultObject domain.Directory

	pathContext := map[string]string{
		"paymentProductId": strconv.FormatInt(int64(paymentProductID), 10),
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/products/{paymentProductId}/directory", pathContext)
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

// CustomerDetails represents the resource /{merchantId}/products/{paymentProductId}/customerDetails - Get customer details
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/products/customerDetails.html
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
func (c *Client) CustomerDetails(paymentProductID int32, body domain.GetCustomerDetailsRequest, context *communicator.CallContext) (domain.GetCustomerDetailsResponse, error) {
	var resultObject domain.GetCustomerDetailsResponse

	pathContext := map[string]string{
		"paymentProductId": strconv.FormatInt(int64(paymentProductID), 10),
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/products/{paymentProductId}/customerDetails", pathContext)
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

// DeviceFingerprint represents the resource /{merchantId}/products/{paymentProductId}/deviceFingerprint - Get device fingerprint
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/products/deviceFingerprint.html
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
func (c *Client) DeviceFingerprint(paymentProductID int32, body domain.DeviceFingerprintRequest, context *communicator.CallContext) (domain.DeviceFingerprintResponse, error) {
	var resultObject domain.DeviceFingerprintResponse

	pathContext := map[string]string{
		"paymentProductId": strconv.FormatInt(int64(paymentProductID), 10),
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/products/{paymentProductId}/deviceFingerprint", pathContext)
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

// Networks represents the resource /{merchantId}/products/{paymentProductId}/networks - Get payment product networks
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/products/networks.html
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
func (c *Client) Networks(paymentProductID int32, query NetworksParams, context *communicator.CallContext) (domain.PaymentProductNetworksResponse, error) {
	var resultObject domain.PaymentProductNetworksResponse

	pathContext := map[string]string{
		"paymentProductId": strconv.FormatInt(int64(paymentProductID), 10),
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/products/{paymentProductId}/networks", pathContext)
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

// Sessions represents the resource /{merchantId}/products/{paymentProductId}/sessions - Create session for payment product
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/products/sessions.html
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
func (c *Client) Sessions(paymentProductID int32, body domain.CreatePaymentProductSessionRequest, context *communicator.CallContext) (domain.CreatePaymentProductSessionResponse, error) {
	var resultObject domain.CreatePaymentProductSessionResponse

	pathContext := map[string]string{
		"paymentProductId": strconv.FormatInt(int64(paymentProductID), 10),
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/v1/{merchantId}/products/{paymentProductId}/sessions", pathContext)
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

// NewClient constructs a Products Client
//
// parent is the *communicator.APIResource on top of which we want to build the new Products Client
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
