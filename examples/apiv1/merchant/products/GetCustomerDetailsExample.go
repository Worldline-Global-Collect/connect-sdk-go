// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func getCustomerDetailsExample() {
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

	var values []domain.KeyValuePair

	var value1 domain.KeyValuePair
	value1.Key = connectsdk.NewString("fiscalNumber")
	value1.Value = connectsdk.NewString("01234567890")

	values = append(values, value1)

	var body domain.GetCustomerDetailsRequest
	body.CountryCode = connectsdk.NewString("SE")
	body.Values = &values

	response, err := client.V1().Merchant("merchantId").Products().CustomerDetails(9000, body, nil)

	fmt.Println(response, err)
}
