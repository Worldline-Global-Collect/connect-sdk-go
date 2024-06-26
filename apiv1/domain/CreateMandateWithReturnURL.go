// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CreateMandateWithReturnURL represents class CreateMandateWithReturnUrl
type CreateMandateWithReturnURL struct {
	Alias                  *string          `json:"alias,omitempty"`
	Customer               *MandateCustomer `json:"customer,omitempty"`
	CustomerReference      *string          `json:"customerReference,omitempty"`
	Language               *string          `json:"language,omitempty"`
	RecurrenceType         *string          `json:"recurrenceType,omitempty"`
	ReturnURL              *string          `json:"returnUrl,omitempty"`
	SignatureType          *string          `json:"signatureType,omitempty"`
	UniqueMandateReference *string          `json:"uniqueMandateReference,omitempty"`
}

// NewCreateMandateWithReturnURL constructs a new CreateMandateWithReturnURL instance
func NewCreateMandateWithReturnURL() *CreateMandateWithReturnURL {
	return &CreateMandateWithReturnURL{}
}
