// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// DisputeOutput represents class DisputeOutput
type DisputeOutput struct {
	AmountOfMoney   *AmountOfMoney         `json:"amountOfMoney,omitempty"`
	ContactPerson   *string                `json:"contactPerson,omitempty"`
	CreationDetails *DisputeCreationDetail `json:"creationDetails,omitempty"`
	EmailAddress    *string                `json:"emailAddress,omitempty"`
	Files           *[]HostedFile          `json:"files,omitempty"`
	Reference       *DisputeReference      `json:"reference,omitempty"`
	ReplyTo         *string                `json:"replyTo,omitempty"`
	RequestMessage  *string                `json:"requestMessage,omitempty"`
	ResponseMessage *string                `json:"responseMessage,omitempty"`
}

// NewDisputeOutput constructs a new DisputeOutput instance
func NewDisputeOutput() *DisputeOutput {
	return &DisputeOutput{}
}
