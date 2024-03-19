// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func createPaymentProductSessionExample() {
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

	var paymentProductSession302SpecificInput domain.MobilePaymentProductSession302SpecificInput
	paymentProductSession302SpecificInput.DisplayName = connectsdk.NewString("Worldline")
	paymentProductSession302SpecificInput.DomainName = connectsdk.NewString("pay1.checkout.worldline-solutions.com")
	paymentProductSession302SpecificInput.ValidationURL = connectsdk.NewString("<VALIDATION URL RECEIVED FROM APPLE>")

	var body domain.CreatePaymentProductSessionRequest
	body.PaymentProductSession302SpecificInput = &paymentProductSession302SpecificInput

	response, err := client.V1().Merchant("merchantId").Products().Sessions(302, body, nil)

	fmt.Println(response, err)
}
