package products

import (
	"runtime"
	"testing"

	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
)

func TestFindParamsToRequestParameters(t *testing.T) {
	lParams := &FindParams{}
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

func paramRequestCmp(t *testing.T, a communication.ParamRequest, b communication.RequestParams) {
	params := a.ToRequestParameters()

	if !requestParamsCmp(params, b) {
		buf := make([]byte, 1<<16)
		runtime.Stack(buf, true)
		t.Fatal("paramRequestCmp failed on equality", params, b, string(buf))
	}
}

func requestParamsCmp(a, b communication.RequestParams) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
