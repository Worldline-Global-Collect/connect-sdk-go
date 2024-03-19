// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/products"
)

func getDirectoryExample() {
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

	var query products.DirectoryParams
	query.CountryCode = connectsdk.NewString("NL")
	query.CurrencyCode = connectsdk.NewString("EUR")

	response, err := client.V1().Merchant("merchantId").Products().Directory(809, query, nil)

	fmt.Println(response, err)
}
