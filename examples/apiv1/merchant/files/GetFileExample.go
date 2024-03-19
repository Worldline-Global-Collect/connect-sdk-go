// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"
	"io"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
)

func getFileExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer func(client *connectsdk.Client) {
		_ = client.Close()
	}(client)

	err := client.V1().Merchant("merchantId").Files().GetFile("fileId", nil, func(headers []communication.Header, bodyReader io.Reader) error {
		// use the headers and body reader, and return any error that occurred
		return nil
	})

	fmt.Println(err)
}
