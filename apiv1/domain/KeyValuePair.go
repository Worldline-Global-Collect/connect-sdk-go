// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// KeyValuePair represents class KeyValuePair
type KeyValuePair struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewKeyValuePair constructs a new KeyValuePair instance
func NewKeyValuePair() *KeyValuePair {
	return &KeyValuePair{}
}
