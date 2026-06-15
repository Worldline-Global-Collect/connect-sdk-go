// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// GetIINDetailsResponse represents class GetIINDetailsResponse.
type GetIINDetailsResponse struct {
	BinProductType     *string      `json:"binProductType,omitempty"`
	CoBrands           *[]IINDetail `json:"coBrands,omitempty"`
	CountryCode        *string      `json:"countryCode,omitempty"`
	IsAllowedInContext *bool        `json:"isAllowedInContext,omitempty"`
	IssuerName         *string      `json:"issuerName,omitempty"`
	PaymentProductID   *int32       `json:"paymentProductId,omitempty"`
}

// NewGetIINDetailsResponse constructs a new GetIINDetailsResponse instance.
func NewGetIINDetailsResponse() *GetIINDetailsResponse {
	return &GetIINDetailsResponse{}
}
