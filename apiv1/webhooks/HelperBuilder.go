// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package webhooks

import (
	"github.com/Worldline-Global-Collect/connect-sdk-go/json"
	"github.com/Worldline-Global-Collect/connect-sdk-go/webhooks/validation"
)

// HelperBuilder is used to build Helper objects
type HelperBuilder struct {
	Marshaller     json.Marshaller
	SecretKeyStore validation.SecretKeyStore
}

// WithMarshaller sets the marshaller
func (h *HelperBuilder) WithMarshaller(marshaller json.Marshaller) *HelperBuilder {
	h.Marshaller = marshaller

	return h
}

// WithSecretKeyStore sets the secretKeyStore
func (h *HelperBuilder) WithSecretKeyStore(secretKeyStore validation.SecretKeyStore) *HelperBuilder {
	h.SecretKeyStore = secretKeyStore

	return h
}

// Build creates the Helper object
func (h *HelperBuilder) Build() (*Helper, error) {
	return NewHelper(h.Marshaller, h.SecretKeyStore)
}

// NewHelperBuilder creates a HelperBuilder object
func NewHelperBuilder() *HelperBuilder {
	return &HelperBuilder{}
}
