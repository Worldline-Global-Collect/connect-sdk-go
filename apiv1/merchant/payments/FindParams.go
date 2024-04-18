// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package payments

import (
	"strconv"

	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
)

// FindParams represents query parameters for Find payments
//
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/payments/find.html
type FindParams struct {
	HostedCheckoutID  *string
	MerchantReference *string
	MerchantOrderID   *int64
	Offset            *int32
	Limit             *int32
}

// ToRequestParameters converts the query to communication.RequestParams
func (params *FindParams) ToRequestParameters() communication.RequestParams {
	reqParams := communication.RequestParams{}

	if params.HostedCheckoutID != nil {
		param, _ := communication.NewRequestParam("hostedCheckoutId", *params.HostedCheckoutID)
		reqParams = append(reqParams, *param)
	}
	if params.MerchantReference != nil {
		param, _ := communication.NewRequestParam("merchantReference", *params.MerchantReference)
		reqParams = append(reqParams, *param)
	}
	if params.MerchantOrderID != nil {
		param, _ := communication.NewRequestParam("merchantOrderId", strconv.FormatInt(*params.MerchantOrderID, 10))
		reqParams = append(reqParams, *param)
	}
	if params.Offset != nil {
		param, _ := communication.NewRequestParam("offset", strconv.FormatInt(int64(*params.Offset), 10))
		reqParams = append(reqParams, *param)
	}
	if params.Limit != nil {
		param, _ := communication.NewRequestParam("limit", strconv.FormatInt(int64(*params.Limit), 10))
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewFindParams constructs a new FindParams instance
func NewFindParams() *FindParams {
	return &FindParams{}
}
