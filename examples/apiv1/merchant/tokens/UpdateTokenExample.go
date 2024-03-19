// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func updateTokenExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer func(client *connectsdk.Client) {
		_ = client.Close()
	}(client)

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions connectsdk.NewInt32 and connectsdk.NewString to overcome this issue.
	// These helper functions are provided by the SDK's root package.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.

	var billingAddress domain.Address
	billingAddress.AdditionalInfo = connectsdk.NewString("b")
	billingAddress.City = connectsdk.NewString("Monument Valley")
	billingAddress.CountryCode = connectsdk.NewString("US")
	billingAddress.HouseNumber = connectsdk.NewString("13")
	billingAddress.State = connectsdk.NewString("Utah")
	billingAddress.Street = connectsdk.NewString("Desertroad")
	billingAddress.Zip = connectsdk.NewString("84536")

	var companyInformation domain.CompanyInformation
	companyInformation.Name = connectsdk.NewString("Acme Labs")

	var name domain.PersonalNameToken
	name.FirstName = connectsdk.NewString("Wile")
	name.Surname = connectsdk.NewString("Coyote")
	name.SurnamePrefix = connectsdk.NewString("E.")

	var personalInformation domain.PersonalInformationToken
	personalInformation.Name = &name

	var customer domain.CustomerToken
	customer.BillingAddress = &billingAddress
	customer.CompanyInformation = &companyInformation
	customer.MerchantCustomerID = connectsdk.NewString("1234")
	customer.PersonalInformation = &personalInformation

	var cardWithoutCvv domain.CardWithoutCvv
	cardWithoutCvv.CardNumber = connectsdk.NewString("4567350000427977")
	cardWithoutCvv.CardholderName = connectsdk.NewString("Wile E. Coyote")
	cardWithoutCvv.ExpiryDate = connectsdk.NewString("1299")
	cardWithoutCvv.IssueNumber = connectsdk.NewString("12")

	var data domain.TokenCardData
	data.CardWithoutCvv = &cardWithoutCvv

	var card domain.TokenCard
	card.Customer = &customer
	card.Data = &data

	var body domain.UpdateTokenRequest
	body.Card = &card
	body.PaymentProductID = connectsdk.NewInt32(1)

	err := client.V1().Merchant("merchantId").Tokens().Update("tokenId", body, nil)

	fmt.Println(err)
}
