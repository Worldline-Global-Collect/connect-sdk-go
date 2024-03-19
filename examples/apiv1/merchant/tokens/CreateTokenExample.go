// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func createTokenExample() {
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
	billingAddress.AdditionalInfo = connectsdk.NewString("Suite II")
	billingAddress.City = connectsdk.NewString("Monument Valley")
	billingAddress.CountryCode = connectsdk.NewString("US")
	billingAddress.HouseNumber = connectsdk.NewString("1")
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

	var bankAccountBban domain.BankAccountBban
	bankAccountBban.AccountNumber = connectsdk.NewString("000000123456")
	bankAccountBban.BankCode = connectsdk.NewString("05428")
	bankAccountBban.BranchCode = connectsdk.NewString("11101")
	bankAccountBban.CheckDigit = connectsdk.NewString("X")
	bankAccountBban.CountryCode = connectsdk.NewString("IT")

	var paymentProduct705SpecificData domain.TokenNonSepaDirectDebitPaymentProduct705SpecificData
	paymentProduct705SpecificData.AuthorisationID = connectsdk.NewString("123456")
	paymentProduct705SpecificData.BankAccountBban = &bankAccountBban

	var mandate domain.MandateNonSepaDirectDebit
	mandate.PaymentProduct705SpecificData = &paymentProduct705SpecificData

	var nonSepaDirectDebit domain.TokenNonSepaDirectDebit
	nonSepaDirectDebit.Customer = &customer
	nonSepaDirectDebit.Mandate = &mandate

	var body domain.CreateTokenRequest
	body.NonSepaDirectDebit = &nonSepaDirectDebit
	body.PaymentProductID = connectsdk.NewInt32(705)

	response, err := client.V1().Merchant("merchantId").Tokens().Create(body, nil)

	fmt.Println(response, err)
}
