// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// DecryptedPaymentData represents class DecryptedPaymentData
type DecryptedPaymentData struct {
	// Deprecated: Use decryptedPaymentData.paymentMethod instead
	AuthMethod     *string `json:"authMethod,omitempty"`
	CardholderName *string `json:"cardholderName,omitempty"`
	Cryptogram     *string `json:"cryptogram,omitempty"`
	Dpan           *string `json:"dpan,omitempty"`
	Eci            *int32  `json:"eci,omitempty"`
	ExpiryDate     *string `json:"expiryDate,omitempty"`
	Pan            *string `json:"pan,omitempty"`
	PaymentMethod  *string `json:"paymentMethod,omitempty"`
}

// NewDecryptedPaymentData constructs a new DecryptedPaymentData instance
func NewDecryptedPaymentData() *DecryptedPaymentData {
	return &DecryptedPaymentData{}
}
