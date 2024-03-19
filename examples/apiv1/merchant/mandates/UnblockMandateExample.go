// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
)

func unblockMandateExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer func(client *connectsdk.Client) {
		_ = client.Close()
	}(client)

	response, err := client.V1().Merchant("merchantId").Mandates().Unblock("42268d8067df43e18a50a2ebf4bdb729", nil)

	fmt.Println(response, err)
}
