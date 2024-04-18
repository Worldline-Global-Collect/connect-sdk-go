// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CreateDisputeRequest represents class CreateDisputeRequest
type CreateDisputeRequest struct {
	AmountOfMoney  *AmountOfMoney `json:"amountOfMoney,omitempty"`
	ContactPerson  *string        `json:"contactPerson,omitempty"`
	EmailAddress   *string        `json:"emailAddress,omitempty"`
	ReplyTo        *string        `json:"replyTo,omitempty"`
	RequestMessage *string        `json:"requestMessage,omitempty"`
}

// NewCreateDisputeRequest constructs a new CreateDisputeRequest instance
func NewCreateDisputeRequest() *CreateDisputeRequest {
	return &CreateDisputeRequest{}
}
