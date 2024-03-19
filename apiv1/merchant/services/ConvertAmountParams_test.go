package services

import (
	communication2 "github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
	"runtime"
	"testing"
)

func TestToRequestParameters(t *testing.T) {
	lParams := &ConvertAmountParams{}
	paramList := communication2.RequestParams{}

	paramRequestCmp(t, lParams, paramList)

	{
		lParams.Source = new(string)
		*lParams.Source = "EUR"

		param, _ := communication2.NewRequestParam("source", "EUR")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.Target = new(string)
		*lParams.Target = "USD"

		param, _ := communication2.NewRequestParam("target", "USD")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.Amount = new(int64)
		*lParams.Amount = 1000

		param, _ := communication2.NewRequestParam("amount", "1000")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.Source = nil

		paramList = paramList[1:]
	}
	paramRequestCmp(t, lParams, paramList)
}

func paramRequestCmp(t *testing.T, a communication2.ParamRequest, b communication2.RequestParams) {
	params := a.ToRequestParameters()

	if !requestParamsCmp(params, b) {
		buf := make([]byte, 1<<16)
		runtime.Stack(buf, true)
		t.Fatal("paramRequestCmp failed on equality", params, b, string(buf))
	}
}

func requestParamsCmp(a, b communication2.RequestParams) bool {
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
