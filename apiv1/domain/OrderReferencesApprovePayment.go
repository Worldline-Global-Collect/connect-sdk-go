// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// OrderReferencesApprovePayment represents class OrderReferencesApprovePayment
type OrderReferencesApprovePayment struct {
	MerchantReference *string `json:"merchantReference,omitempty"`
}

// NewOrderReferencesApprovePayment constructs a new OrderReferencesApprovePayment
func NewOrderReferencesApprovePayment() *OrderReferencesApprovePayment {
	return &OrderReferencesApprovePayment{}
}
