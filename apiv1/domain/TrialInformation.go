// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// TrialInformation represents class TrialInformation
type TrialInformation struct {
	AmountOfMoneyAfterTrial *AmountOfMoney `json:"amountOfMoneyAfterTrial,omitempty"`
	EndDate                 *string        `json:"endDate,omitempty"`
	IsRecurring             *bool          `json:"isRecurring,omitempty"`
	TrialPeriod             *TrialPeriod   `json:"trialPeriod,omitempty"`
	TrialPeriodRecurring    *Frequency     `json:"trialPeriodRecurring,omitempty"`
}

// NewTrialInformation constructs a new TrialInformation instance
func NewTrialInformation() *TrialInformation {
	return &TrialInformation{}
}
