// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RiskAssessmentBankAccount represents class RiskAssessmentBankAccount
type RiskAssessmentBankAccount struct {
	BankAccountBban  *BankAccountBban        `json:"bankAccountBban,omitempty"`
	BankAccountIban  *BankAccountIban        `json:"bankAccountIban,omitempty"`
	FraudFields      *FraudFields            `json:"fraudFields,omitempty"`
	Merchant         *MerchantRiskAssessment `json:"merchant,omitempty"`
	Order            *OrderRiskAssessment    `json:"order,omitempty"`
	PaymentProductID *int32                  `json:"paymentProductId,omitempty"`
}

// NewRiskAssessmentBankAccount constructs a new RiskAssessmentBankAccount
func NewRiskAssessmentBankAccount() *RiskAssessmentBankAccount {
	return &RiskAssessmentBankAccount{}
}
