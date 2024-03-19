// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func getDeviceFingerprintExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer func(client *connectsdk.Client) {
		_ = client.Close()
	}(client)

	var body domain.DeviceFingerprintRequest

	response, err := client.V1().Merchant("merchantId").Products().DeviceFingerprint(1, body, nil)

	fmt.Println(response, err)
}
