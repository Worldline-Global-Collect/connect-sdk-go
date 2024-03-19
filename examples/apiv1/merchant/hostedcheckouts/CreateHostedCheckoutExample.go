// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func createHostedCheckoutExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer func(client *connectsdk.Client) {
		_ = client.Close()
	}(client)

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions connectsdk.NewInt64 and connectsdk.NewString to overcome this issue.
	// These helper functions are provided by the SDK's root package.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.

	var hostedCheckoutSpecificInput domain.HostedCheckoutSpecificInput
	hostedCheckoutSpecificInput.Locale = connectsdk.NewString("en_GB")
	hostedCheckoutSpecificInput.Variant = connectsdk.NewString("testVariant")

	var amountOfMoney domain.AmountOfMoney
	amountOfMoney.Amount = connectsdk.NewInt64(2345)
	amountOfMoney.CurrencyCode = connectsdk.NewString("USD")

	var billingAddress domain.Address
	billingAddress.CountryCode = connectsdk.NewString("US")

	var customer domain.Customer
	customer.BillingAddress = &billingAddress
	customer.MerchantCustomerID = connectsdk.NewString("1234")

	var order domain.Order
	order.AmountOfMoney = &amountOfMoney
	order.Customer = &customer

	var body domain.CreateHostedCheckoutRequest
	body.HostedCheckoutSpecificInput = &hostedCheckoutSpecificInput
	body.Order = &order

	response, err := client.V1().Merchant("merchantId").Hostedcheckouts().Create(body, nil)

	fmt.Println(response, err)
}
