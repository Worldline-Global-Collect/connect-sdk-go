// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/products"
)

func getPaymentProductsExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer func(client *connectsdk.Client) {
		_ = client.Close()
	}(client)

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions connectsdk.NewBool, connectsdk.NewInt64 and connectsdk.NewString to overcome this issue.
	// These helper functions are provided by the SDK's root package.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.

	var query products.FindParams
	query.CountryCode = connectsdk.NewString("US")
	query.CurrencyCode = connectsdk.NewString("USD")
	query.Locale = connectsdk.NewString("en_US")
	query.Amount = connectsdk.NewInt64(1000)
	query.IsRecurring = connectsdk.NewBool(true)
	query.IsInstallments = connectsdk.NewBool(true)
	query.AddHide("fields")

	response, err := client.V1().Merchant("merchantId").Products().Find(query, nil)

	fmt.Println(response, err)
}
