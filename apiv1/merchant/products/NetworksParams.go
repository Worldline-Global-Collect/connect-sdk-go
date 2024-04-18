// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package products

import (
	"strconv"

	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
)

// NetworksParams represents query parameters for Get payment product networks
//
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/products/networks.html
type NetworksParams struct {
	CountryCode  *string
	CurrencyCode *string
	Amount       *int64
	IsRecurring  *bool
}

// ToRequestParameters converts the query to communication.RequestParams
func (params *NetworksParams) ToRequestParameters() communication.RequestParams {
	reqParams := communication.RequestParams{}

	if params.CountryCode != nil {
		param, _ := communication.NewRequestParam("countryCode", *params.CountryCode)
		reqParams = append(reqParams, *param)
	}
	if params.CurrencyCode != nil {
		param, _ := communication.NewRequestParam("currencyCode", *params.CurrencyCode)
		reqParams = append(reqParams, *param)
	}
	if params.Amount != nil {
		param, _ := communication.NewRequestParam("amount", strconv.FormatInt(*params.Amount, 10))
		reqParams = append(reqParams, *param)
	}
	if params.IsRecurring != nil {
		param, _ := communication.NewRequestParam("isRecurring", strconv.FormatBool(*params.IsRecurring))
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewNetworksParams constructs a new NetworksParams instance
func NewNetworksParams() *NetworksParams {
	return &NetworksParams{}
}
