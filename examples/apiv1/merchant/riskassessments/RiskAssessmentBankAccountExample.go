// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func riskAssessmentBankAccountExample() {
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

	var bankAccountBban domain.BankAccountBban
	bankAccountBban.AccountNumber = connectsdk.NewString("0532013000")
	bankAccountBban.BankCode = connectsdk.NewString("37040044")
	bankAccountBban.CountryCode = connectsdk.NewString("DE")

	var amountOfMoney domain.AmountOfMoney
	amountOfMoney.Amount = connectsdk.NewInt64(100)
	amountOfMoney.CurrencyCode = connectsdk.NewString("EUR")

	var billingAddress domain.Address
	billingAddress.CountryCode = connectsdk.NewString("US")

	var customer domain.CustomerRiskAssessment
	customer.BillingAddress = &billingAddress
	customer.Locale = connectsdk.NewString("en_US")

	var order domain.OrderRiskAssessment
	order.AmountOfMoney = &amountOfMoney
	order.Customer = &customer

	var body domain.RiskAssessmentBankAccount
	body.BankAccountBban = &bankAccountBban
	body.Order = &order

	response, err := client.V1().Merchant("merchantId").Riskassessments().Bankaccounts(body, nil)

	fmt.Println(response, err)
}
