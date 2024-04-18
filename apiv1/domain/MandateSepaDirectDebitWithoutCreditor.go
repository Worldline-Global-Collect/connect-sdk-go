// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MandateSepaDirectDebitWithoutCreditor represents class MandateSepaDirectDebitWithoutCreditor
type MandateSepaDirectDebitWithoutCreditor struct {
	BankAccountIban            *BankAccountIban `json:"bankAccountIban,omitempty"`
	CustomerContractIdentifier *string          `json:"customerContractIdentifier,omitempty"`
	Debtor                     *Debtor          `json:"debtor,omitempty"`
	IsRecurring                *bool            `json:"isRecurring,omitempty"`
	MandateApproval            *MandateApproval `json:"mandateApproval,omitempty"`
	PreNotification            *string          `json:"preNotification,omitempty"`
}

// NewMandateSepaDirectDebitWithoutCreditor constructs a new MandateSepaDirectDebitWithoutCreditor instance
func NewMandateSepaDirectDebitWithoutCreditor() *MandateSepaDirectDebitWithoutCreditor {
	return &MandateSepaDirectDebitWithoutCreditor{}
}
