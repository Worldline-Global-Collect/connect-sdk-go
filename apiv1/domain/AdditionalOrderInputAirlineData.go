// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// AdditionalOrderInputAirlineData represents class AdditionalOrderInputAirlineData
type AdditionalOrderInputAirlineData struct {
	AirlineData *AirlineData `json:"airlineData,omitempty"`
	LodgingData *LodgingData `json:"lodgingData,omitempty"`
}

// NewAdditionalOrderInputAirlineData constructs a new AdditionalOrderInputAirlineData
func NewAdditionalOrderInputAirlineData() *AdditionalOrderInputAirlineData {
	return &AdditionalOrderInputAirlineData{}
}
