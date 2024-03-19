// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RecurringPaymentsData represents class RecurringPaymentsData
type RecurringPaymentsData struct {
	RecurringInterval *Frequency        `json:"recurringInterval,omitempty"`
	TrialInformation  *TrialInformation `json:"trialInformation,omitempty"`
}

// NewRecurringPaymentsData constructs a new RecurringPaymentsData
func NewRecurringPaymentsData() *RecurringPaymentsData {
	return &RecurringPaymentsData{}
}
