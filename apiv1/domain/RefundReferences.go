// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// RefundReferences represents class RefundReferences
type RefundReferences struct {
	Descriptor        *string `json:"descriptor,omitempty"`
	MerchantReference *string `json:"merchantReference,omitempty"`
}

// NewRefundReferences constructs a new RefundReferences instance
func NewRefundReferences() *RefundReferences {
	return &RefundReferences{}
}
