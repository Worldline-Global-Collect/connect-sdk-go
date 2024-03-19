// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// GetCustomerDetailsRequest represents class GetCustomerDetailsRequest
type GetCustomerDetailsRequest struct {
	CountryCode *string         `json:"countryCode,omitempty"`
	Values      *[]KeyValuePair `json:"values,omitempty"`
}

// NewGetCustomerDetailsRequest constructs a new GetCustomerDetailsRequest
func NewGetCustomerDetailsRequest() *GetCustomerDetailsRequest {
	return &GetCustomerDetailsRequest{}
}
