package connectsdk

import (
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/services"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/tokens"
	commErrors "github.com/Worldline-Global-Collect/connect-sdk-go/communicator/errors"
	"github.com/Worldline-Global-Collect/connect-sdk-go/logging"
)

var testConnectionJSON = `{
	"result": "OK"
}`

var convertAmountJSON = `{
	"convertedAmount": 4547504
}`

var createPaymentUnicodeJSON = `{
    "creationOutput": {
        "additionalReference": "00000012341000059598",
        "externalReference": "000000123410000595980000100001"
    },
    "payment": {
        "id": "000000123410000595980000100001",
        "paymentOutput": {
            "amountOfMoney": {
                "amount": 2345,
                "currencyCode": "CAD"
            },
            "references": {
                "paymentReference": "0"
            },
            "paymentMethod": "redirect",
            "redirectPaymentMethodSpecificOutput":{
               "paymentProductId":840,
               "paymentProduct840SpecificOutput":{
                  "customerAccount":{
                     "firstName":"Theresa",
                     "surname":"Schröder"
                  },
                  "customerAddress":{
                     "city":"sittensen",
                     "countryCode":"DE",
                     "street":"Westerberg 25",
                     "zip":"27419"
                  }
               }
            }
        },
        "status": "PENDING_APPROVAL",
        "statusOutput": {
            "isCancellable": true,
            "statusCategory": "PENDING_MERCHANT",
            "statusCode": 600,
            "statusCodeChangeDateTime": "20160310094054",
            "isAuthorized": true
        }
    }
}
`

var createPaymentJSON = `{
	"creationOutput": {
		"additionalReference": "00000012341000059598",
		"externalReference": "000000123410000595980000100001"
	},
	"payment": {
		"id": "000000123410000595980000100001",
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
					"cardNumber": "************3456",
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
			"statusCategory": "PENDING_MERCHANT",
			"statusCode": 600,
			"statusCodeChangeDateTime": "20160310094054",
			"isAuthorized": true
		}
	}
}`

var createPaymentFailedInvalidCardNumberJSON = `{
	"errorId": "0953f236-9e54-4f23-9556-d66bc757dda8",
	"errors": [{
		"code": "21000020",
		"requestId": "24146",
		"message": "VALUE **************** OF FIELD CREDITCARDNUMBER DID NOT PASS THE LUHNCHECK",
		"httpStatusCode": 400
	}]
}`

var createPaymentFailedRejectedJSON = `{
	"errorId": "833dfd83-52ae-419c-b871-9df1278da93e",
	"errors": [{
		"code": "430330",
		"message": "Not authorised",
		"httpStatusCode": 402
	}],
	"paymentResult": {
		"creationOutput": {
			"additionalReference": "00000200001000059614",
			"externalReference": "000002000010000596140000100001"
		},
		"payment": {
			"id": "000002000010000596140000100001",
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
				"errors": [{
					"code": "430330",
					"requestId": "24162",
					"message": "Not authorised",
					"httpStatusCode": 402
				}],
				"isCancellable": false,
				"statusCategory": "UNSUCCESSFUL",
				"statusCode": 100,
				"statusCodeChangeDateTime": "20160310121151",
				"isAuthorized": false
			}
		}
	}
}`

var unknownServerErrorJSON = `{
	"errorId": "fbff1179-7ba4-4894-9021-d8a0011d23a7",
	"errors": [{
		"code": "9999",
		"message": "UNKNOWN_SERVER_ERROR",
		"httpStatusCode": 500
	}]
}`

var notFoundErrorHTML = `Not Found`

