// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package productgroups

import (
	"strconv"

	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
)

// GetParams represents query parameters for Get payment product group
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/productgroups/get.html
type GetParams struct {
	CountryCode    *string
	CurrencyCode   *string
	Locale         *string
	Amount         *int64
	IsRecurring    *bool
	IsInstallments *bool
	Hide           []string
}

// AddHide adds an element to the Hide array.
func (params *GetParams) AddHide(value string) {
	params.Hide = append(params.Hide, value)
	return
}

// ToRequestParameters converts the query to communication.RequestParams
func (params *GetParams) ToRequestParameters() communication.RequestParams {
	reqParams := communication.RequestParams{}

	if params.CountryCode != nil {
		param, _ := communication.NewRequestParam("countryCode", *params.CountryCode)
		reqParams = append(reqParams, *param)
	}
	if params.CurrencyCode != nil {
		param, _ := communication.NewRequestParam("currencyCode", *params.CurrencyCode)
		reqParams = append(reqParams, *param)
	}
	if params.Locale != nil {
		param, _ := communication.NewRequestParam("locale", *params.Locale)
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
	if params.IsInstallments != nil {
		param, _ := communication.NewRequestParam("isInstallments", strconv.FormatBool(*params.IsInstallments))
		reqParams = append(reqParams, *param)
	}
	for _, hideElement := range params.Hide {
		param, _ := communication.NewRequestParam("hide", hideElement)
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewGetParams constructs an instance of GetParams
func NewGetParams() *GetParams {
	return &GetParams{}
}
