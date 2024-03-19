// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
)

func getPaymentDeviceFingerprintDetailsExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer func(client *connectsdk.Client) {
		_ = client.Close()
	}(client)

	response, err := client.V1().Merchant("merchantId").Payments().Devicefingerprint("paymentId", nil)

	fmt.Println(response, err)
}
