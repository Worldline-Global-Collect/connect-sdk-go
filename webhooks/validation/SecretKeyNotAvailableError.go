package validation

import "fmt"

// SecretKeyNotAvailableError represents an error that causes a secret key to not be available.
type SecretKeyNotAvailableError struct {
	keyID   string
	message string
}

// Error implements the error interface
func (ske SecretKeyNotAvailableError) Error() string {
	return ske.message
}

// KeyID returns the keyID which produced the error
func (ske SecretKeyNotAvailableError) KeyID() string {
	return ske.keyID
}

// NewSecretKeyNotAvailableError creates a SecretKeyNotAvailableError with the given keyID
func NewSecretKeyNotAvailableError(keyID string) *SecretKeyNotAvailableError {
	message := fmt.Sprintf(`could not find secret key for key id "%v"`, keyID)

	return &SecretKeyNotAvailableError{keyID, message}
}
