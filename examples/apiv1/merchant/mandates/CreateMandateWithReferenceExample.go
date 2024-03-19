// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func createMandateWithReferenceExample() {
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

	var bankAccountIban domain.BankAccountIban
	bankAccountIban.Iban = connectsdk.NewString("DE46720200700359736690")

	var contactDetails domain.MandateContactDetails
	contactDetails.EmailAddress = connectsdk.NewString("wile.e.coyote@acmelabs.com")

	var mandateAddress domain.MandateAddress
	mandateAddress.City = connectsdk.NewString("Monumentenvallei")
	mandateAddress.CountryCode = connectsdk.NewString("NL")
	mandateAddress.Street = connectsdk.NewString("Woestijnweg")
	mandateAddress.Zip = connectsdk.NewString("1337XD")

	var name domain.MandatePersonalName
	name.FirstName = connectsdk.NewString("Wile")
	name.Surname = connectsdk.NewString("Coyote")

	var personalInformation domain.MandatePersonalInformation
	personalInformation.Name = &name
	personalInformation.Title = connectsdk.NewString("Miss")

	var customer domain.MandateCustomer
	customer.BankAccountIban = &bankAccountIban
	customer.CompanyName = connectsdk.NewString("Acme labs")
	customer.ContactDetails = &contactDetails
	customer.MandateAddress = &mandateAddress
	customer.PersonalInformation = &personalInformation

	var body domain.CreateMandateRequest
	body.Customer = &customer
	body.CustomerReference = connectsdk.NewString("idonthaveareference")
	body.Language = connectsdk.NewString("nl")
	body.RecurrenceType = connectsdk.NewString("UNIQUE")
	body.SignatureType = connectsdk.NewString("UNSIGNED")

	response, err := client.V1().Merchant("merchantId").Mandates().CreateWithMandateReference("42268d8067df43e18a50a2ebf4bdb729", body, nil)

	fmt.Println(response, err)
}
