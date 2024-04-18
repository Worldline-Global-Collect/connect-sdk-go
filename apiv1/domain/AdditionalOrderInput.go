// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// AdditionalOrderInput represents class AdditionalOrderInput
type AdditionalOrderInput struct {
	AccountFundingRecipient *AccountFundingRecipient `json:"accountFundingRecipient,omitempty"`
	AirlineData             *AirlineData             `json:"airlineData,omitempty"`
	Installments            *Installments            `json:"installments,omitempty"`
	// Deprecated: Use Order.shoppingCart.amountBreakdown instead
	Level3SummaryData       *Level3SummaryData       `json:"level3SummaryData,omitempty"`
	// Deprecated: No replacement
	LoanRecipient           *LoanRecipient           `json:"loanRecipient,omitempty"`
	LodgingData             *LodgingData             `json:"lodgingData,omitempty"`
	// Deprecated: Use installments.numberOfInstallments instead
	NumberOfInstallments    *int64                   `json:"numberOfInstallments,omitempty"`
	OrderDate               *string                  `json:"orderDate,omitempty"`
	TypeInformation         *OrderTypeInformation    `json:"typeInformation,omitempty"`
}

// NewAdditionalOrderInput constructs a new AdditionalOrderInput instance
func NewAdditionalOrderInput() *AdditionalOrderInput {
	return &AdditionalOrderInput{}
}
