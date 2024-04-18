// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ThreeDSecureBase represents class ThreeDSecureBase
type ThreeDSecureBase struct {
	AuthenticationAmount  *AmountOfMoney    `json:"authenticationAmount,omitempty"`
	AuthenticationFlow    *string           `json:"authenticationFlow,omitempty"`
	ChallengeCanvasSize   *string           `json:"challengeCanvasSize,omitempty"`
	ChallengeIndicator    *string           `json:"challengeIndicator,omitempty"`
	ExemptionRequest      *string           `json:"exemptionRequest,omitempty"`
	PriorThreeDSecureData *ThreeDSecureData `json:"priorThreeDSecureData,omitempty"`
	SdkData               *SdkDataInput     `json:"sdkData,omitempty"`
	SkipAuthentication    *bool             `json:"skipAuthentication,omitempty"`
	TransactionRiskLevel  *string           `json:"transactionRiskLevel,omitempty"`
}

// NewThreeDSecureBase constructs a new ThreeDSecureBase instance
func NewThreeDSecureBase() *ThreeDSecureBase {
	return &ThreeDSecureBase{}
}
