package products

import (
	"testing"

	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
)

func TestDirectoryParamsToRequestParameters(t *testing.T) {
	lParams := &DirectoryParams{}
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
}
