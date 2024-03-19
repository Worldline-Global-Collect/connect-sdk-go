// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func getInstallmentsInfoExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer func(client *connectsdk.Client) {
		_ = client.Close()
	}(client)

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions connectsdk.NewInt32, connectsdk.NewInt64 and connectsdk.NewString to overcome this issue.
	// These helper functions are provided by the SDK's root package.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.

	var amountOfMoney domain.AmountOfMoney
	amountOfMoney.Amount = connectsdk.NewInt64(123)
	amountOfMoney.CurrencyCode = connectsdk.NewString("EUR")

	var body domain.GetInstallmentRequest
	body.AmountOfMoney = &amountOfMoney
	body.Bin = connectsdk.NewString("123455")
	body.CountryCode = connectsdk.NewString("NL")
	body.PaymentProductID = connectsdk.NewInt32(123)

	response, err := client.V1().Merchant("merchantId").Installments().GetInstallmentsInfo(body, nil)

	fmt.Println(response, err)
}
