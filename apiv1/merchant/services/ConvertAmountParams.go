// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package services

import (
	"strconv"

	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
)

// ConvertAmountParams represents query parameters for Convert amount
//
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/services/convertAmount.html
type ConvertAmountParams struct {
	Source *string
	Target *string
	Amount *int64
}

// ToRequestParameters converts the query to communication.RequestParams
func (params *ConvertAmountParams) ToRequestParameters() communication.RequestParams {
	reqParams := communication.RequestParams{}

	if params.Source != nil {
		param, _ := communication.NewRequestParam("source", *params.Source)
		reqParams = append(reqParams, *param)
	}
	if params.Target != nil {
		param, _ := communication.NewRequestParam("target", *params.Target)
		reqParams = append(reqParams, *param)
	}
	if params.Amount != nil {
		param, _ := communication.NewRequestParam("amount", strconv.FormatInt(*params.Amount, 10))
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewConvertAmountParams constructs a new ConvertAmountParams instance
func NewConvertAmountParams() *ConvertAmountParams {
	return &ConvertAmountParams{}
}
