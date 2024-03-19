package tokens

import (
	"runtime"
	"testing"

	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
)

func TestDeleteParamsToRequestParameters(t *testing.T) {
	lParams := &DeleteParams{}
	paramList := communication.RequestParams{}

	paramRequestCmp(t, lParams, paramList)

	{
		lParams.MandateCancelDate = new(string)
		*lParams.MandateCancelDate = "20160610"

		param, _ := communication.NewRequestParam("mandateCancelDate", "20160610")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.MandateCancelDate = new(string)
		*lParams.MandateCancelDate = ""

		param, _ := communication.NewRequestParam("mandateCancelDate", "")
		paramList = communication.RequestParams{}
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.MandateCancelDate = nil

		paramList = communication.RequestParams{}
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
