// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func createCaptureDisputeExample() {
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

	var amountOfMoney domain.AmountOfMoney
	amountOfMoney.Amount = connectsdk.NewInt64(1234)
	amountOfMoney.CurrencyCode = connectsdk.NewString("USD")

	var body domain.CreateDisputeRequest
	body.AmountOfMoney = &amountOfMoney
	body.ContactPerson = connectsdk.NewString("Wile Coyote")
	body.EmailAddress = connectsdk.NewString("wile.e.coyote@acmelabs.com")
	body.ReplyTo = connectsdk.NewString("r.runner@acmelabs.com")
	body.RequestMessage = connectsdk.NewString("This is the message from the merchant to GlobalCollect. It is a a freeform text field.")

	response, err := client.V1().Merchant("merchantId").Captures().Dispute("captureId", body, nil)

	fmt.Println(response, err)
}
