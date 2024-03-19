// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// ValueMappingElement represents class ValueMappingElement
type ValueMappingElement struct {
	DisplayElements *[]PaymentProductFieldDisplayElement `json:"displayElements,omitempty"`
	// Deprecated: Use displayElements instead with ID 'displayName'
	DisplayName     *string                              `json:"displayName,omitempty"`
	Value           *string                              `json:"value,omitempty"`
}

// NewValueMappingElement constructs a new ValueMappingElement
func NewValueMappingElement() *ValueMappingElement {
	return &ValueMappingElement{}
}
