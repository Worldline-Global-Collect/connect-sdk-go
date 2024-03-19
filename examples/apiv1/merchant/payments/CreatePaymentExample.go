// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors"
)

func createPaymentExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer func(client *connectsdk.Client) {
		_ = client.Close()
	}(client)

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions connectsdk.NewBool, connectsdk.NewInt32, connectsdk.NewInt64 and connectsdk.NewString to overcome this issue.
	// These helper functions are provided by the SDK's root package.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.

	var card domain.Card
	card.CardNumber = connectsdk.NewString("4567350000427977")
	card.CardholderName = connectsdk.NewString("Wile E. Coyote")
	card.Cvv = connectsdk.NewString("123")
	card.ExpiryDate = connectsdk.NewString("1299")

	var authenticationAmount domain.AmountOfMoney
	authenticationAmount.Amount = connectsdk.NewInt64(2980)
	authenticationAmount.CurrencyCode = connectsdk.NewString("EUR")

	var redirectionData domain.RedirectionData
	redirectionData.ReturnURL = connectsdk.NewString("https://hostname.myownwebsite.url")

	var threeDSecure domain.ThreeDSecure
	threeDSecure.AuthenticationAmount = &authenticationAmount
	threeDSecure.AuthenticationFlow = connectsdk.NewString("browser")
	threeDSecure.ChallengeCanvasSize = connectsdk.NewString("600x400")
	threeDSecure.ChallengeIndicator = connectsdk.NewString("challenge-requested")
	threeDSecure.ExemptionRequest = connectsdk.NewString("none")
	threeDSecure.RedirectionData = &redirectionData
	threeDSecure.SkipAuthentication = connectsdk.NewBool(false)

	var cardPaymentMethodSpecificInput domain.CardPaymentMethodSpecificInput
	cardPaymentMethodSpecificInput.Card = &card
	cardPaymentMethodSpecificInput.IsRecurring = connectsdk.NewBool(false)
	cardPaymentMethodSpecificInput.MerchantInitiatedReasonIndicator = connectsdk.NewString("delayedCharges")
	cardPaymentMethodSpecificInput.PaymentProductID = connectsdk.NewInt32(1)
	cardPaymentMethodSpecificInput.ThreeDSecure = &threeDSecure
	cardPaymentMethodSpecificInput.TransactionChannel = connectsdk.NewString("ECOMMERCE")

	var amountOfMoney domain.AmountOfMoney
	amountOfMoney.Amount = connectsdk.NewInt64(2980)
	amountOfMoney.CurrencyCode = connectsdk.NewString("EUR")

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
	companyInformation.VatNumber = connectsdk.NewString("1234AB5678CD")

	var contactDetails domain.ContactDetails
	contactDetails.EmailAddress = connectsdk.NewString("wile.e.coyote@acmelabs.com")
	contactDetails.FaxNumber = connectsdk.NewString("+1234567891")
	contactDetails.PhoneNumber = connectsdk.NewString("+1234567890")

	var browserData domain.BrowserData
	browserData.ColorDepth = connectsdk.NewInt32(24)
	browserData.JavaEnabled = connectsdk.NewBool(false)
	browserData.ScreenHeight = connectsdk.NewString("1200")
	browserData.ScreenWidth = connectsdk.NewString("1920")

	var device domain.CustomerDevice
	device.AcceptHeader = connectsdk.NewString("texthtml,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	device.BrowserData = &browserData
	device.IPAddress = connectsdk.NewString("123.123.123.123")
	device.Locale = connectsdk.NewString("en-US")
	device.TimezoneOffsetUtcMinutes = connectsdk.NewString("420")
	device.UserAgent = connectsdk.NewString("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1 Safari/605.1.15")

	var name domain.PersonalName
	name.FirstName = connectsdk.NewString("Wile")
	name.Surname = connectsdk.NewString("Coyote")
	name.SurnamePrefix = connectsdk.NewString("E.")
	name.Title = connectsdk.NewString("Mr.")

	var personalInformation domain.PersonalInformation
	personalInformation.DateOfBirth = connectsdk.NewString("19490917")
	personalInformation.Gender = connectsdk.NewString("male")
	personalInformation.Name = &name

	var customer domain.Customer
	customer.AccountType = connectsdk.NewString("none")
	customer.BillingAddress = &billingAddress
	customer.CompanyInformation = &companyInformation
	customer.ContactDetails = &contactDetails
	customer.Device = &device
	customer.Locale = connectsdk.NewString("en_US")
	customer.MerchantCustomerID = connectsdk.NewString("1234")
	customer.PersonalInformation = &personalInformation

	var invoiceData domain.OrderInvoiceData
	invoiceData.InvoiceDate = connectsdk.NewString("20140306191500")
	invoiceData.InvoiceNumber = connectsdk.NewString("000000123")

	var references domain.OrderReferences
	references.Descriptor = connectsdk.NewString("Fast and Furry-ous")
	references.InvoiceData = &invoiceData
	references.MerchantOrderID = connectsdk.NewInt64(123456)
	references.MerchantReference = connectsdk.NewString("AcmeOrder0001")

	var shippingName domain.PersonalName
	shippingName.FirstName = connectsdk.NewString("Road")
	shippingName.Surname = connectsdk.NewString("Runner")
	shippingName.Title = connectsdk.NewString("Miss")

	var address domain.AddressPersonal
	address.AdditionalInfo = connectsdk.NewString("Suite II")
	address.City = connectsdk.NewString("Monument Valley")
	address.CountryCode = connectsdk.NewString("US")
	address.HouseNumber = connectsdk.NewString("1")
	address.Name = &shippingName
	address.State = connectsdk.NewString("Utah")
	address.Street = connectsdk.NewString("Desertroad")
	address.Zip = connectsdk.NewString("84536")

	var shipping domain.Shipping
	shipping.Address = &address

	var items []domain.LineItem

	var item1AmountOfMoney domain.AmountOfMoney
	item1AmountOfMoney.Amount = connectsdk.NewInt64(2500)
	item1AmountOfMoney.CurrencyCode = connectsdk.NewString("EUR")

	var item1InvoiceData domain.LineItemInvoiceData
	item1InvoiceData.Description = connectsdk.NewString("ACME Super Outfit")
	item1InvoiceData.NrOfItems = connectsdk.NewString("1")
	item1InvoiceData.PricePerItem = connectsdk.NewInt64(2500)

	var item1 domain.LineItem
	item1.AmountOfMoney = &item1AmountOfMoney
	item1.InvoiceData = &item1InvoiceData

	items = append(items, item1)

	var item2AmountOfMoney domain.AmountOfMoney
	item2AmountOfMoney.Amount = connectsdk.NewInt64(480)
	item2AmountOfMoney.CurrencyCode = connectsdk.NewString("EUR")

	var item2InvoiceData domain.LineItemInvoiceData
	item2InvoiceData.Description = connectsdk.NewString("Aspirin")
	item2InvoiceData.NrOfItems = connectsdk.NewString("12")
	item2InvoiceData.PricePerItem = connectsdk.NewInt64(40)

	var item2 domain.LineItem
	item2.AmountOfMoney = &item2AmountOfMoney
	item2.InvoiceData = &item2InvoiceData

	items = append(items, item2)

	var shoppingCart domain.ShoppingCart
	shoppingCart.Items = &items

	var order domain.Order
	order.AmountOfMoney = &amountOfMoney
	order.Customer = &customer
	order.References = &references
	order.Shipping = &shipping
	order.ShoppingCart = &shoppingCart

	var body domain.CreatePaymentRequest
	body.CardPaymentMethodSpecificInput = &cardPaymentMethodSpecificInput
	body.Order = &order

	response, err := client.V1().Merchant("merchantId").Payments().Create(body, nil)
	switch realError := err.(type) {
	case *v1Errors.DeclinedPaymentError:
		{
			handleDeclinedPayment(realError.PaymentResult())

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
