package connectsdk

import (
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors"
	"github.com/Worldline-Global-Collect/connect-sdk-go/authentication/v1hmac"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
	commErrors "github.com/Worldline-Global-Collect/connect-sdk-go/communicator/errors"
	"github.com/Worldline-Global-Collect/connect-sdk-go/logging"
	"github.com/Worldline-Global-Collect/connect-sdk-go/logging/obfuscation"
)

func checkSuccess(connection *TestConnection, response domain.CreatePaymentResponse, err error) string {
	if err != nil {
		return "Error during request: " + err.Error()
	}
	if payment := response.Payment; payment == nil || payment.ID == nil || *payment.ID != "000002000020142549460000100001" {
		return "The ID of the payment is not equal to 000002000020142549460000100001"
	}
	if payment := response.Payment; payment == nil || payment.Status == nil || *payment.Status != "PENDING_APPROVAL" {
		return "The status of the payment is unequal to PENDING_APPROVAL"
	}
	return ""
}

func checkRejected(connection *TestConnection, response domain.CreatePaymentResponse, err error) string {
	if err == nil {
		return "Error expected"
	}
	if dpe, ok := err.(*v1Errors.DeclinedPaymentError); ok {
		res := dpe.PaymentResult()
		if res == nil {
			return "PaymentResult is nil"
		}
		if res.Payment.ID == nil {
			return "PaymentResult.ID is nil"
		}
		if *res.Payment.ID != "000002000020142544360000100001" {
			return "PaymentResult.ID is not 000002000020142544360000100001"
		}
		if res.Payment.Status == nil {
			return "PaymentResult.Status is nil"
		}
		if *res.Payment.Status != "REJECTED" {
			return "PaymentResult.Status is not equal to " + "REJECTED"
		}
	} else {
		return "Expected DeclinedPaymentError, but got different error"
	}
	errorString := err.Error()
	if !strings.Contains(errorString, "payment '000002000020142544360000100001'") {
		return "Error message does not contain" + "payment '000002000020142544360000100001'"
	}
	if !strings.Contains(errorString, "status 'REJECTED'") {
		return "Error message: " + errorString + " does not contain " + "status 'REJECTED'"
	}
	if !strings.Contains(errorString, rejectedJSON) {
		return "Error message: " + errorString + " does not contain response body"
	}
	return ""
}

func checkInvalidRequest(connection *TestConnection, response domain.CreatePaymentResponse, err error) string {
	if err == nil {
		return "Error expected"
	}
	if _, ok := err.(*v1Errors.ValidationError); !ok {
		return "Expected ValidationError, but got different error"
	}
	errorString := err.Error()
	if !strings.Contains(errorString, invalidRequestJSON) {
		return "Error message: " + errorString + " does not contain response body"
	}
	return ""
}

func checkInvalidAuthorization(connection *TestConnection, response domain.CreatePaymentResponse, err error) string {
	if err == nil {
		return "Error expected"
	}
	if _, ok := err.(v1Errors.APIError); !ok {
		return "Expected APIError, but got different error"
	}
	errorString := err.Error()
	if !strings.Contains(errorString, invalidAuthorizationJSON) {
		return "Error message: " + errorString + " does not contain response body"
	}
	return ""
}

func checkReferenceError(connection *TestConnection, response domain.CreatePaymentResponse, err error) string {
	if err == nil {
		return "Error expected"
	}
	if _, ok := err.(*v1Errors.ReferenceError); !ok {
		return "Expected ReferenceError, but got different error"
	}
	errorString := err.Error()
	if !strings.Contains(errorString, duplicateRequestJSON) {
		return "Error message: " + errorString + " does not contain response body"
	}
	return ""
}

func checkIdempotenceError(connection *TestConnection, response domain.CreatePaymentResponse, err error) string {
	if err == nil {
		return "Error expected"
	}
	if ie, ok := err.(*v1Errors.IdempotenceError); ok {
		if ie.IdempotenceKey() != connection.idempotenceKey {
			return "Returned wrong idempotenceKey"
		}
	} else {
		return "Expected IdempotenceError, but got different error"
	}
	errorString := err.Error()
	if !strings.Contains(errorString, duplicateRequestJSON) {
		return "Error message: " + errorString + " does not contain response body"
	}
	return ""
}

func checkNotFound(connection *TestConnection, response domain.CreatePaymentResponse, err error) string {
	if err == nil {
		return "Error expected"
	}
	if nfe, ok := err.(*commErrors.NotFoundError); ok {
		if nfe.InternalError() == nil {
			return "Returned NotFoundError without internal error"
		}
		if _, ok := nfe.InternalError().(*commErrors.ResponseError); !ok {
			return "NotFoundError has a different internal error than ResponseError"
		}
		if !strings.Contains(nfe.InternalError().Error(), notFoundHTML) {
			return "Error message: " + nfe.InternalError().Error() + " does not contain response body"
		}
	} else {
		return "Expected NotFoundError, but got different error"
	}

	return ""
}

