// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MerchantAction represents class MerchantAction
type MerchantAction struct {
	ActionType                            *string                                `json:"actionType,omitempty"`
	FormFields                            *[]PaymentProductField                 `json:"formFields,omitempty"`
	MobileThreeDSecureChallengeParameters *MobileThreeDSecureChallengeParameters `json:"mobileThreeDSecureChallengeParameters,omitempty"`
	RedirectData                          *RedirectData                          `json:"redirectData,omitempty"`
	RenderingData                         *string                                `json:"renderingData,omitempty"`
	ShowData                              *[]KeyValuePair                        `json:"showData,omitempty"`
	ThirdPartyData                        *ThirdPartyData                        `json:"thirdPartyData,omitempty"`
}

// NewMerchantAction constructs a new MerchantAction instance
func NewMerchantAction() *MerchantAction {
	return &MerchantAction{}
}
