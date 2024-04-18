// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentReferences represents class PaymentReferences
type PaymentReferences struct {
	MerchantOrderID      *int64  `json:"merchantOrderId,omitempty"`
	MerchantReference    *string `json:"merchantReference,omitempty"`
	PaymentReference     *string `json:"paymentReference,omitempty"`
	ProviderID           *string `json:"providerId,omitempty"`
	ProviderMerchantID   *string `json:"providerMerchantId,omitempty"`
	ProviderReference    *string `json:"providerReference,omitempty"`
	ReferenceOrigPayment *string `json:"referenceOrigPayment,omitempty"`
}

// NewPaymentReferences constructs a new PaymentReferences instance
func NewPaymentReferences() *PaymentReferences {
	return &PaymentReferences{}
}