func TestLoggingTestConnection(t *testing.T) {
	logPrefix := "TestLoggingTestConnection"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/services/testconnection",
		createRecordRequest(http.StatusOK, testConnectionJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, client *Client) {
		_ = listener.Close()
		_ = sl.Close()
		_ = client.Close()
	}(listener, sl, client)

	logger := &testLogger{}
	client.EnableLogging(logger)

	response, err := client.V1().Merchant("1234").Services().Testconnection(nil)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if response.Result != nil && *response.Result != "OK" {
		t.Fatalf("%v: responseResult %v", logPrefix, response.Result)
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'GET'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/1234/services/testconnection'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-GCS-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if strings.Contains(requestEntry.message, "body:") {
		t.Fatalf("%v: requeestEntryRequestBody %v", logPrefix, requestEntry.message)
	}

	responseEntry := logger.entries[1]
	if !strings.Contains(responseEntry.message, "status-code:  '200'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'application/json'") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "body:         '") || strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
}

func TestLoggingConvertAmount(t *testing.T) {
	logPrefix := "TestLoggingConvertAmount"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/services/convert/amount",
		createRecordRequest(http.StatusOK, convertAmountJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, client *Client) {
		_ = listener.Close()
		_ = sl.Close()
		_ = client.Close()
	}(listener, sl, client)

	logger := &testLogger{}
	client.EnableLogging(logger)

	var query services.ConvertAmountParams
	amount := int64(1000)
	query.Amount = &amount
	source := "EUR"
	query.Source = &source
	target := "USD"
	query.Target = &target

	response, err := client.V1().Merchant("1234").Services().ConvertAmount(query, nil)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if response.ConvertedAmount == nil {
		t.Fatalf("%v: responseResult %v", logPrefix, response.ConvertedAmount)
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'GET'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/1234/services/convert/amount'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-GCS-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if strings.Contains(requestEntry.message, "body:") {
		t.Fatalf("%v: requestEntryRequestBody %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}

	responseEntry := logger.entries[1]
	if !strings.Contains(responseEntry.message, "status-code:  '200'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'application/json'") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "body:         '") || strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
}

func TestDeleteToken(t *testing.T) {
	logPrefix := "TestDeleteToken"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/tokens/5678",
		createRecordRequest(http.StatusNoContent, "", responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, client *Client) {
		_ = listener.Close()
		_ = sl.Close()
		_ = client.Close()
	}(listener, sl, client)

	logger := &testLogger{}
	client.EnableLogging(logger)

	var query tokens.DeleteParams

	err = client.V1().Merchant("1234").Tokens().Delete("5678", query, nil)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'DELETE'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/1234/tokens/5678'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-GCS-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if strings.Contains(requestEntry.message, "body:") {
		t.Fatalf("%v: requestEntryRequestBody %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}

	responseEntry := logger.entries[1]
	if !strings.Contains(responseEntry.message, "status-code:  '204'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'application/json'") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
	if !strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
}

func TestLoggingCreatePaymentUnicode(t *testing.T) {
	logPrefix := "TestLoggingCreatePayment"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
		"Location":     "http://localhost/v1/1234/payments/000000123410000595980000100001",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/payments",
		createRecordRequest(http.StatusCreated, createPaymentUnicodeJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, client *Client) {
		_ = listener.Close()
		_ = sl.Close()
		_ = client.Close()
	}(listener, sl, client)

	logger := &testLogger{}
	client.EnableLogging(logger)

	var card domain.Card
	card.CardNumber = NewString("1234567890123456")
	card.Cvv = NewString("123")
	card.ExpiryDate = NewString("1220")

	var cardPaymentMethodSpecificInput domain.CardPaymentMethodSpecificInput
	cardPaymentMethodSpecificInput.Card = &card
	cardPaymentMethodSpecificInput.PaymentProductID = NewInt32(1)

	var amountOfMoney domain.AmountOfMoney
	amountOfMoney.Amount = NewInt64(2345)
	amountOfMoney.CurrencyCode = NewString("CAD")

	var billingAddress domain.Address
	billingAddress.CountryCode = NewString("CA")

	var customer domain.Customer
	customer.BillingAddress = &billingAddress

	var order domain.Order
	order.AmountOfMoney = &amountOfMoney
	order.Customer = &customer

	var requestBody domain.CreatePaymentRequest
	requestBody.CardPaymentMethodSpecificInput = &cardPaymentMethodSpecificInput
	requestBody.Order = &order

	response, err := client.V1().Merchant("1234").Payments().Create(requestBody, nil)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	if response.Payment == nil {
		t.Fatalf("%v: responsePayment nil", logPrefix)
	}
	if response.Payment.ID == nil {
		t.Fatalf("%v: responsePaymentID nil", logPrefix)
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'POST'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/1234/payments'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-GCS-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "body:         '") || strings.Contains(requestEntry.message, "body:         ''") {
		t.Fatalf("%v: requestEntryRequestBody %v", logPrefix, requestEntry.message)
	}
	if strings.Contains(requestEntry.message, "\"1234567890123456\"") || strings.Contains(requestEntry.message, "\"123\"") || strings.Contains(requestEntry.message, "\"1220\"") {
		t.Fatalf("%v: requestEntryRequestBodyObfuscation %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}

	responseEntry := logger.entries[1]
	if !strings.Contains(responseEntry.message, "Schröder") {
		t.Fatalf("%v: responseEntryResponse %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "status-code:  '201'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'application/json'") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "body:         '") || strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
}
func TestLoggingCreatePayment(t *testing.T) {
	logPrefix := "TestLoggingCreatePayment"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
		"Location":     "http://localhost/v1/1234/payments/000000123410000595980000100001",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/payments",
		createRecordRequest(http.StatusCreated, createPaymentJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, client *Client) {
		_ = listener.Close()
		_ = sl.Close()
		_ = client.Close()
	}(listener, sl, client)

	logger := &testLogger{}
	client.EnableLogging(logger)

	var card domain.Card
	card.CardNumber = NewString("1234567890123456")
	card.Cvv = NewString("123")
	card.ExpiryDate = NewString("1220")

	var cardPaymentMethodSpecificInput domain.CardPaymentMethodSpecificInput
	cardPaymentMethodSpecificInput.Card = &card
	cardPaymentMethodSpecificInput.PaymentProductID = NewInt32(1)

	var amountOfMoney domain.AmountOfMoney
	amountOfMoney.Amount = NewInt64(2345)
	amountOfMoney.CurrencyCode = NewString("CAD")

	var billingAddress domain.Address
	billingAddress.CountryCode = NewString("CA")

	var customer domain.Customer
	customer.BillingAddress = &billingAddress

	var order domain.Order
	order.AmountOfMoney = &amountOfMoney
	order.Customer = &customer

	var requestBody domain.CreatePaymentRequest
	requestBody.CardPaymentMethodSpecificInput = &cardPaymentMethodSpecificInput
	requestBody.Order = &order

	response, err := client.V1().Merchant("1234").Payments().Create(requestBody, nil)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	if response.Payment == nil {
		t.Fatalf("%v: responsePayment nil", logPrefix)
	}
	if response.Payment.ID == nil {
		t.Fatalf("%v: responsePaymentID nil", logPrefix)
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'POST'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/1234/payments'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-GCS-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "body:         '") || strings.Contains(requestEntry.message, "body:         ''") {
		t.Fatalf("%v: requestEntryRequestBody %v", logPrefix, requestEntry.message)
	}
	if strings.Contains(requestEntry.message, "\"1234567890123456\"") || strings.Contains(requestEntry.message, "\"123\"") || strings.Contains(requestEntry.message, "\"1220\"") {
		t.Fatalf("%v: requestEntryRequestBodyObfuscation %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}

	responseEntry := logger.entries[1]
	if !strings.Contains(responseEntry.message, "status-code:  '201'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'application/json'") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "body:         '") || strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
}

func TestCreatePaymentInvalidCardNumber(t *testing.T) {
	logPrefix := "TestCreatePaymentInvalidCardNumber"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/payments",
		createRecordRequest(http.StatusBadRequest, createPaymentFailedInvalidCardNumberJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, client *Client) {
		_ = listener.Close()
		_ = sl.Close()
		_ = client.Close()
	}(listener, sl, client)

	logger := &testLogger{}
	client.EnableLogging(logger)

	var card domain.Card
	card.CardNumber = NewString("1234567890123456")
	card.Cvv = NewString("123")
	card.ExpiryDate = NewString("1220")

	var cardPaymentMethodSpecificInput domain.CardPaymentMethodSpecificInput
	cardPaymentMethodSpecificInput.Card = &card
	cardPaymentMethodSpecificInput.PaymentProductID = NewInt32(1)

	var amountOfMoney domain.AmountOfMoney
	amountOfMoney.Amount = NewInt64(2345)
	amountOfMoney.CurrencyCode = NewString("CAD")

	var billingAddress domain.Address
	billingAddress.CountryCode = NewString("CA")

	var customer domain.Customer
	customer.BillingAddress = &billingAddress

	var order domain.Order
	order.AmountOfMoney = &amountOfMoney
	order.Customer = &customer

	var requestBody domain.CreatePaymentRequest
	requestBody.CardPaymentMethodSpecificInput = &cardPaymentMethodSpecificInput
	requestBody.Order = &order

	_, err = client.V1().Merchant("1234").Payments().Create(requestBody, nil)
	switch ce := err.(type) {
	case *v1Errors.ValidationError:
		{
			if ce.StatusCode() != http.StatusBadRequest {
				t.Fatalf("%v: statusCode %v", logPrefix, ce.StatusCode())
			}
			if ce.ResponseBody() != createPaymentFailedInvalidCardNumberJSON {
				t.Fatalf("%v: responseBody %v", logPrefix, ce.ResponseBody())
			}

			break
		}
	default:
		{
			t.Fatalf("%v: %v", logPrefix, err)
		}
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'POST'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/1234/payments'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-GCS-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "body:         '") || strings.Contains(requestEntry.message, "body:         ''") {
		t.Fatalf("%v: requestEntryRequestBody %v", logPrefix, requestEntry.message)
	}
	if strings.Contains(requestEntry.message, "\"1234567890123456\"") || strings.Contains(requestEntry.message, "\"123\"") || strings.Contains(requestEntry.message, "\"1220\"") {
		t.Fatalf("%v: requestEntryRequestBodyObfuscation %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}

	responseEntry := logger.entries[1]
	if !strings.Contains(responseEntry.message, "status-code:  '400'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'application/json'") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "body:         '") || strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
}

func TestCreatePaymentRejected(t *testing.T) {
	logPrefix := "TestCreatePaymentRejected"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/payments",
		createRecordRequest(http.StatusPaymentRequired, createPaymentFailedRejectedJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, client *Client) {
		_ = listener.Close()
		_ = sl.Close()
		_ = client.Close()
	}(listener, sl, client)

	logger := &testLogger{}
	client.EnableLogging(logger)

	var card domain.Card
	card.CardNumber = NewString("1234567890123456")
	card.Cvv = NewString("123")
	card.ExpiryDate = NewString("1220")

	var cardPaymentMethodSpecificInput domain.CardPaymentMethodSpecificInput
	cardPaymentMethodSpecificInput.Card = &card
	cardPaymentMethodSpecificInput.PaymentProductID = NewInt32(1)

	var amountOfMoney domain.AmountOfMoney
	amountOfMoney.Amount = NewInt64(2345)
	amountOfMoney.CurrencyCode = NewString("CAD")

	var billingAddress domain.Address
	billingAddress.CountryCode = NewString("CA")

	var customer domain.Customer
	customer.BillingAddress = &billingAddress

	var order domain.Order
	order.AmountOfMoney = &amountOfMoney
	order.Customer = &customer

	var requestBody domain.CreatePaymentRequest
	requestBody.CardPaymentMethodSpecificInput = &cardPaymentMethodSpecificInput
	requestBody.Order = &order

	_, err = client.V1().Merchant("1234").Payments().Create(requestBody, nil)
	switch ce := err.(type) {
	case *v1Errors.DeclinedPaymentError:
		{
			if ce.StatusCode() != http.StatusPaymentRequired {
				t.Fatalf("%v: statusCode %v", logPrefix, ce.StatusCode())
			}
			if ce.ResponseBody() != createPaymentFailedRejectedJSON {
				t.Fatalf("%v: responseBody %v", logPrefix, ce.ResponseBody())
			}

			break
		}
	default:
		{
			t.Fatalf("%v: %v", logPrefix, err)
		}
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'POST'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/1234/payments'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-GCS-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "body:         '") || strings.Contains(requestEntry.message, "body:         ''") {
		t.Fatalf("%v: requestEntryRequestBody %v", logPrefix, requestEntry.message)
	}
	if strings.Contains(requestEntry.message, "\"1234567890123456\"") || strings.Contains(requestEntry.message, "\"123\"") || strings.Contains(requestEntry.message, "\"1220\"") {
		t.Fatalf("%v: requestEntryRequestBodyObfuscation %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}

	responseEntry := logger.entries[1]
	if !strings.Contains(responseEntry.message, "status-code:  '402'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'application/json'") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "body:         '") || strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
}

func TestLoggingUnknownServerError(t *testing.T) {
	logPrefix := "TestLoggingUnknownServerError"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/services/testconnection",
		createRecordRequest(http.StatusInternalServerError, unknownServerErrorJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, client *Client) {
		_ = listener.Close()
		_ = sl.Close()
		_ = client.Close()
	}(listener, sl, client)

	logger := &testLogger{}
	client.EnableLogging(logger)

	_, err = client.V1().Merchant("1234").Services().Testconnection(nil)
	switch ce := err.(type) {
	case *v1Errors.PlatformError:
		{
			if ce.StatusCode() != http.StatusInternalServerError {
				t.Fatalf("%v: statusCode %v", logPrefix, ce.StatusCode())
			}
			if ce.ResponseBody() != unknownServerErrorJSON {
				t.Fatalf("%v: responseBody %v", logPrefix, ce.ResponseBody())
			}

			break
		}
	default:
		{
			t.Fatalf("%v: %v", logPrefix, err)
		}
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'GET'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/1234/services/testconnection'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-GCS-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}

	responseEntry := logger.entries[1]
	if !strings.Contains(responseEntry.message, "status-code:  '500'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'application/json'") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "body:         '") || strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
}

func TestNonJson(t *testing.T) {
	logPrefix := "TestNonJson"

	responseHeaders := map[string]string{
		"Dummy": "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/services/testconnection",
		createRecordRequest(http.StatusNotFound, notFoundErrorHTML, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, client *Client) {
		_ = listener.Close()
		_ = sl.Close()
		_ = client.Close()
	}(listener, sl, client)

	logger := &testLogger{}
	client.EnableLogging(logger)

	_, err = client.V1().Merchant("1234").Services().Testconnection(nil)
	switch err.(type) {
	case *commErrors.NotFoundError:
		{
			break
		}
	default:
		{
			t.Fatalf("%v: %v", logPrefix, err)
		}
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'GET'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/1234/services/testconnection'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-GCS-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}

	responseEntry := logger.entries[1]
	if !strings.Contains(responseEntry.message, "status-code:  '404'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'text/plain") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "body:         '") || strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
}

func TestReadTimeout(t *testing.T) {
	logPrefix := "TestReadTimeout"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTimedTestEnvironment(
		"/v1/1234/services/testconnection",
		createDelayedRecordRequest(http.StatusOK, testConnectionJSON, responseHeaders, requestHeaders, 1*time.Second),
		1*time.Millisecond,
		10*time.Millisecond)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, client *Client) {
		_ = listener.Close()
		_ = sl.Close()
		_ = client.Close()
	}(listener, sl, client)

	logger := &testLogger{}
	client.EnableLogging(logger)

	_, err = client.V1().Merchant("1234").Services().Testconnection(nil)
	switch ce := err.(type) {
	case *commErrors.CommunicationError:
		{
			internalError := ce.InternalError()

			if uErr, ok := internalError.(*url.Error); ok && uErr.Timeout() {
				break
			}

			t.Fatalf("%v: %v", logPrefix, internalError)
		}
	default:
		{
			t.Fatalf("%v: %v", logPrefix, err)
		}
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'GET'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/1234/services/testconnection'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-GCS-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}

	errorEntry := logger.entries[1]
	if errorEntry.err == nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, errorEntry.err)
	}
}

func TestLogRequestOnly(t *testing.T) {
	logPrefix := "TestLogRequestOnly"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, mux, err := createEmptyTestEnvironment()
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, client *Client) {
		_ = listener.Close()
		_ = sl.Close()
		_ = client.Close()
	}(listener, sl, client)

	mux.HandleFunc("/v1/1234/services/testconnection",
		createNonLoggedRecordRequest(http.StatusOK, testConnectionJSON, responseHeaders, requestHeaders, client))

	logger := &testLogger{}
	client.EnableLogging(logger)

	_, err = client.V1().Merchant("1234").Services().Testconnection(nil)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if len(logger.entries) != 1 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'GET'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/1234/services/testconnection'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-GCS-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}
}

func TestLogResponseOnly(t *testing.T) {
	logPrefix := "TestLogResponseOnly"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, mux, err := createEmptyTestEnvironment()
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, client *Client) {
		_ = listener.Close()
		_ = sl.Close()
		_ = client.Close()
	}(listener, sl, client)

	logger := &testLogger{}

	mux.HandleFunc("/v1/1234/services/testconnection",
		createLoggedRecordRequest(http.StatusOK, testConnectionJSON, responseHeaders, requestHeaders, client, logger))

	_, err = client.V1().Merchant("1234").Services().Testconnection(nil)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if len(logger.entries) != 1 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	responseEntry := logger.entries[0]
	if !strings.Contains(responseEntry.message, "status-code:  '200'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'application/json'") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "body:         '") || strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
}

func TestLogErrorOnly(t *testing.T) {
	logPrefix := "TestLogErrorOnly"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, mux, err := createEmptyTimedTestEnvironment(
		100*time.Millisecond,
		100*time.Millisecond)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, client *Client) {
		_ = listener.Close()
		_ = sl.Close()
		_ = client.Close()
	}(listener, sl, client)

	logger := &testLogger{}

	mux.HandleFunc("/v1/1234/services/testconnection",
		createLoggedDelayedRecordRequest(http.StatusOK, testConnectionJSON, responseHeaders, requestHeaders, 1*time.Second, client, logger))

	_, err = client.V1().Merchant("1234").Services().Testconnection(nil)
	switch ce := err.(type) {
	case *commErrors.CommunicationError:
		{
			internalError := ce.InternalError()

			if uErr, ok := internalError.(*url.Error); ok && uErr.Timeout() {
				break
			}

			t.Fatalf("%v: %v", logPrefix, internalError)
		}
	default:
		{
			t.Fatalf("%v: %v", logPrefix, err)
		}
	}

	if len(logger.entries) != 1 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	errorEntry := logger.entries[0]
	if errorEntry.err == nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, errorEntry.err)
	}
}

func createNonLoggedRecordRequest(statusCode int, body string, responseHeaders map[string]string, requestHeaders map[string][]string, client *Client) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		client.DisableLogging()

		for k, v := range r.Header {
			if k == "X-Gcs-Idempotence-Key" {
				k = "X-GCS-Idempotence-Key"
			}

			requestHeaders[k] = v
		}

		for k, v := range responseHeaders {
			rw.Header()[k] = []string{v}
		}

		rw.WriteHeader(statusCode)

		_, _ = rw.Write([]byte(body))
	}
}

