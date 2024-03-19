// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// AirlinePassenger represents class AirlinePassenger
type AirlinePassenger struct {
	FirstName     *string `json:"firstName,omitempty"`
	Surname       *string `json:"surname,omitempty"`
	SurnamePrefix *string `json:"surnamePrefix,omitempty"`
	Title         *string `json:"title,omitempty"`
}

// NewAirlinePassenger constructs a new AirlinePassenger
func NewAirlinePassenger() *AirlinePassenger {
	return &AirlinePassenger{}
}
