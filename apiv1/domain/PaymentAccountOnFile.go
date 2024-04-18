// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentAccountOnFile represents class PaymentAccountOnFile
type PaymentAccountOnFile struct {
	CreateDate                                    *string `json:"createDate,omitempty"`
	NumberOfCardOnFileCreationAttemptsLast24Hours *int32  `json:"numberOfCardOnFileCreationAttemptsLast24Hours,omitempty"`
}

// NewPaymentAccountOnFile constructs a new PaymentAccountOnFile instance
func NewPaymentAccountOnFile() *PaymentAccountOnFile {
	return &PaymentAccountOnFile{}
}
