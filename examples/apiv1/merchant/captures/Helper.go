// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func getClient() (*connectsdk.Client, error) {
	apiKeyID := "someKey"
	secretAPIKey := "someSecret"

	return connectsdk.CreateClient(apiKeyID, secretAPIKey, "Worldline")
}

func handleDeclinedRefund(refundResult *domain.RefundResult) {
	// handle the result here
}

func handleAPIErrors(errors []domain.APIError) {
	// handle the errors here
}
