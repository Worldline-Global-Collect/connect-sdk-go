// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func approveSepaDirectDebitTokenExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer func(client *connectsdk.Client) {
		_ = client.Close()
	}(client)

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions connectsdk.NewBool and connectsdk.NewString to overcome this issue.
	// These helper functions are provided by the SDK's root package.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.

	var body domain.ApproveTokenRequest
	body.MandateSignatureDate = connectsdk.NewString("20150201")
	body.MandateSignaturePlace = connectsdk.NewString("Monument Valley")
	body.MandateSigned = connectsdk.NewBool(true)

	err := client.V1().Merchant("merchantId").Tokens().Approvesepadirectdebit("tokenId", body, nil)

	fmt.Println(err)
}
