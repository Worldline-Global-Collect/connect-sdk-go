// This file was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package examples

import (
	"fmt"

	"github.com/Worldline-Global-Collect/connect-sdk-go"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
)

func riskAssessmentCardsExample() {
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
	card.Cvv = connectsdk.NewString("123")
	card.ExpiryDate = connectsdk.NewString("1299")

	var flightLegs []domain.AirlineFlightLeg

	var flightLeg1 domain.AirlineFlightLeg
	flightLeg1.AirlineClass = connectsdk.NewString("1")
	flightLeg1.ArrivalAirport = connectsdk.NewString("AMS")
	flightLeg1.CarrierCode = connectsdk.NewString("KL")
	flightLeg1.Date = connectsdk.NewString("20150102")
	flightLeg1.DepartureTime = connectsdk.NewString("17:59")
	flightLeg1.Fare = connectsdk.NewString("fare")
	flightLeg1.FareBasis = connectsdk.NewString("INTERNET")
	flightLeg1.FlightNumber = connectsdk.NewString("791")
	flightLeg1.Number = connectsdk.NewInt32(1)
	flightLeg1.OriginAirport = connectsdk.NewString("BCN")
	flightLeg1.StopoverCode = connectsdk.NewString("non-permitted")

	flightLegs = append(flightLegs, flightLeg1)

	var flightLeg2 domain.AirlineFlightLeg
	flightLeg2.AirlineClass = connectsdk.NewString("1")
	flightLeg2.ArrivalAirport = connectsdk.NewString("BCN")
	flightLeg2.CarrierCode = connectsdk.NewString("KL")
	flightLeg2.Date = connectsdk.NewString("20150102")
	flightLeg2.DepartureTime = connectsdk.NewString("23:59")
	flightLeg2.Fare = connectsdk.NewString("fare")
	flightLeg2.FareBasis = connectsdk.NewString("INTERNET")
	flightLeg2.FlightNumber = connectsdk.NewString("792")
	flightLeg2.Number = connectsdk.NewInt32(2)
	flightLeg2.OriginAirport = connectsdk.NewString("AMS")
	flightLeg2.StopoverCode = connectsdk.NewString("non-permitted")

	flightLegs = append(flightLegs, flightLeg2)

	var airlineData domain.AirlineData
	airlineData.AgentNumericCode = connectsdk.NewString("123321")
	airlineData.Code = connectsdk.NewString("123")
	airlineData.FlightDate = connectsdk.NewString("20150102")
	airlineData.FlightLegs = &flightLegs
	airlineData.InvoiceNumber = connectsdk.NewString("123456")
	airlineData.IsETicket = connectsdk.NewBool(true)
	airlineData.IsRestrictedTicket = connectsdk.NewBool(true)
	airlineData.IsThirdParty = connectsdk.NewBool(true)
	airlineData.IssueDate = connectsdk.NewString("20150101")
	airlineData.MerchantCustomerID = connectsdk.NewString("14")
	airlineData.Name = connectsdk.NewString("Air France KLM")
	airlineData.PassengerName = connectsdk.NewString("WECOYOTE")
	airlineData.PlaceOfIssue = connectsdk.NewString("Utah")
	airlineData.PNR = connectsdk.NewString("4JTGKT")
	airlineData.PointOfSale = connectsdk.NewString("IATA point of sale name")
	airlineData.PosCityCode = connectsdk.NewString("AMS")
	airlineData.TicketDeliveryMethod = connectsdk.NewString("e-ticket")
	airlineData.TicketNumber = connectsdk.NewString("KLM20050000")

	var additionalInput domain.AdditionalOrderInputAirlineData
	additionalInput.AirlineData = &airlineData

	var amountOfMoney domain.AmountOfMoney
	amountOfMoney.Amount = connectsdk.NewInt64(100)
	amountOfMoney.CurrencyCode = connectsdk.NewString("EUR")

	var billingAddress domain.Address
	billingAddress.CountryCode = connectsdk.NewString("US")

	var customer domain.CustomerRiskAssessment
	customer.AccountType = connectsdk.NewString("existing")
	customer.BillingAddress = &billingAddress
	customer.Locale = connectsdk.NewString("en_US")

	var order domain.OrderRiskAssessment
	order.AdditionalInput = &additionalInput
	order.AmountOfMoney = &amountOfMoney
	order.Customer = &customer

	var body domain.RiskAssessmentCard
	body.Card = &card
	body.Order = &order

	response, err := client.V1().Merchant("merchantId").Riskassessments().Cards(body, nil)

	fmt.Println(response, err)
}
