// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// Swift represents class Swift
type Swift struct {
	Bic           *string `json:"bic,omitempty"`
	Category      *string `json:"category,omitempty"`
	ChipsUID      *string `json:"chipsUID,omitempty"`
	ExtraInfo     *string `json:"extraInfo,omitempty"`
	PoBoxCountry  *string `json:"poBoxCountry,omitempty"`
	PoBoxLocation *string `json:"poBoxLocation,omitempty"`
	PoBoxNumber   *string `json:"poBoxNumber,omitempty"`
	PoBoxZip      *string `json:"poBoxZip,omitempty"`
	RoutingBic    *string `json:"routingBic,omitempty"`
	Services      *string `json:"services,omitempty"`
}

// NewSwift constructs a new Swift instance
func NewSwift() *Swift {
	return &Swift{}
}
