// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/services"
)

func convertAmountExample() {
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

	var query services.ConvertAmountParams
	query.Source = connectsdk.NewString("EUR")
	query.Target = connectsdk.NewString("USD")
	query.Amount = connectsdk.NewInt64(100)

	response, err := client.V1().Merchant("merchantId").Services().ConvertAmount(query, nil)

	fmt.Println(response, err)
}
