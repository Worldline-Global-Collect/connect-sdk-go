// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RecurringPaymentsData represents class RecurringPaymentsData
type RecurringPaymentsData struct {
	PaymentProduct302SpecificInput *PaymentProduct302SpecificInput `json:"paymentProduct302SpecificInput,omitempty"`
	RecurringEndDate               *string                         `json:"recurringEndDate,omitempty"`
	RecurringInterval              *Frequency                      `json:"recurringInterval,omitempty"`
	RecurringStartDate             *string                         `json:"recurringStartDate,omitempty"`
	TrialInformation               *TrialInformation               `json:"trialInformation,omitempty"`
}

// NewRecurringPaymentsData constructs a new RecurringPaymentsData instance
func NewRecurringPaymentsData() *RecurringPaymentsData {
	return &RecurringPaymentsData{}
}
