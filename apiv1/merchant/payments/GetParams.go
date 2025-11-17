// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package payments

import (
	"strconv"

	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
)

// GetParams represents query parameters for Get payment
//
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/payments/get.html
type GetParams struct {
	ReturnOperations *bool
}

// ToRequestParameters converts the query to communication.RequestParams
func (params *GetParams) ToRequestParameters() communication.RequestParams {
	reqParams := communication.RequestParams{}

	if params.ReturnOperations != nil {
		param, _ := communication.NewRequestParam("returnOperations", strconv.FormatBool(*params.ReturnOperations))
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewGetParams constructs a new GetParams instance
func NewGetParams() *GetParams {
	return &GetParams{}
}
