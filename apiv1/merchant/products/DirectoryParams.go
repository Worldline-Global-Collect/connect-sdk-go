// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package products

import "github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"

// DirectoryParams represents query parameters for Get payment product directory
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/products/directory.html
type DirectoryParams struct {
	CountryCode  *string
	CurrencyCode *string
}

// ToRequestParameters converts the query to communication.RequestParams
func (params *DirectoryParams) ToRequestParameters() communication.RequestParams {
	reqParams := communication.RequestParams{}

	if params.CountryCode != nil {
		param, _ := communication.NewRequestParam("countryCode", *params.CountryCode)
		reqParams = append(reqParams, *param)
	}
	if params.CurrencyCode != nil {
		param, _ := communication.NewRequestParam("currencyCode", *params.CurrencyCode)
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewDirectoryParams constructs an instance of DirectoryParams
func NewDirectoryParams() *DirectoryParams {
	return &DirectoryParams{}
}
