// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RefundRequest represents class RefundRequest
type RefundRequest struct {
	AmountOfMoney                 *AmountOfMoney                 `json:"amountOfMoney,omitempty"`
	BankRefundMethodSpecificInput *BankRefundMethodSpecificInput `json:"bankRefundMethodSpecificInput,omitempty"`
	Customer                      *RefundCustomer                `json:"customer,omitempty"`
	RefundDate                    *string                        `json:"refundDate,omitempty"`
	RefundReferences              *RefundReferences              `json:"refundReferences,omitempty"`
}

// NewRefundRequest constructs a new RefundRequest instance
func NewRefundRequest() *RefundRequest {
	return &RefundRequest{}
}
