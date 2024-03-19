// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package services

import (
	"strconv"

	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
)

// PrivacypolicyParams represents query parameters for Get privacy policy
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/services/privacypolicy.html
type PrivacypolicyParams struct {
	Locale           *string
	PaymentProductID *int32
}

// ToRequestParameters converts the query to communication.RequestParams
func (params *PrivacypolicyParams) ToRequestParameters() communication.RequestParams {
	reqParams := communication.RequestParams{}

	if params.Locale != nil {
		param, _ := communication.NewRequestParam("locale", *params.Locale)
		reqParams = append(reqParams, *param)
	}
	if params.PaymentProductID != nil {
		param, _ := communication.NewRequestParam("paymentProductId", strconv.FormatInt(int64(*params.PaymentProductID), 10))
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewPrivacypolicyParams constructs an instance of PrivacypolicyParams
func NewPrivacypolicyParams() *PrivacypolicyParams {
	return &PrivacypolicyParams{}
}
