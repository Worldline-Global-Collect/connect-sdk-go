package products

import (
	"testing"

	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
)

func TestGetParamsToRequestParameters(t *testing.T) {
	lParams := &GetParams{}
	paramList := communication.RequestParams{}

	paramRequestCmp(t, lParams, paramList)

	{
		lParams.CountryCode = new(string)
		*lParams.CountryCode = "NL"

		param, _ := communication.NewRequestParam("countryCode", "NL")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.CurrencyCode = new(string)
		*lParams.CurrencyCode = "EUR"

		param, _ := communication.NewRequestParam("currencyCode", "EUR")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.Locale = new(string)
		*lParams.Locale = "nl_NL"

		param, _ := communication.NewRequestParam("locale", "nl_NL")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		amount := new(int64)
		*amount = 1000
		lParams.Amount = amount

		param, _ := communication.NewRequestParam("amount", "1000")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		isRecurring := new(bool)
		*isRecurring = true
		lParams.IsRecurring = isRecurring

		param, _ := communication.NewRequestParam("isRecurring", "true")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.AddHide("fields")

		param, _ := communication.NewRequestParam("hide", "fields")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.AddHide("accountsOnFile")

		param, _ := communication.NewRequestParam("hide", "accountsOnFile")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.Amount = nil

		paramList = append(paramList[0:3], paramList[4:]...)
	}
	paramRequestCmp(t, lParams, paramList)
}
