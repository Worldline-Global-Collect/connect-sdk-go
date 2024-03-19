// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors"
)

func createPayoutExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer func(client *connectsdk.Client) {
		_ = client.Close()
	}(client)

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions connectsdk.NewInt64 and connectsdk.NewString to overcome this issue.
	// These helper functions are provided by the SDK's root package.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.

	var bankAccountIban domain.BankAccountIban
	bankAccountIban.AccountHolderName = connectsdk.NewString("Wile E. Coyote")
	bankAccountIban.Iban = connectsdk.NewString("IT60X0542811101000000123456")

	var bankTransferPayoutMethodSpecificInput domain.BankTransferPayoutMethodSpecificInput
	bankTransferPayoutMethodSpecificInput.BankAccountIban = &bankAccountIban
	bankTransferPayoutMethodSpecificInput.PayoutDate = connectsdk.NewString("20150102")
	bankTransferPayoutMethodSpecificInput.PayoutText = connectsdk.NewString("Payout Acme")
	bankTransferPayoutMethodSpecificInput.SwiftCode = connectsdk.NewString("swift")

	var amountOfMoney domain.AmountOfMoney
	amountOfMoney.Amount = connectsdk.NewInt64(2345)
	amountOfMoney.CurrencyCode = connectsdk.NewString("EUR")

	var address domain.Address
	address.City = connectsdk.NewString("Burbank")
	address.CountryCode = connectsdk.NewString("US")
	address.HouseNumber = connectsdk.NewString("411")
	address.State = connectsdk.NewString("California")
	address.Street = connectsdk.NewString("N Hollywood Way")
	address.Zip = connectsdk.NewString("91505")

	var companyInformation domain.CompanyInformation
	companyInformation.Name = connectsdk.NewString("Acme Labs")

	var contactDetails domain.ContactDetailsBase
	contactDetails.EmailAddress = connectsdk.NewString("wile.e.coyote@acmelabs.com")

	var name domain.PersonalName
	name.FirstName = connectsdk.NewString("Wile")
	name.Surname = connectsdk.NewString("Coyote")
	name.SurnamePrefix = connectsdk.NewString("E.")
	name.Title = connectsdk.NewString("Mr.")

	var customer domain.PayoutCustomer
	customer.Address = &address
	customer.CompanyInformation = &companyInformation
	customer.ContactDetails = &contactDetails
	customer.Name = &name

	var references domain.PayoutReferences
	references.MerchantReference = connectsdk.NewString("AcmeOrder001")

	var payoutDetails domain.PayoutDetails
	payoutDetails.AmountOfMoney = &amountOfMoney
	payoutDetails.Customer = &customer
	payoutDetails.References = &references

	var body domain.CreatePayoutRequest
	body.BankTransferPayoutMethodSpecificInput = &bankTransferPayoutMethodSpecificInput
	body.PayoutDetails = &payoutDetails

	response, err := client.V1().Merchant("merchantId").Payouts().Create(body, nil)
	switch realError := err.(type) {
	case *v1Errors.DeclinedPayoutError:
		{
			handleDeclinedPayout(realError.PayoutResult())

			break
		}
	case v1Errors.APIError:
		{
			handleAPIErrors(realError.Errors())

			break
		}
	}

	fmt.Println(response, err)
}
