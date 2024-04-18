// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ThreeDSecure represents class ThreeDSecure
type ThreeDSecure struct {
	AuthenticationAmount                 *AmountOfMoney                        `json:"authenticationAmount,omitempty"`
	AuthenticationFlow                   *string                               `json:"authenticationFlow,omitempty"`
	ChallengeCanvasSize                  *string                               `json:"challengeCanvasSize,omitempty"`
	ChallengeIndicator                   *string                               `json:"challengeIndicator,omitempty"`
	ExemptionRequest                     *string                               `json:"exemptionRequest,omitempty"`
	ExternalCardholderAuthenticationData *ExternalCardholderAuthenticationData `json:"externalCardholderAuthenticationData,omitempty"`
	PriorThreeDSecureData                *ThreeDSecureData                     `json:"priorThreeDSecureData,omitempty"`
	RedirectionData                      *RedirectionData                      `json:"redirectionData,omitempty"`
	SdkData                              *SdkDataInput                         `json:"sdkData,omitempty"`
	SkipAuthentication                   *bool                                 `json:"skipAuthentication,omitempty"`
	TransactionRiskLevel                 *string                               `json:"transactionRiskLevel,omitempty"`
}

// NewThreeDSecure constructs a new ThreeDSecure instance
func NewThreeDSecure() *ThreeDSecure {
	return &ThreeDSecure{}
}
