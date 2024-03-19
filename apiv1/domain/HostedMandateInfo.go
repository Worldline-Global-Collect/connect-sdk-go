// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// HostedMandateInfo represents class HostedMandateInfo
type HostedMandateInfo struct {
	Alias                  *string          `json:"alias,omitempty"`
	Customer               *MandateCustomer `json:"customer,omitempty"`
	CustomerReference      *string          `json:"customerReference,omitempty"`
	RecurrenceType         *string          `json:"recurrenceType,omitempty"`
	SignatureType          *string          `json:"signatureType,omitempty"`
	UniqueMandateReference *string          `json:"uniqueMandateReference,omitempty"`
}

// NewHostedMandateInfo constructs a new HostedMandateInfo
func NewHostedMandateInfo() *HostedMandateInfo {
	return &HostedMandateInfo{}
}
