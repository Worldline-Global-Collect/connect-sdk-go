// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
)

func getHostedMandateManagementExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer func(client *connectsdk.Client) {
		_ = client.Close()
	}(client)

	response, err := client.V1().Merchant("merchantId").Hostedmandatemanagements().Get("hostedMandateManagementId", nil)

	fmt.Println(response, err)
}
