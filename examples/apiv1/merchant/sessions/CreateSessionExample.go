// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func createSessionExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer func(client *connectsdk.Client) {
		_ = client.Close()
	}(client)

	var tokens []string
	tokens = append(tokens, "126166b16ed04b3ab85fb06da1d7a167")
	tokens = append(tokens, "226166b16ed04b3ab85fb06da1d7a167")
	tokens = append(tokens, "122c5b4d-dd40-49f0-b7c9-3594212167a9")
	tokens = append(tokens, "326166b16ed04b3ab85fb06da1d7a167")
	tokens = append(tokens, "426166b16ed04b3ab85fb06da1d7a167")

	var body domain.SessionRequest
	body.Tokens = &tokens

	response, err := client.V1().Merchant("merchantId").Sessions().Create(body, nil)

	fmt.Println(response, err)
}
