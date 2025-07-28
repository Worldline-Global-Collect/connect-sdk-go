package connectsdk

import (
	"crypto/rand"
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/productgroups"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/products"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/services"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/tokens"
	"github.com/Worldline-Global-Collect/connect-sdk-go/authentication/v1hmac"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator"
	"github.com/Worldline-Global-Collect/connect-sdk-go/configuration"
)

var envMerchantID = os.Getenv("connect.api.merchantId")
var envAPIKeyID = os.Getenv("connect.api.apiKeyId")
var envSecretAPIKey = os.Getenv("connect.api.secretApiKey")
var envEndpointURL = os.Getenv("connect.api.endpointUrl")
var envProxyURL = os.Getenv("connect.api.proxyUrl")

func TestIntegratedConvertAmount(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer func(client *Client) {
		_ = client.Close()
	}(client)

	var query services.ConvertAmountParams
	query.Amount = NewInt64(123)
	query.Source = NewString("USD")
	query.Target = NewString("EUR")

	response, err := client.V1().Merchant(envMerchantID).Services().ConvertAmount(query, nil)
	if err != nil {
		t.Fatal(err)
	}

	if response.ConvertedAmount == nil {
		t.Fatal("nil converted amount")
	}
}

func TestIntegratedConnection(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer func(client *Client) {
		_ = client.Close()
	}(client)

	response, err := client.V1().Merchant(envMerchantID).Services().Testconnection(nil)
	if err != nil {
		t.Fatal(err)
	}

	if response.Result == nil {
		t.Fatal("nil result")
	}
}

func TestIntegratedPayment(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer func(client *Client) {
		_ = client.Close()
	}(client)

	var request domain.CreatePaymentRequest
	var order domain.Order

	var amountOfMoney domain.AmountOfMoney
	amountOfMoney.Amount = NewInt64(100)
	amountOfMoney.CurrencyCode = NewString("EUR")
	order.AmountOfMoney = &amountOfMoney

	var customer domain.Customer
	customer.Locale = NewString("en")

	var address domain.Address
	address.CountryCode = NewString("NL")
	customer.BillingAddress = &address

	order.Customer = &customer
	request.Order = &order

	var paymentMethodSpecificInput domain.RedirectPaymentMethodSpecificInput
	paymentMethodSpecificInput.ReturnURL = NewString("http://example.com/")
	paymentMethodSpecificInput.PaymentProductID = NewInt32(809)

	var paymentProductSpecificInput domain.RedirectPaymentProduct809SpecificInput
	paymentProductSpecificInput.IssuerID = NewString("INGBNL2A")
	paymentMethodSpecificInput.PaymentProduct809SpecificInput = &paymentProductSpecificInput

	request.RedirectPaymentMethodSpecificInput = &paymentMethodSpecificInput

	idempotenceKey, err := pseudoUUID()
	if err != nil {
		t.Fatal(err)
	}

	context := NewCallContext(idempotenceKey)

	result, err := doCreatePayment(client, request, context)
	if err != nil {
		t.Fatal(err)
	}
	paymentID := result.Payment.ID
	status := result.Payment.Status

	if idempotenceKey != context.IdempotenceKey {
		t.Fatalf("idempotence key mismatch")
	}
	if context.IdempotenceRequestTimestamp != nil {
		t.Fatalf("timestamp not nil")
	}

	secondResult, err := doCreatePayment(client, request, context)
	if err != nil {
		t.Fatal(err)
	}
	if idempotenceKey != context.IdempotenceKey {
		t.Fatalf("idempotence key mismatch")
	}
	if context.IdempotenceRequestTimestamp == nil {
		t.Fatalf("timestamp nil")
	}

	if *secondResult.Payment.ID != *paymentID {
		t.Fatalf("payment id mismatch")
	}
	if *secondResult.Payment.Status != *status {
		t.Fatalf("status mismatch")
	}
}

