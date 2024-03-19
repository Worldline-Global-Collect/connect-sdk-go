// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// FraudFieldsShippingDetails represents class FraudFieldsShippingDetails
//
// Deprecated: No replacement
type FraudFieldsShippingDetails struct {
	// Deprecated: No replacement
	MethodDetails *string `json:"methodDetails,omitempty"`
	// Deprecated: No replacement
	MethodSpeed   *int32  `json:"methodSpeed,omitempty"`
	// Deprecated: No replacement
	MethodType    *int32  `json:"methodType,omitempty"`
}

// NewFraudFieldsShippingDetails constructs a new FraudFieldsShippingDetails
func NewFraudFieldsShippingDetails() *FraudFieldsShippingDetails {
	return &FraudFieldsShippingDetails{}
}