func createLoggedRecordRequest(statusCode int, body string, responseHeaders map[string]string, requestHeaders map[string][]string, client *Client, logger logging.CommunicatorLogger) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		client.EnableLogging(logger)

		for k, v := range r.Header {
			if k == "X-Gcs-Idempotence-Key" {
				k = "X-GCS-Idempotence-Key"
			}

			requestHeaders[k] = v
		}

		for k, v := range responseHeaders {
			rw.Header()[k] = []string{v}
		}

		rw.WriteHeader(statusCode)

		_, _ = rw.Write([]byte(body))
	}
}

func createDelayedRecordRequest(statusCode int, body string, responseHeaders map[string]string, requestHeaders map[string][]string, delay time.Duration) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)

		for k, v := range r.Header {
			if k == "X-Gcs-Idempotence-Key" {
				k = "X-GCS-Idempotence-Key"
			}

			requestHeaders[k] = v
		}

		for k, v := range responseHeaders {
			rw.Header()[k] = []string{v}
		}

		rw.WriteHeader(statusCode)

		_, _ = rw.Write([]byte(body))
	}
}

func createLoggedDelayedRecordRequest(statusCode int, body string, responseHeaders map[string]string, requestHeaders map[string][]string, delay time.Duration, client *Client, logger logging.CommunicatorLogger) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		client.EnableLogging(logger)
		time.Sleep(delay)

		for k, v := range r.Header {
			if k == "X-Gcs-Idempotence-Key" {
				k = "X-GCS-Idempotence-Key"
			}

			requestHeaders[k] = v
		}

		for k, v := range responseHeaders {
			rw.Header()[k] = []string{v}
		}

		rw.WriteHeader(statusCode)

		_, _ = rw.Write([]byte(body))
	}
}