func TestIntegratedPaymentProducts(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer func(client *Client) {
		_ = client.Close()
	}(client)

	lParams := products.NewFindParams()
	lParams.CountryCode = NewString("NL")
	lParams.CurrencyCode = NewString("EUR")

	_, err = client.V1().Merchant(envMerchantID).Products().Find(*lParams, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIntegratedPaymentProductDirectories(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer func(client *Client) {
		_ = client.Close()
	}(client)

	lParams := products.NewDirectoryParams()
	lParams.CountryCode = NewString("NL")
	lParams.CurrencyCode = NewString("EUR")

	_, err = client.V1().Merchant(envMerchantID).Products().Directory(809, *lParams, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIntegratedPaymentProductsGroups(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer func(client *Client) {
		_ = client.Close()
	}(client)

	lParams := productgroups.NewGetParams()
	lParams.CountryCode = NewString("NL")
	lParams.CurrencyCode = NewString("EUR")

	_, err = client.V1().Merchant(envMerchantID).Productgroups().Get("cards", *lParams, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIntegratedCreateSession(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer func(client *Client) {
		_ = client.Close()
	}(client)

	lParams := domain.NewSessionRequest()

	_, err = client.V1().Merchant(envMerchantID).Sessions().Create(*lParams, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIntegratedCreateToken(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer func(client *Client) {
		_ = client.Close()
	}(client)

	var request domain.CreateTokenRequest
	request.PaymentProductID = NewInt32(1)

	tokenCard := domain.NewTokenCard()
	request.Card = tokenCard

	customerToken := domain.NewCustomerToken()
	tokenCard.Customer = customerToken

	address := domain.NewAddress()
	customerToken.BillingAddress = address
	address.CountryCode = NewString("NL")

	mandate := domain.NewTokenCardData()
	tokenCard.Data = mandate

	cardWithoutCVV := domain.NewCardWithoutCvv()
	mandate.CardWithoutCvv = cardWithoutCVV
	cardWithoutCVV.CardholderName = NewString("Jan")
	cardWithoutCVV.IssueNumber = NewString("12")
	cardWithoutCVV.CardNumber = NewString("4567350000427977")
	cardWithoutCVV.ExpiryDate = NewString("1225")

	response, err := client.V1().Merchant(envMerchantID).Tokens().Create(request, nil)
	if err != nil {
		t.Fatal(err)
	}

	if response.Token == nil {
		t.Fatal("nil token")
	}

	var query tokens.DeleteParams

	err = client.V1().Merchant(envMerchantID).Tokens().Delete(*response.Token, query, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIntegratedRiskAssessments(t *testing.T) {
	t.Skip("Risk assessments are not available for pre-prod sandbox accounts")
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer func(client *Client) {
		_ = client.Close()
	}(client)

	var request domain.RiskAssessmentBankAccount

	bankAccountBBan := domain.NewBankAccountBban()
	bankAccountBBan.CountryCode = NewString("DE")
	bankAccountBBan.AccountNumber = NewString("0532013000")
	bankAccountBBan.BankCode = NewString("37040044")
	request.BankAccountBban = bankAccountBBan

	order := domain.NewOrderRiskAssessment()

	amountOfMoney := domain.NewAmountOfMoney()
	amountOfMoney.Amount = NewInt64(100)
	amountOfMoney.CurrencyCode = NewString("EUR")
	order.AmountOfMoney = amountOfMoney

	customer := domain.NewCustomerRiskAssessment()
	customer.Locale = NewString("en_GB")
	order.Customer = customer

	request.Order = order

	response, err := client.V1().Merchant(envMerchantID).Riskassessments().Bankaccounts(request, nil)
	if err != nil {
		t.Fatal(err)
	}

	if response.Results == nil {
		t.Fatal("nil results")
	}
	if len(*response.Results) == 0 {
		t.Fatal("empty results")
	}
}

func TestIntegratedMultilineHeader(t *testing.T) {
	//skipTestIfNeeded(t)
	t.Skip("fails inside sandbox")

	conf, _ := configuration.DefaultV1HMACConfiguration(envAPIKeyID, envSecretAPIKey, "Worldline")
	if len(envEndpointURL) == 0 {
		conf.APIEndpoint.Host = "api.preprod.connect.worldline-solutions.com"
	} else {
		endpoint, err := url.Parse(envEndpointURL)
		if err != nil {
			t.Fatal(err)
		}
		conf.APIEndpoint = *endpoint
	}
	conf.SetAPIKeyID(envAPIKeyID)
	conf.SetSecretAPIKey(envSecretAPIKey)

	connection, err := communicator.NewDefaultConnection(conf.SocketTimeout,
		conf.ConnectTimeout,
		conf.KeepAliveTimeout,
		conf.IdleTimeout,
		conf.MaxConnections,
		conf.Proxy)
	if err != nil {
		t.Fatal(err)
	}

	multiLineHeader := newHeader("X-GCS-MultiLineHeader", "some\nvalue")

	metadataProviderBuilder, _ := communicator.NewMetadataProviderBuilder(conf.Integrator)
	metadataProviderBuilder.ShoppingCartExtension = conf.ShoppingCartExtension
	metadataProviderBuilder.AdditionalRequestHeaders = append(metadataProviderBuilder.AdditionalRequestHeaders, multiLineHeader)

	metadataProvider, err := metadataProviderBuilder.Build()
	if err != nil {
		t.Fatal(err)
	}

	authenticator, err := v1hmac.NewAuthenticator(conf.GetAPIKeyID(), conf.GetSecretAPIKey())
	if err != nil {
		t.Fatal(err)
	}

	client, err := CreateClientWithDefaultMarshaller(&conf.APIEndpoint, connection, authenticator, metadataProvider)
	if err != nil {
		t.Fatal(err)
	}
	defer func(client *Client) {
		_ = client.Close()
	}(client)

	_, err = client.V1().Merchant(envMerchantID).Services().Testconnection(nil)
	if err != nil {
		t.Fatal(err)
	}
}

func getClientIntegration() (*Client, error) {
	conf, _ := configuration.DefaultV1HMACConfiguration(envAPIKeyID, envSecretAPIKey, "Worldline")
	if len(envEndpointURL) == 0 {
		conf.APIEndpoint.Host = "api.preprod.connect.worldline-solutions.com"
	} else {
		endpoint, err := url.Parse(envEndpointURL)
		if err != nil {
			return nil, err
		}
		conf.APIEndpoint = *endpoint
	}

	if len(envProxyURL) > 0 {
		proxy, err := url.Parse(envProxyURL)
		if err != nil {
			return nil, err
		}
		conf.Proxy = proxy
	}

	client, err := CreateClientFromConfiguration(conf)
	if err != nil {
		return nil, err
	}

	return client.WithClientMetaInfo("{\"test\":\"test\"}")
}

func doCreatePayment(client *Client, request domain.CreatePaymentRequest, context *communicator.CallContext) (domain.CreatePaymentResult, error) {
	// For this test it doesn't matter if the response is successful or declined,
	// as long as idempotence is handled correctly

	response, err := client.V1().Merchant(envMerchantID).Payments().Create(request, context)
	if err == nil {
		result := domain.CreatePaymentResult{
			CreationOutput: response.CreationOutput,
			MerchantAction: response.MerchantAction,
			Payment:        response.Payment,
		}
		return result, nil
	}
	switch realError := err.(type) {
	case *v1Errors.DeclinedPaymentError:
		return *realError.PaymentResult(), nil
	default:
		return domain.CreatePaymentResult{}, err
	}
}

func pseudoUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

func skipTestIfNeeded(t *testing.T) {
	if len(envMerchantID) == 0 {
		t.Skip("empty env connect.api.merchantId")
	}
	if len(envAPIKeyID) == 0 {
		t.Skip("empty env connect.api.apiKeyId")
	}
	if len(envSecretAPIKey) == 0 {
		t.Skip("empty env connect.api.secretApiKey")
	}
}