func checkMethodNotAllowed(connection *TestConnection, response domain.CreatePaymentResponse, err error) string {
	if err == nil {
		return "Error expected"
	}
	if nfe, ok := err.(*commErrors.CommunicationError); ok {
		if nfe.InternalError() == nil {
			return "Returned CommunicationError without internal error"
		}
		if _, ok := nfe.InternalError().(*commErrors.ResponseError); !ok {
			return "CommunicationError has a different internal error than ResponseError"
		}
		if !strings.Contains(nfe.InternalError().Error(), methodNotAllowedHTML) {
			return "Error message: " + nfe.InternalError().Error() + " does not contain response body"
		}
	} else {
		return "Expected CommunicationError, but got different error"
	}

	return ""
}

var table = []TestConnection{
	// Tests that a non-failure response will not throw an exception.
	{pendingApprovalJSON, 201, nil, "", checkSuccess},

	//Tests that a failure response with a payment result will return a DeclinedPaymentError.
	{rejectedJSON, 400, nil, "", checkRejected},

	// Tests that a 400 failure response without a payment result will return a ValidateError.
	{invalidRequestJSON, 400, nil, "", checkInvalidRequest},

	// Tests that a 401 failure response without a payment result will return a APIError.
	{invalidAuthorizationJSON, 401, nil, "", checkInvalidAuthorization},

	// Tests that a 409 failure response with a duplicate request code but without an idempotence key will return a ReferenceError
	{duplicateRequestJSON, 409, nil, "", checkReferenceError},

	// Tests that a 409 failure response with a duplicate request code and an idempotence key will return a IdempotenceError.
	{duplicateRequestJSON, 409, nil, "key", checkIdempotenceError},

	// Tests that a 404 response with a non-JSON response will throw a NotFoundException.
	{notFoundHTML, 404, []communication.Header{newHeader("content-type", "text/html")}, "", checkNotFound},

	// Tests that a 405 response with a non-JSON response will throw a CommunicationException.
	{methodNotAllowedHTML, 405, []communication.Header{newHeader("content-type", "text/html")}, "", checkMethodNotAllowed},
}

func newHeader(name, value string) communication.Header {
	h, _ := communication.NewHeader(name, value)
	return *h
}

type TestConnection struct {
	body           string
	statusCode     int
	headers        []communication.Header
	idempotenceKey string
	checkResult    CheckResult
}

const pendingApprovalJSON string = `{
    "creationOutput": {
        "additionalReference": "00000200002014254946",
        "externalReference": "000002000020142549460000100001"
    },
    "payment": {
        "id": "000002000020142549460000100001",
        "paymentOutput": {
            "amountOfMoney": {
                "amount": 2345,
                "currencyCode": "CAD"
            },
            "references": {
                "paymentReference": "0"
            },
            "paymentMethod": "card",
            "cardPaymentMethodSpecificOutput": {
                "paymentProductId": 1,
                "authorisationCode": "OK1131",
                "card": {
                    "cardNumber": "************9176",
                    "expiryDate": "1220"
                },
                "fraudResults": {
                    "fraudServiceResult": "error",
                    "avsResult": "X",
                    "cvvResult": "M"
                }
            }
        },
        "status": "PENDING_APPROVAL",
        "statusOutput": {
            "isCancellable": true,
            "statusCode": 600,
            "statusCodeChangeDateTime": "20150331120036",
            "isAuthorized": true
        }
    }
}`

const rejectedJSON = `{
    "errorId": "2c164323-20d3-4e9e-8578-dc562cd7506c-0000003c",
    "errors": [
        {
            "code": "21000020",
            "requestId": "2001798",
            "message": "VALUE **************** OF FIELD CREDITCARDNUMBER DID NOT PASS THE LUHNCHECK"
        }
    ],
    "paymentResult": {
        "creationOutput": {
            "additionalReference": "00000200002014254436",
            "externalReference": "000002000020142544360000100001"
        },
        "payment": {
            "id": "000002000020142544360000100001",
            "paymentOutput": {
                "amountOfMoney": {
                    "amount": 2345,
                    "currencyCode": "CAD"
                },
                "references": {
                    "paymentReference": "0"
                },
                "paymentMethod": "card",
                "cardPaymentMethodSpecificOutput": {
                    "paymentProductId": 1
                }
            },
            "status": "REJECTED",
            "statusOutput": {
                "errors": [
                    {
                        "code": "21000020",
                        "requestId": "2001798",
                        "message": "VALUE **************** OF FIELD CREDITCARDNUMBER DID NOT PASS THE LUHNCHECK"
                    }
                ],
                "isCancellable": false,
                "statusCode": 100,
                "statusCodeChangeDateTime": "20150330173151",
                "isAuthorized": false
            }
        }
    }
}`

