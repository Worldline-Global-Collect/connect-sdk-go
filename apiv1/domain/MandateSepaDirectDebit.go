// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MandateSepaDirectDebit represents class MandateSepaDirectDebit
type MandateSepaDirectDebit struct {
	BankAccountIban            *BankAccountIban `json:"bankAccountIban,omitempty"`
	Creditor                   *Creditor        `json:"creditor,omitempty"`
	CustomerContractIdentifier *string          `json:"customerContractIdentifier,omitempty"`
	Debtor                     *Debtor          `json:"debtor,omitempty"`
	IsRecurring                *bool            `json:"isRecurring,omitempty"`
	MandateApproval            *MandateApproval `json:"mandateApproval,omitempty"`
	MandateID                  *string          `json:"mandateId,omitempty"`
	PreNotification            *string          `json:"preNotification,omitempty"`
}

// NewMandateSepaDirectDebit constructs a new MandateSepaDirectDebit instance
func NewMandateSepaDirectDebit() *MandateSepaDirectDebit {
	return &MandateSepaDirectDebit{}
}
