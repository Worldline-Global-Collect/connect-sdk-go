// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// MandateResponse represents class MandateResponse
type MandateResponse struct {
	Alias                  *string          `json:"alias,omitempty"`
	Customer               *MandateCustomer `json:"customer,omitempty"`
	CustomerReference      *string          `json:"customerReference,omitempty"`
	RecurrenceType         *string          `json:"recurrenceType,omitempty"`
	Status                 *string          `json:"status,omitempty"`
	UniqueMandateReference *string          `json:"uniqueMandateReference,omitempty"`
}

// NewMandateResponse constructs a new MandateResponse
func NewMandateResponse() *MandateResponse {
	return &MandateResponse{}
}