const invalidRequestJSON string = `{
    "errorId": "2c164323-20d3-4e9e-8578-dc562cd7506c-00000058",
    "errors": [
        {
            "code": "21000120",
            "requestId": "2001803",
            "propertyName": "cardPaymentMethodSpecificInput.card.expiryDate",
            "message": "paymentMethodSpecificInput.card.expiryDate (1210) IS IN THE PAST OR NOT IN CORRECT MMYY FORMAT"
        }
    ]
}`

const invalidAuthorizationJSON string = `{
    "errorId": "fbd8d914-c889-45d3-a396-9e0d9ff9db88-0000006f",
    "errors": [
        {
            "code": "9002",
            "message": "MISSING_OR_INVALID_AUTHORIZATION"
        }
    ]
}`

const duplicateRequestJSON string = `{
   "errorId" : "75b0f13a-04a5-41b3-83b8-b295ddb23439-000013c6",
   "errors" : [ {
      "code" : "1409",
      "message" : "DUPLICATE REQUEST IN PROGRESS",
      "httpStatusCode" : 409
   } ]
}`

const notFoundHTML string = "Not Found"

const methodNotAllowedHTML string = "Not Allowed"

type CheckResult func(connection *TestConnection, response domain.CreatePaymentResponse, err error) string

func (t *TestConnection) CloseExpiredConnections() {
}
func (t *TestConnection) CloseIdleConnections(num time.Duration) {
}
func (t *TestConnection) Close() error {
	return nil
}
func (t *TestConnection) Get(resourceURI url.URL, requestHeaders []communication.Header, handler communication.ResponseHandler) (interface{}, error) {
	return nil, nil
}
func (t *TestConnection) Delete(resourceURI url.URL, requestHeaders []communication.Header, handler communication.ResponseHandler) (interface{}, error) {
	return nil, nil
}
func (t *TestConnection) Put(resourceURI url.URL, requestHeaders []communication.Header, body string, handler communication.ResponseHandler) (interface{}, error) {
	return nil, nil
}
func (t *TestConnection) PutMultipart(resourceURI url.URL, requestHeaders []communication.Header, body *communication.MultipartFormDataObject, handler communication.ResponseHandler) (interface{}, error) {
	return nil, nil
}
func (t *TestConnection) Post(resourceURI url.URL, requestHeaders []communication.Header, body string, handler communication.ResponseHandler) (interface{}, error) {
	return handler(t.statusCode, t.headers, strings.NewReader(t.body))
}
func (t *TestConnection) PostMultipart(resourceURI url.URL, requestHeaders []communication.Header, body *communication.MultipartFormDataObject, handler communication.ResponseHandler) (interface{}, error) {
	return nil, nil
}
func (t *TestConnection) DisableLogging() {
}
func (t *TestConnection) EnableLogging(logger logging.CommunicatorLogger) {
}
func (t *TestConnection) SetBodyObfuscator(obfuscator obfuscation.BodyObfuscator) {
}
func (t *TestConnection) SetHeaderObfuscator(obfuscator obfuscation.HeaderObfuscator) {
}

func TestPayment(t *testing.T) {
	for _, connection := range table {
		apiEndpoint, _ := url.Parse("http://localhost")

		authenticator, _ := v1hmac.NewAuthenticator("test", "test")

		metadataProvider, err := communicator.NewMetadataProvider("Worldline")
		if err != nil {
			t.Fatal("Cannot create MetadataProvider")
		}

		client, err := CreateClientWithDefaultMarshaller(apiEndpoint, &connection, authenticator, metadataProvider)
		if err != nil {
			t.Fatal("Cannot create Client: ", err)
		}
		body := domain.CreatePaymentRequest{}
		cc := NewCallContext("")
		var response domain.CreatePaymentResponse
		if connection.idempotenceKey != "" {
			cc.IdempotenceKey = connection.idempotenceKey
			response, err = client.V1().Merchant("merchantId").Payments().Create(body, cc)
		} else {
			response, err = client.V1().Merchant("merchantId").Payments().Create(body, nil)
		}
		if str := connection.checkResult(&connection, response, err); str != "" {
			t.Fatal(str)
		}
	}
}
