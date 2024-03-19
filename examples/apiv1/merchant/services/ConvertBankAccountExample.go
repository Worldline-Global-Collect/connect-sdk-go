// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func convertBankAccountExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer func(client *connectsdk.Client) {
		_ = client.Close()
	}(client)

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper function connectsdk.NewString to overcome this issue.
	// This helper function is provided by the SDK's root package.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.

	var bankAccountBban domain.BankAccountBban
	bankAccountBban.AccountNumber = connectsdk.NewString("0532013000")
	bankAccountBban.BankCode = connectsdk.NewString("37040044")
	bankAccountBban.CountryCode = connectsdk.NewString("DE")

	var body domain.BankDetailsRequest
	body.BankAccountBban = &bankAccountBban

	response, err := client.V1().Merchant("merchantId").Services().Bankaccount(body, nil)

	fmt.Println(response, err)
}
