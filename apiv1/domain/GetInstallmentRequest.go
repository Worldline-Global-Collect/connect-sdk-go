// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// GetInstallmentRequest represents class GetInstallmentRequest
type GetInstallmentRequest struct {
	AmountOfMoney    *AmountOfMoney `json:"amountOfMoney,omitempty"`
	Bin              *string        `json:"bin,omitempty"`
	CountryCode      *string        `json:"countryCode,omitempty"`
	PaymentProductID *int32         `json:"paymentProductId,omitempty"`
}

// NewGetInstallmentRequest constructs a new GetInstallmentRequest
func NewGetInstallmentRequest() *GetInstallmentRequest {
	return &GetInstallmentRequest{}
}
