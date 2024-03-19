// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func createHostedMandateManagementExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer func(client *connectsdk.Client) {
		_ = client.Close()
	}(client)

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper function connectsdk.NewString to overcome this issue.
	// This helper function is provided by the SDK's root package.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.

	var createMandateInfo domain.HostedMandateInfo
	createMandateInfo.CustomerReference = connectsdk.NewString("idonthaveareference")
	createMandateInfo.RecurrenceType = connectsdk.NewString("RECURRING")
	createMandateInfo.SignatureType = connectsdk.NewString("UNSIGNED")

	var hostedMandateManagementSpecificInput domain.HostedMandateManagementSpecificInput
	hostedMandateManagementSpecificInput.Locale = connectsdk.NewString("fr_FR")
	hostedMandateManagementSpecificInput.ReturnURL = connectsdk.NewString("http://www.example.com")
	hostedMandateManagementSpecificInput.Variant = connectsdk.NewString("101")

	var body domain.CreateHostedMandateManagementRequest
	body.CreateMandateInfo = &createMandateInfo
	body.HostedMandateManagementSpecificInput = &hostedMandateManagementSpecificInput

	response, err := client.V1().Merchant("merchantId").Hostedmandatemanagements().Create(body, nil)

	fmt.Println(response, err)
}
