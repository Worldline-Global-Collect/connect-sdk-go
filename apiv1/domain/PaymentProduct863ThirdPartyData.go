// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProduct863ThirdPartyData represents class PaymentProduct863ThirdPartyData
type PaymentProduct863ThirdPartyData struct {
	AppID       *string `json:"appId,omitempty"`
	NonceStr    *string `json:"nonceStr,omitempty"`
	PackageSign *string `json:"packageSign,omitempty"`
	PaySign     *string `json:"paySign,omitempty"`
	PrepayID    *string `json:"prepayId,omitempty"`
	SignType    *string `json:"signType,omitempty"`
	TimeStamp   *string `json:"timeStamp,omitempty"`
}

// NewPaymentProduct863ThirdPartyData constructs a new PaymentProduct863ThirdPartyData instance
func NewPaymentProduct863ThirdPartyData() *PaymentProduct863ThirdPartyData {
	return &PaymentProduct863ThirdPartyData{}
}
