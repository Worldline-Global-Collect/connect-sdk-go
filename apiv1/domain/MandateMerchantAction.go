// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MandateMerchantAction represents class MandateMerchantAction
type MandateMerchantAction struct {
	ActionType   *string              `json:"actionType,omitempty"`
	RedirectData *MandateRedirectData `json:"redirectData,omitempty"`
}

// NewMandateMerchantAction constructs a new MandateMerchantAction instance
func NewMandateMerchantAction() *MandateMerchantAction {
	return &MandateMerchantAction{}
}
