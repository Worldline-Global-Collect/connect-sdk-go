// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// IINDetail represents class IINDetail
type IINDetail struct {
	IsAllowedInContext *bool  `json:"isAllowedInContext,omitempty"`
	PaymentProductID   *int32 `json:"paymentProductId,omitempty"`
}

// NewIINDetail constructs a new IINDetail instance
func NewIINDetail() *IINDetail {
	return &IINDetail{}
}
