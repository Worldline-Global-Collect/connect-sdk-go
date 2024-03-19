// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package tokens

import "github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"

// DeleteParams represents query parameters for Delete token
// Documentation can be found at https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/tokens/delete.html
type DeleteParams struct {
	MandateCancelDate *string
}

// ToRequestParameters converts the query to communication.RequestParams
func (params *DeleteParams) ToRequestParameters() communication.RequestParams {
	reqParams := communication.RequestParams{}

	if params.MandateCancelDate != nil {
		param, _ := communication.NewRequestParam("mandateCancelDate", *params.MandateCancelDate)
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewDeleteParams constructs an instance of DeleteParams
func NewDeleteParams() *DeleteParams {
	return &DeleteParams{}
}
