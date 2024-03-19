// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors"
)

func createRefundCaptureExample() {
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
	amountOfMoney.Amount = connectsdk.NewInt64(500)
	amountOfMoney.CurrencyCode = connectsdk.NewString("EUR")

	var refundReferences domain.RefundReferences
	refundReferences.MerchantReference = connectsdk.NewString("AcmeOrder0001")

	var body domain.RefundRequest
	body.AmountOfMoney = &amountOfMoney
	body.RefundReferences = &refundReferences

	response, err := client.V1().Merchant("merchantId").Captures().Refund("captureId", body, nil)
	switch realError := err.(type) {
	case *v1Errors.DeclinedRefundError:
		{
			handleDeclinedRefund(realError.RefundResult())

			break
		}
	case v1Errors.APIError:
		{
			handleAPIErrors(realError.Errors())

			break
		}
	}

	fmt.Println(response, err)
}
