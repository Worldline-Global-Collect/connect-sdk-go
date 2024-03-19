// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func completePaymentExample() {
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

	var card domain.CardWithoutCvv
	card.CardNumber = connectsdk.NewString("67030000000000003")
	card.CardholderName = connectsdk.NewString("Wile E. Coyote")
	card.ExpiryDate = connectsdk.NewString("1299")

	var cardPaymentMethodSpecificInput domain.CompletePaymentCardPaymentMethodSpecificInput
	cardPaymentMethodSpecificInput.Card = &card

	var body domain.CompletePaymentRequest
	body.CardPaymentMethodSpecificInput = &cardPaymentMethodSpecificInput

	response, err := client.V1().Merchant("merchantId").Payments().Complete("paymentId", body, nil)

	fmt.Println(response, err)
}
