// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"github.com/Worldline-Global-Collect/connect-sdk-go"
)

func getClient() (*connectsdk.Client, error) {
	apiKeyID := "someKey"
	secretAPIKey := "someSecret"

	return connectsdk.CreateClient(apiKeyID, secretAPIKey, "Worldline")
}
