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

func handleDeclinedPayment(paymentResult *domain.CreatePaymentResult) {
	// handle the result here
}

func handleDeclinedRefund(refundResult *domain.RefundResult) {
	// handle the result here
}

func handleErrorResponse(errorID string, errors []domain.APIError) {
	// handle the error response here
}