func createTimedTestEnvironment(path string, handleFunc http.HandlerFunc, socketTimeout, connectTimeout time.Duration) (net.Listener, *stoppableListener, *Client, error) {
	mux := http.NewServeMux()
	mux.Handle(path, handleFunc)

	httpServer := &http.Server{
		Handler: mux,
	}

	randomPort := (1 << 12) + rand.Intn(1<<15)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(randomPort))
	if err != nil {
		return nil, nil, nil, err
	}

	sl, err := mockServer(httpServer, listener)
	if err != nil {
		return nil, nil, nil, err
	}

	client, err := createClient(socketTimeout, connectTimeout, randomPort)
	if err != nil {
		return nil, nil, nil, err
	}

	return listener, sl, client, nil
}

func createEmptyTestEnvironment() (net.Listener, *stoppableListener, *Client, *http.ServeMux, error) {
	mux := http.NewServeMux()

	httpServer := &http.Server{
		Handler: mux,
	}

	randomPort := (1 << 12) + rand.Intn(1<<15)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(randomPort))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	sl, err := mockServer(httpServer, listener)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	client, err := createClient(50*time.Second, 50*time.Second, randomPort)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return listener, sl, client, mux, nil
}

func createEmptyTimedTestEnvironment(socketTimeout, connectTimeout time.Duration) (net.Listener, *stoppableListener, *Client, *http.ServeMux, error) {
	mux := http.NewServeMux()

	httpServer := &http.Server{
		Handler: mux,
	}

	randomPort := (1 << 12) + rand.Intn(1<<15)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(randomPort))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	sl, err := mockServer(httpServer, listener)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	client, err := createClient(socketTimeout, connectTimeout, randomPort)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return listener, sl, client, mux, nil
}

type testLogger struct {
	entries []testLoggerEntry
}

func (t *testLogger) Log(message string) {
	t.entries = append(t.entries, testLoggerEntry{message, nil})
}

func (t *testLogger) LogError(message string, err error) {
	t.entries = append(t.entries, testLoggerEntry{message, err})
}

type testLoggerEntry struct {
	message string
	err     error
}
