// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package webhooks

import (
	"errors"

	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
	"github.com/Worldline-Global-Collect/connect-sdk-go/json"
	"github.com/Worldline-Global-Collect/connect-sdk-go/webhooks/validation"
)

// Helper is the Worldline Global Collect platform v1 webhooks helper. Immutable and thread-safe.
type Helper struct {
	marshaller         json.Marshaller
	signatureValidator *validation.SignatureValidator
}

// Unmarshal unmarshalls the given body and validates it using the requestHeaders
//
// Can return any of the following errors:
// validation.SignatureValidationError if the body could not be validated successfully
// validation.APIVersionMismatchError if the resulting event has an API version that this version of the SDK does not support
func (h Helper) Unmarshal(body string, requestHeaders []communication.Header) (*domain.WebhooksEvent, error) {
	err := h.signatureValidator.Validate(body, requestHeaders)
	if err != nil {
		return nil, err
	}

	event := domain.NewWebhooksEvent()

	err = h.marshaller.Unmarshal(body, event)
	if err != nil {
		return nil, err
	}

	err = validateAPIVersion(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

// Marshaller returns the marshaller of the webhooks helper
func (h Helper) Marshaller() json.Marshaller {
	return h.marshaller
}

// SecretKeyStore returns the secret key store of the webhooks helper
func (h Helper) SecretKeyStore() validation.SecretKeyStore {
	return h.signatureValidator.SecretKeyStore()
}

// NewHelper creates a Helper with the given marshaller and secretKeyStore
func NewHelper(marshaller json.Marshaller, secretKeyStore validation.SecretKeyStore) (*Helper, error) {
	if marshaller == nil {
		return nil, errors.New("nil marshaller")
	}

	signatureValidator, err := validation.NewSignatureValidator(secretKeyStore)
	if err != nil {
		return nil, err
	}

	return &Helper{marshaller, signatureValidator}, nil
}

func validateAPIVersion(event *domain.WebhooksEvent) error {
	if *event.APIVersion != "v1" {
		err := validation.NewAPIVersionMismatchError(*event.APIVersion, "v1")

		return err
	}

	return nil
}
