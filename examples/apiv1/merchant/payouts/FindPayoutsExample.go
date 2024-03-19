// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/payouts"
)

func findPayoutsExample() {
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

	var query payouts.FindParams
	query.MerchantReference = connectsdk.NewString("AcmeOrder0001")
	query.MerchantOrderID = connectsdk.NewInt64(123456)
	query.Offset = connectsdk.NewInt32(0)
	query.Limit = connectsdk.NewInt32(10)

	response, err := client.V1().Merchant("merchantId").Payouts().Find(query, nil)

	fmt.Println(response, err)
}
