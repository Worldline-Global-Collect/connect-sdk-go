// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"
	"os"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/disputes"
	"github.com/Worldline-Global-Collect/connect-sdk-go/domain"
)

func uploadDisputeFileExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer func(client *connectsdk.Client) {
		_ = client.Close()
	}(client)

	var body disputes.UploadFileRequest
	fileContent, _ := os.Open("file.pdf")
	defer fileContent.Close()
	body.File, _ = domain.NewUploadableFile("file.pdf", fileContent, "application/pdf")

	response, err := client.V1().Merchant("merchantId").Disputes().UploadFile("disputeId", body, nil)

	fmt.Println(response, err)
}
