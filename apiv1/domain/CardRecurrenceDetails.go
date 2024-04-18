// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CardRecurrenceDetails represents class CardRecurrenceDetails
type CardRecurrenceDetails struct {
	EndDate                           *string `json:"endDate,omitempty"`
	MinFrequency                      *int32  `json:"minFrequency,omitempty"`
	RecurringPaymentSequenceIndicator *string `json:"recurringPaymentSequenceIndicator,omitempty"`
}

// NewCardRecurrenceDetails constructs a new CardRecurrenceDetails instance
func NewCardRecurrenceDetails() *CardRecurrenceDetails {
	return &CardRecurrenceDetails{}
}
