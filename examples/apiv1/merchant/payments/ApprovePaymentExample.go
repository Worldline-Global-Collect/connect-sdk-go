// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func approvePaymentExample() {
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

	var directDebitPaymentMethodSpecificInput domain.ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput
	directDebitPaymentMethodSpecificInput.DateCollect = connectsdk.NewString("20150201")
	directDebitPaymentMethodSpecificInput.Token = connectsdk.NewString("bfa8a7e4-4530-455a-858d-204ba2afb77e")

	var references domain.OrderReferencesApprovePayment
	references.MerchantReference = connectsdk.NewString("AcmeOrder0001")

	var order domain.OrderApprovePayment
	order.References = &references

	var body domain.ApprovePaymentRequest
	body.Amount = connectsdk.NewInt64(2980)
	body.DirectDebitPaymentMethodSpecificInput = &directDebitPaymentMethodSpecificInput
	body.Order = &order

	response, err := client.V1().Merchant("merchantId").Payments().Approve("paymentId", body, nil)

	fmt.Println(response, err)
}
