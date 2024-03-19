// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// Installments represents class Installments
type Installments struct {
	AmountOfMoneyPerInstallment *AmountOfMoney `json:"amountOfMoneyPerInstallment,omitempty"`
	AmountOfMoneyTotal          *AmountOfMoney `json:"amountOfMoneyTotal,omitempty"`
	FrequencyOfInstallments     *string        `json:"frequencyOfInstallments,omitempty"`
	InstallmentPlanCode         *int32         `json:"installmentPlanCode,omitempty"`
	InterestRate                *string        `json:"interestRate,omitempty"`
	NumberOfInstallments        *int64         `json:"numberOfInstallments,omitempty"`
}

// NewInstallments constructs a new Installments
func NewInstallments() *Installments {
	return &Installments{}
}
