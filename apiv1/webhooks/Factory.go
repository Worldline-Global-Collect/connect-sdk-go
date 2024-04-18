// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package webhooks

import (
	"github.com/Worldline-Global-Collect/connect-sdk-go/json"
	"github.com/Worldline-Global-Collect/connect-sdk-go/webhooks/validation"
)

// Factory is a factory for several v1 webhooks components
type Factory struct {
}

// NewHelperBuilder creates a new HelperBuilder that will use the given secretKeyStore
func (f Factory) NewHelperBuilder(secretKeyStore validation.SecretKeyStore) *HelperBuilder {
	marshaller := json.DefaultMarshaller()

	return NewHelperBuilder().WithMarshaller(marshaller).WithSecretKeyStore(secretKeyStore)
}

// NewHelper creates a new Helper that will use the given secretKeyStore
func (f Factory) NewHelper(secretKeyStore validation.SecretKeyStore) (*Helper, error) {
	helperBuilder := f.NewHelperBuilder(secretKeyStore)

	return helperBuilder.Build()
}

// NewFactory creates a factory for several v1 webhooks components
func NewFactory() *Factory {
	return &Factory{}
}
